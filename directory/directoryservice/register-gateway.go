package directoryservice

import (
	"context"
	"fmt"

	"github.com/VNG-Realisatie/nlx/directory/directoryapi"
)

type registerGatewayHandler struct{}

func newRegisterGatewayHandler() (*registerGatewayHandler, error) {
	return &registerGatewayHandler{}, nil
}

func (p *registerGatewayHandler) RegisterGateway(ctx context.Context, req *directoryapi.RegisterGatewayRequest) (*directoryapi.RegisterGatewayResponse, error) {
	fmt.Printf("rpc request RegisterGateway(%s)\n", req.Name)
	repl := &directoryapi.RegisterGatewayResponse{}

	return repl, nil
}
