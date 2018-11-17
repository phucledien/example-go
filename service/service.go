package service

import (
	"github.com/phucledien/example-go/service/category"
	"github.com/phucledien/example-go/service/user"
)

// Service define list of all services in projects
type Service struct {
	UserService     user.Service
	CategoryService category.Service
}
