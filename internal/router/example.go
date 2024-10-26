package router

import (
	"context"
	"fmt"

	"connectrpc.com/connect"
	protogen "github.com/dena-autumn-2024-g/wrp-server/internal/router/protogen/protobuf"
	"github.com/dena-autumn-2024-g/wrp-server/internal/router/protogen/protobuf/protogenconnect"
)

var _ protogenconnect.ExampleServiceHandler = &Example{}

type Example struct{}

func NewExample() *Example {
	return &Example{}
}

func (g *Example) Greet(ctx context.Context, req *connect.Request[protogen.GreetRequest]) (*connect.Response[protogen.GreetResponse], error) {
	res := connect.NewResponse(&protogen.GreetResponse{
		Greeting: fmt.Sprintf("Hello, %s!", req.Msg.Name),
	})

	return res, nil
}
