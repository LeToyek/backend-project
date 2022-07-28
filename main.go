package main

import (
	"database/sql"
	"fmt"
	"project/handler"
	"project/repository"
	"project/server"
	"project/service"

	"github.com/gin-gonic/gin"

	_ "github.com/lib/pq"
)

const (
	user      = "postgres"
	password  = "handoko"
	dbname    = "project"
	port      = 5432
	ssl       = "disable"
	timeZone  = "Asia/Jakarta"
	localhost = "localhost:5000"
)

func main() {

	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s", user, password, dbname, port, ssl, timeZone)

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
	panic(router.Run(localhost))
}
