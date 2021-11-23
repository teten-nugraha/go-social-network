package main

import (
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/hako/branca"
	"github.com/teten-nugraha/go-social-network/internal/handler"
	"github.com/teten-nugraha/go-social-network/internal/service"
	"log"
	"net/http"
)

const (
	dbURL = "postgres://postgres:postgres@localhost:5432/nakama?sslmode=disable"
	port  = 3000
)

func main() {
	db, err := sql.Open("pgx", dbURL)
	if err != nil {
		log.Fatalf("could not open db connection: %v\n", err)
		return
	}

	defer db.Close()
	if err = db.Ping(); err != nil {
		log.Fatalf("could not ping to db: %v\n", err)
		return
	}

	codec := branca.NewBranca("supersecretkeyyoushouldnotcommit")
	s := service.New(db, codec)

	h := handler.New(s)
	addr := fmt.Sprintf(":%d", port)
	log.Printf("accepting connections on port %d", port)
	if err = http.ListenAndServe(addr, h); err != nil {
		log.Fatalf("could not start server: %v\n", err)
	}
}
