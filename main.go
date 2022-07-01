package main

import (
	"coin-batam/handler"
	"coin-batam/repository"
	"coin-batam/server"
	"coin-batam/service"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

const (
	user      = "postgres"
	password  = "handoko"
	dbname    = "coin-batam"
	port      = 5432
	ssl       = "disable"
	time      = "Asia/Jakarta"
	localhost = "localhost:5000"
)

func main() {

	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s", user, password, dbname, port, ssl, time)

	db, err := sql.Open("postgres", dbInfo)

	if err != nil {
		panic(err.Error())
	}

	repository := repository.Repository{
		DB: db,
	}
	service := service.Service{
		Repository: repository,
	}
	handler := handler.Handler{
		Service: service,
	}

	router := gin.Default()

	start := server.Server{
		Router:  router,
		Handler: &handler,
	}

	start.StartServer()

	router.Run(localhost)

}
