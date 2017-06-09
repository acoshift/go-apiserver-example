package user

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/acoshift/go-apiserver-example/pkg/api"
)

// New creates new user controller
func New(db *sql.DB) api.UserController {
	return &userCtrl{db}
}

type userCtrl struct {
	db *sql.DB
}

func (c *userCtrl) Get(ctx context.Context, req *api.IDRequest) (*api.User, error) {
	if req.ID <= 0 {
		return nil, fmt.Errorf("invalid id")
	}

	// better to move database logic to model
	resp := api.User{}
	err := c.db.QueryRowContext(ctx, `
		select id, name, email
		from users
		where id = $1
	`, req.ID).Scan(&resp.ID, &resp.Name, &resp.Email)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (c *userCtrl) List(ctx context.Context, req *api.ListRequest) (*api.UsersResponse, error) {
	rows, err := c.db.QueryContext(ctx, `
		select id, name, email
		from users
		limit $1 offset $2
	`, req.Limit, req.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := make([]*api.User, 0)
	for rows.Next() {
		var user api.User
		err = rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			return nil, err
		}
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &api.UsersResponse{Users: users}, nil
}

func (c *userCtrl) Delete(ctx context.Context, req *api.IDRequest) (*api.Empty, error) {
	_, err := c.db.ExecContext(ctx, `
		delete from users
		where id = $1
	`, req.ID)
	if err != nil {
		return nil, err
	}
	return new(api.Empty), nil
}
