package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/acoshift/configfile"
	"github.com/acoshift/go-apiserver-example/pkg/app"
	"github.com/acoshift/go-apiserver-example/pkg/user"
	_ "github.com/lib/pq"
)

func main() {
	config := configfile.NewReader("config")
	sqlURL := config.String("sql_url")

	db, err := sql.Open("postgres", sqlURL)
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	app.MountUserController(mux, user.New(db))

	log.Println("start server at :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
