package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"github.com/l1rn/balance-app/controllers"
	"github.com/l1rn/balance-app/services"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(
		sqlite.Open("db/data.db"),
		&gorm.Config{},
	)

	if err != nil {
		fmt.Println("Failed to connect to db: ", err)
		return
	}

	err = db.AutoMigrate(
		&services.User{},
	)

	if err != nil {
		fmt.Println("Failed to init db: ", err)
	}

	r := gin.Default()
	userService := services.NewUserService(db)
	userCtrl := controllers.NewUserController(userService)

	admin := r.Group("/api/admin")
	admin.Use(controllers.AdminOnly())
	{
		admin.GET("/all-users", userCtrl.GetAllUsers)
		admin.POST("/create-user", userCtrl.CreateUser)
		admin.POST("/top-up", userCtrl.TopUpBalance)
	}

	r.Run(":8080")
}
