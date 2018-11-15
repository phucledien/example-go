package endpoints

import (
	"github.com/go-kit/kit/endpoint"
	"github.com/phucledien/example-go/service"

	"github.com/phucledien/example-go/endpoints/user"

	"github.com/phucledien/example-go/endpoints/category"
)

// Endpoints .
type Endpoints struct {
	FindUser    endpoint.Endpoint
	FindAllUser endpoint.Endpoint
	CreateUser  endpoint.Endpoint
	UpdateUser  endpoint.Endpoint
	DeleteUser  endpoint.Endpoint

	FindCategory    endpoint.Endpoint
	FindAllCategory endpoint.Endpoint
	CreateCategory  endpoint.Endpoint
	UpdateCategory  endpoint.Endpoint
	DeleteCategory  endpoint.Endpoint
}

// MakeServerEndpoints returns an Endpoints struct
func MakeServerEndpoints(s service.Service) Endpoints {
	return Endpoints{
		FindUser:    user.MakeFindEndpoint(s),
		FindAllUser: user.MakeFindAllEndpoint(s),
		CreateUser:  user.MakeCreateEndpoint(s),
		UpdateUser:  user.MakeUpdateEndpoint(s),
		DeleteUser:  user.MakeDeleteEndpoint(s),

		FindCategory:    category.MakeFindEndpoint(s),
		FindAllCategory: category.MakeFindAllEndpoint(s),
		CreateCategory:  category.MakeCreateEndpoint(s),
		UpdateCategory:  category.MakeUpdateEndpoint(s),
		DeleteCategory:  category.MakeDeleteEndpoint(s),
	}
}
