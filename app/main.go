package main

import (
	"github.com/joho/godotenv"
	"log"
	"main/api"
	"main/api/fetch"
	"os"
	"strconv"
	_ "strconv"
)

func getEnv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	maxOpenConnections, _ := strconv.Atoi(getEnv("PG_MAX_OPEN_CONNS", "30"))
	maxIdleConnections, _ := strconv.Atoi(getEnv("PG_MAX_IDLE_CONNS", "30"))
	maxIdleTime := getEnv("PG_IDLE_TIME", "15min")

	dbConfig := api.DbConfig{
		Address:            getEnv("PG_ADRESS", "postgres://admin:adminpassword@localhost/GoCard?sslmode=disable"),
		MaxOpenConnections: maxOpenConnections,
		MaxIdleConnections: maxIdleConnections,
		MaxIdleTime:        maxIdleTime,
	}

	cfg := api.Config{
		Address: ":8080",
		Version: "0.0.1",
		Db:      dbConfig,
	}

	db, err := fetch.New(
		cfg.Db.Address,
		cfg.Db.MaxOpenConnections,
		cfg.Db.MaxIdleConnections,
		cfg.Db.MaxIdleTime,
	)

	if err != nil {
		log.Panic(err)
	}

	defer db.Close()
	log.Printf("database connection pool established")

	store := fetch.NewPostGresStorage(db)

	app := &api.Application{
		Config: cfg,
		Store:  store,
	}

	mux := app.Mount()

	log.Fatal(app.Run(mux))
}
