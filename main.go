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

	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"https://localhost:5000/register/"},
	// 	AllowMethods:     []string{"PUT", "PATCH", "POST"},
	// 	AllowHeaders:     []string{"Content-Type"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	AllowOriginFunc: func(origin string) bool {
	// 		return origin == "https://github.com"
	// 	},
	// 	MaxAge: 12 * time.Hour,
	// }))
	router.Run(localhost)
}
