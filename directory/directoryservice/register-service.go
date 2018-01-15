package directoryservice

import (
	"context"
	"fmt"

	"github.com/VNG-Realisatie/nlx/directory/directoryapi"
)

type registerServiceHandler struct{}

func newRegisterServiceHandler() (*registerServiceHandler, error) {
	return &registerServiceHandler{}, nil
}

func (p *registerServiceHandler) RegisterService(ctx context.Context, req *directoryapi.RegisterServiceRequest) (*directoryapi.RegisterServiceResponse, error) {
	fmt.Printf("rpc request RegisterService(%s)\n", req.Name)
	repl := &directoryapi.RegisterServiceResponse{}

	return repl, nil
}
