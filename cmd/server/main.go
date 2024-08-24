// cmd/server/main.go
package main

import (
	"log"
	"practice02/configs"
	"practice02/internal/domain"
	"practice02/internal/handler"
	"practice02/internal/repository/mysql"
	"practice02/internal/usecase"

	"github.com/labstack/echo/v4"
	gormMysql "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
    db, err := gorm.Open(gormMysql.Open(configs.GetDSN()), &gorm.Config{})
    if err != nil {
        log.Fatalf("failed to connect to database: %v", err)
    }

    db.AutoMigrate(&domain.Todo{})

    todoRepo := mysql.NewTodoRepositoryMysql(db)
    todoUsecase := usecase.NewTodoUsecase(todoRepo)

    e := echo.New()
    handler.NewTodoHandler(e, todoUsecase)

    log.Fatal(e.Start(":8080"))
}
