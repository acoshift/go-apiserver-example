package user_test

import (
	"context"
	"database/sql"
	"testing"

	"github.com/acoshift/go-apiserver-example/pkg/api"
	"github.com/acoshift/go-apiserver-example/pkg/user"
	_ "github.com/lib/pq"
)

func getMockSQL() (*sql.DB, error) {
	// TODO
	return sql.Open("postgres", "postgres://localhost:5432")
}

func generateMockData(db *sql.DB) {
	// TODO
}

func cleanup(db *sql.DB) {
	// TODO:
}

func TestGet(t *testing.T) {
	db, err := getMockSQL()
	if err != nil {
		t.Fatalf("can not create mock database; %v", err)
	}
	generateMockData(db)
	defer cleanup(db)

	ctrl := user.New(db)
	resp, err := ctrl.Get(context.Background(), &api.IDRequest{ID: 12345})
	if err != nil {
		t.Errorf("expected get user 12345 return valid user; got %v", err)
	}
	if resp.ID != 12345 {
		t.Errorf("expetect user id to be 12345; got %v", resp.ID)
	}
	// TODO
}
