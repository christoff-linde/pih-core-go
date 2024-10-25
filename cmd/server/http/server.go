package http

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	db "github.com/christoff-linde/pih-core-go/internal/database"
	"github.com/jackc/pgx/v5/pgxpool"

	_ "github.com/joho/godotenv/autoload"
)

func initDb(databaseUrl string) *db.Queries {
	connection, err := pgxpool.New(context.Background(), databaseUrl)
	if err != nil {
		log.Panic(err, "Unable to connect to database: %v\n", err)
	}

	db := db.New(connection)
	return db
}

type Server struct {
	db   *db.Queries
	port int
}

func NewServer(port int, dbConn *db.Queries) *http.Server {
	NewServer := &Server{
		port: port,
		db:   dbConn,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
