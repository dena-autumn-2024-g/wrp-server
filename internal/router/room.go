package router

import (
	"context"
	"errors"
	"fmt"

	"connectrpc.com/connect"
	"github.com/dena-autumn-2024-g/wrp-server/internal/memdb"
	protogen "github.com/dena-autumn-2024-g/wrp-server/internal/router/protogen/protobuf"
	"github.com/google/uuid"
)

type Room struct{}

func NewRoom() *Room {
	return &Room{}
}

func (r *Room) CreateRoom(ctx context.Context, req *connect.Request[protogen.CreateRoomRequest]) (*connect.Response[protogen.CreateRoomResponse], error) {
	roomID := uuid.New()

	memdb.CreateRoom(roomID)

	return &connect.Response[protogen.CreateRoomResponse]{
		Msg: &protogen.CreateRoomResponse{
			RoomId:  roomID.String(),
			RoomUrl: fmt.Sprintf("https://wrp.mazrean.com/?roomID=%s&serverName=%s", roomID.String(), "a"),
		},
	}, nil
}

func (r *Room) WaitForUserJoin(ctx context.Context, req *connect.Request[protogen.WaitForUserJoinRequest], stream *connect.ServerStream[protogen.WaitForUserJoinResponse]) error {
	var (
		joinChan <-chan uint32
		err      error
	)
	joinChan, err = memdb.GetStream(uuid.MustParse(req.Msg.RoomId))
	if err != nil {
		return fmt.Errorf("failed to get stream: %w", err)
	}

	for {
		select {
		case userID := <-joinChan:
			if err := stream.Send(&protogen.WaitForUserJoinResponse{
				UserId: int32(userID),
			}); err != nil {
				return fmt.Errorf("failed to send stream: %w", err)
			}
		case <-ctx.Done():
			return nil
		}
	}
}

func (r *Room) JoinRoom(ctx context.Context, req *connect.Request[protogen.JoinRoomRequest]) (*connect.Response[protogen.JoinRoomResponse], error) {
	userID, err := memdb.Join(uuid.MustParse(req.Msg.RoomId))
	switch {
	case errors.Is(err, memdb.ErrRoomNotFound):
		return nil, connect.NewError(connect.CodeNotFound, err)
	case errors.Is(err, memdb.ErrRoomIsFull):
		return nil, connect.NewError(connect.CodeResourceExhausted, err)
	case err != nil:
		return nil, fmt.Errorf("failed to join room: %w", err)
	}

	return &connect.Response[protogen.JoinRoomResponse]{
		Msg: &protogen.JoinRoomResponse{
			UserId: int32(userID),
		},
	}, nil
}

func (r *Room) CloseRoom(ctx context.Context, req *connect.Request[protogen.CloseRoomRequest]) (*connect.Response[protogen.CloseRoomResponse], error) {
	roomID, err := uuid.Parse(req.Msg.RoomId)
	if err != nil {
		return nil, fmt.Errorf("failed to parse room id: %w", err)
	}

	memdb.DeleteRoom(roomID)

	return &connect.Response[protogen.CloseRoomResponse]{
		Msg: &protogen.CloseRoomResponse{},
	}, nil
}

func (r *Room) CheckLiveness(ctx context.Context, req *connect.Request[protogen.CheckLivenessRequest]) (*connect.Response[protogen.CheckLivenessResponse], error) {
	_, err := memdb.GetStream(uuid.MustParse(req.Msg.RoomId))
	if err != nil {
		return &connect.Response[protogen.CheckLivenessResponse]{
			Msg: &protogen.CheckLivenessResponse{
				IsAlive: false,
			},
		}, nil
	}
	return &connect.Response[protogen.CheckLivenessResponse]{
		Msg: &protogen.CheckLivenessResponse{
			IsAlive: true,
		},
	}, nil
}
