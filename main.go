package main

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/christoff-linde/pih-core-go/cmd/server/http"
	"github.com/christoff-linde/pih-core-go/cmd/subscriber"
	db "github.com/christoff-linde/pih-core-go/internal/database"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
)

func initDb(databaseUrl string) *db.Queries {
	connection, err := pgxpool.New(context.Background(), databaseUrl)
	if err != nil {
		log.Panic(err, "Unable to connect to database: %v\n", err)
	}

	db := db.New(connection)
	return db
}

//go:embed db/schema/*.sql
var embedMigrations embed.FS

func runMigrations(databaseUrl string) {
	sqlDB, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err, "Unable to create sql.DB connection: %v\n", err)
	}
	defer sqlDB.Close()

	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal(err)
	}
	if err := goose.Up(sqlDB, "db/schema"); err != nil {
		log.Fatal(err)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file", err)
	}

	brokerUrl := os.Getenv("BROKER_URL")
	fmt.Println("Found BROKER_URL with value:", brokerUrl)

	databaseUrl := os.Getenv("DB_URL")
	fmt.Println("Found DB_URL with value:", databaseUrl)

	runMigrations(databaseUrl)

	dbConn := initDb(databaseUrl)

	subscriberCfg := subscriber.AppConfig{DB: dbConn}

	go subscriber.StartSubscriber(subscriberCfg, brokerUrl)

	port, _ := strconv.Atoi(os.Getenv("API_PORT"))
	server := http.NewServer(port, dbConn)
	log.Fatal(server.ListenAndServe())
}
