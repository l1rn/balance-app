package services

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
)

var SecretKey = []byte("super_key")

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
		Role: RoleAdmin,
	}

	s.AddUser(request)
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
