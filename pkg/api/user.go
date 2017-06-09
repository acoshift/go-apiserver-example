package api

import "context"

// UserController type
type UserController interface {
	Get(context.Context, *IDRequest) (*User, error)
	List(context.Context, *ListRequest) (*UsersResponse, error)
	Delete(context.Context, *IDRequest) (*Empty, error)
}

// User type
type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// UsersResponse type
type UsersResponse struct {
	Users []*User `json:"users"`
}
