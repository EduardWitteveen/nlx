package directoryservice

import (
	"context"
	"fmt"

	"github.com/VNG-Realisatie/nlx/directory/directoryapi"
)

type getServicesHandler struct{}

func newGetServicesHandler() (*getServicesHandler, error) {
	return &getServicesHandler{}, nil
}

func (p *getServicesHandler) GetServices(ctx context.Context, req *directoryapi.GetServicesRequest) (*directoryapi.GetServicesResponse, error) {
	fmt.Println("rpc request GetServices()")
	repl := &directoryapi.GetServicesResponse{}

	return repl, nil
}
