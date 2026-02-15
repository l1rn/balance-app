package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"github.com/joho/godotenv"
	"github.com/l1rn/balance-app/controllers"
	"github.com/l1rn/balance-app/services"
	"gorm.io/gorm"
)	

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file! Check if it exists")
	}
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
		&services.TransactionItem{},
	)

	if err != nil {
		fmt.Println("Failed to init db: ", err)
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"https://30c9-2a09-bac5-3b6c-16b4-00-243-7c.ngrok-free.app", "http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))
	userService := services.NewUserService(db)
	userCtrl := controllers.NewUserController(userService)

	transactionService := services.NewTransactionService(db)
	transactionCtrl := controllers.NewTransactionController(transactionService)

	var count int64
	err = db.Find(&services.User{}).Count(&count).Error
	if err != nil {
		log.Println("Failed to execute counter for user data. SQL seed is stopped.")
		return
	}
	fmt.Println("count:", count)
	if count <= 0 {
		services.SeedAdmin(userService)
	}

	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		auth.POST("/login", func(ctx *gin.Context) {
			var req services.UserRequest

			if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
				ctx.JSON(400, gin.H{"error": err.Error()})
				return
			}

			token, err := services.Login(db, req.Username, req.Password)
			if err != nil {
				ctx.JSON(401, gin.H{"error": err})
				return
			}
			ctx.SetCookie("token", token, 3600*24, "/", "", true, true)
		})
		admin := api.Group("/admin")
		admin.Use(controllers.AuthMiddleware("admin"))
		{
			admin.GET("/all-users", userCtrl.GetAllUsers)
			admin.GET("/check-balance/:id", userCtrl.CheckBalance)
			admin.POST("/create-user", userCtrl.CreateUser)
			admin.POST("/top-up", userCtrl.TopUpBalance)
			admin.POST("/withdraw", userCtrl.WithdrawBalance)

			admin.GET("/users/:id", transactionCtrl.GetUserHistory)
		}
	}
	
	r.Static("/assets", "./dist/assets")
	r.StaticFile("/favicon.ico", "./dist/favicon.ico")		
	r.StaticFile("/", "./dist/index.html")
	
	r.NoRoute(func(c *gin.Context) {
    	path := c.Request.URL.Path
		if strings.HasPrefix(path, "/api") {
			c.JSON(404, gin.H{"error": "API route not found"})
			return
		}
    	c.File("./dist/index.html")
	})
	r.Run(":8080")
}
