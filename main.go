package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/christoff-linde/pih-core-go/cmd/server/http"
	"github.com/christoff-linde/pih-core-go/cmd/subscriber"
	db "github.com/christoff-linde/pih-core-go/internal/database"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func initDb(databaseUrl string) *db.Queries {
	connection, err := pgxpool.New(context.Background(), databaseUrl)
	if err != nil {
		log.Panic(err, "Unable to connect to database: %v\n", err)
	}

	db := db.New(connection)
	return db
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	brokerUrl := os.Getenv("BROKER_URL")
	fmt.Println("Found BROKER_URL with value:", brokerUrl)

	databaseUrl := os.Getenv("DB_URL")
	fmt.Println("Found DB_URL with value:", databaseUrl)

	dbConn := initDb(databaseUrl)
	subscriberCfg := subscriber.AppConfig{DB: dbConn}

	go subscriber.StartSubscriber(subscriberCfg, brokerUrl)

	port, _ := strconv.Atoi(os.Getenv("API_PORT"))
	server := http.NewServer(port, dbConn)
	log.Fatal(server.ListenAndServe())

}
