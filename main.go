package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	todos "jubo-hw/apiv1/todos"
	docs "jubo-hw/docs"
	models "jubo-hw/models"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db_database := os.Getenv("DB_DATABASE")
	db_username := os.Getenv("DB_USERNAME")
	db_password := os.Getenv("DB_PASSWORD")
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	models.InitDB(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&timeout=5000ms",
		db_username, db_password, db_host, db_port, db_database))
	defer models.CloseDB()

	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		todosRoute := v1.Group("/todos")
		{
			todosRoute.GET("/", todos.GetTODOItems)
			todosRoute.GET("/:id", todos.GetTODOItem)
			todosRoute.POST("/", todos.CreateTODOItem)
			todosRoute.PUT("/:id", todos.UpdateTODOItem)
			todosRoute.DELETE("/:id", todos.DeleteTODOItem)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8787")
}
