package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/KKGo-Software-engineering/fun-exercise-api/postgres"
	"github.com/KKGo-Software-engineering/fun-exercise-api/wallet"
	"github.com/labstack/echo/v4"

	_ "github.com/KKGo-Software-engineering/fun-exercise-api/docs"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title			Wallet API
// @version		1.0
// @description	Sophisticated Wallet API
// @host			localhost:1323
func main() {
	host := os.Getenv("DB_HOST")
	port, _ := strconv.Atoi(os.Getenv("DB_PORT")) // Convert port to int
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	dbConfig := postgres.Config{
		DatabaseURL: fmt.Sprintf("host=%s port=%d user=%s "+
			"password=%s dbname=%s sslmode=disable",
			host, port, user, password, dbname),
	}
	p, err := postgres.New(dbConfig)
	if err != nil {
		panic(err)
	}

	e := echo.New()
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	handler := wallet.New(p)
	e.GET("/api/v1/wallets", handler.GetAllWalletsHandler)
	e.GET("/api/v1/users/:id/wallets", handler.GetWalletByIDHandler)
	e.POST("/api/v1/wallets", handler.CreateWalletHandler)
	e.PUT("/api/v1/wallets/:id", handler.UpdateWalletHandler)
	e.DELETE("/api/v1/users/:id/wallets", handler.DeleteWalletByIDHandler)
	e.Logger.Fatal(e.Start(":1323"))
}
