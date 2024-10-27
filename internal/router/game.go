package router

import (
	"context"
	"fmt"
	"sync"

	"connectrpc.com/connect"
	protogen "github.com/dena-autumn-2024-g/wrp-server/internal/router/protogen/protobuf"
)

type Game struct{}

func NewGame() *Game {
	return &Game{}
}

var (
	gameMap       = map[string]chan *protogen.StartGameStreamResponse{}
	gameMapLocker = &sync.RWMutex{}
)

func (g *Game) StartGameStream(ctx context.Context, req *connect.Request[protogen.StartGameStreamRequest], stream *connect.ServerStream[protogen.StartGameStreamResponse]) error {
	streamChan := make(chan *protogen.StartGameStreamResponse, 10000)
	func() {
		gameMapLocker.Lock()
		defer gameMapLocker.Unlock()

		gameMap[req.Msg.RoomId] = streamChan
	}()

	for {
		select {
		case <-ctx.Done():
			return nil
		case res := <-streamChan:
			if err := stream.Send(res); err != nil {
				return fmt.Errorf("failed to send stream: %w", err)
			}
		}
	}
}

func (g *Game) Move(ctx context.Context, req *connect.Request[protogen.MoveRequest]) (*connect.Response[protogen.MoveResponse], error) {
	streamChan, ok := func() (chan<- *protogen.StartGameStreamResponse, bool) {
		gameMapLocker.RLock()
		defer gameMapLocker.RUnlock()

		if _, ok := gameMap[req.Msg.RoomId]; !ok {
			return nil, false
		}

		return gameMap[req.Msg.RoomId], true
	}()
	if !ok {
		return &connect.Response[protogen.MoveResponse]{
			Msg: &protogen.MoveResponse{},
		}, nil
	}

	streamChan <- &protogen.StartGameStreamResponse{
		Event: &protogen.StartGameStreamResponse_MoveButton{
			MoveButton: &protogen.MoveRequest{
				UserId: req.Msg.UserId,
			},
		},
	}

	return &connect.Response[protogen.MoveResponse]{
		Msg: &protogen.MoveResponse{},
	}, nil
}

func (g *Game) PushButton(tcx context.Context, req *connect.Request[protogen.PushButtonRequest]) (*connect.Response[protogen.PushButtonResponse], error) {
	streamChan, ok := func() (chan<- *protogen.StartGameStreamResponse, bool) {
		gameMapLocker.RLock()
		defer gameMapLocker.RUnlock()

		if _, ok := gameMap[req.Msg.RoomId]; !ok {
			return nil, false
		}

		return gameMap[req.Msg.RoomId], true
	}()
	if !ok {
		return &connect.Response[protogen.PushButtonResponse]{
			Msg: &protogen.PushButtonResponse{},
		}, nil
	}

	streamChan <- &protogen.StartGameStreamResponse{
		Event: &protogen.StartGameStreamResponse_PushButtonPressed{
			PushButtonPressed: &protogen.PushButtonRequest{
				UserId: req.Msg.UserId,
				RoomId: req.Msg.RoomId,
			},
		},
	}

	return &connect.Response[protogen.PushButtonResponse]{
		Msg: &protogen.PushButtonResponse{},
	}, nil
}

func (g *Game) ReleaseButton(ctx context.Context, req *connect.Request[protogen.ReleaseButtonRequest]) (*connect.Response[protogen.ReleaseButtonResponse], error) {
	streamChan, ok := func() (chan<- *protogen.StartGameStreamResponse, bool) {
		gameMapLocker.RLock()
		defer gameMapLocker.RUnlock()

		if _, ok := gameMap[req.Msg.RoomId]; !ok {
			return nil, false
		}

		return gameMap[req.Msg.RoomId], true
	}()
	if !ok {
		return &connect.Response[protogen.ReleaseButtonResponse]{
			Msg: &protogen.ReleaseButtonResponse{},
		}, nil
	}

	streamChan <- &protogen.StartGameStreamResponse{
		Event: &protogen.StartGameStreamResponse_PushButtonReleased{
			PushButtonReleased: &protogen.ReleaseButtonRequest{
				UserId: req.Msg.UserId,
				RoomId: req.Msg.RoomId,
			},
		},
	}

	return &connect.Response[protogen.ReleaseButtonResponse]{
		Msg: &protogen.ReleaseButtonResponse{},
	}, nil
}
