package memdb

import (
	"errors"
	"sync"
	"sync/atomic"

	"github.com/google/uuid"
)

type RoomItem struct {
	userCount *atomic.Uint32
	joinChan  chan uint32
}

var (
	userCountMap       = map[uuid.UUID]*RoomItem{}
	userCountMapLocker = &sync.RWMutex{}
)

func CreateRoom(roomID uuid.UUID) {
	atomicUint := &atomic.Uint32{}
	atomicUint.Store(0)

	userCountMapLocker.Lock()
	defer userCountMapLocker.Unlock()

	userCountMap[roomID] = &RoomItem{
		userCount: atomicUint,
		joinChan:  make(chan uint32, 100),
	}
}

func Join(roomID uuid.UUID) (uint32, error) {
	roomItem, ok := func() (*RoomItem, bool) {
		userCountMapLocker.RLock()
		defer userCountMapLocker.RUnlock()

		if _, ok := userCountMap[roomID]; !ok {
			return nil, false
		}

		return userCountMap[roomID], true
	}()
	if !ok {
		return 0, errors.New("room not found")
	}

	userID := roomItem.userCount.Add(1)
	roomItem.joinChan <- userID

	return userID - 1, nil
}

func GetStream(roomID uuid.UUID) (chan uint32, error) {
	roomItem, ok := func() (*RoomItem, bool) {
		userCountMapLocker.RLock()
		defer userCountMapLocker.RUnlock()

		if _, ok := userCountMap[roomID]; !ok {
			return nil, false
		}

		return userCountMap[roomID], true
	}()
	if !ok {
		return nil, errors.New("room not found")
	}

	return roomItem.joinChan, nil
}

func DeleteRoom(roomID uuid.UUID) {
	userCountMapLocker.Lock()
	defer userCountMapLocker.Unlock()

	roomItem, ok := userCountMap[roomID]
	if !ok {
		return
	}

	close(roomItem.joinChan)
	delete(userCountMap, roomID)
}
