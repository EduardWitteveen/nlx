package directoryservice

import (
	"github.com/VNG-Realisatie/nlx/directory/directoryapi"
)

var _ directoryapi.DirectoryServer = &DirectoryService{}

// DirectoryService handles all requests for a directory api
type DirectoryService struct {
	*registerGatewayHandler
	*registerServiceHandler
	*getServicesHandler
}
