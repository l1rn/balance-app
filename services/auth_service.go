package services

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

var s, _ = os.LookupEnv("JWT_SECRET")
var SecretKey = []byte(s)

func SeedAdmin(s UserService) {
	user, exists := os.LookupEnv("ADMIN_USERNAME")
	if !exists {
		log.Println("⚠️ WARNING: ADMIN_USERNAME not found in .env, using default 'admin'")
		user = "admin"
	}

	pass, _ := os.LookupEnv("ADMIN_PASSWORD")
	if pass == "" {
		pass = "admin"
	}

	request := UserRequest{
		Username: user,
		Password: pass,
		Role:     RoleAdmin,
	}

	user1, _ := os.LookupEnv("USERNAME1")
	pwd1, _ := os.LookupEnv("USERNAME1")
	user2, _ := os.LookupEnv("USERNAME2")
	pwd2, _ := os.LookupEnv("USERNAME2")

	request1 := UserRequest{
		Username: user1,
		Password: pwd1,
		Role:     RoleUser,
	}

	request2 := UserRequest{
		Username: user2,
		Password: pwd2,
		Role:     RoleUser,
	}

	s.AddUser(request)
	s.AddUser(request1)
	s.AddUser(request2)
}

func Login(db *gorm.DB, username, password string) (string, error) {
	var user User
	if err := db.Where("username = ?", username).First(&user).Error; err != nil {
		return "", err
	}

	if err := user.CheckPassword(password); err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":  user.ID,
		"role": user.Role,
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
	})

	return token.SignedString(SecretKey)
}
