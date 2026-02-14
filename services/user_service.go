package services

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique;not null" json:"username"`
	Password string `gorm:"not null" json:"password"`
	Role     Role   `gorm:"type:varchar(20);default:'user'" json:"role"`
	Balance  int64  `gorm:"default:0" json:"balance"`
}

type UserRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     Role   `json:"role"`
}

type TopUpRequest struct {
	gorm.Model
	UserId uint  `json:"user_id"`
	Amount int64 `json:"amount"`
}

type UserService interface {
	FindAll() ([]User, error)
	FindById(id uint) (*User, error)
	TopUpBalanceByUserId(userId uint, amount int64) error
	CheckBalance(userId uint) (*int64, error)
	AddUser(req UserRequest) error
}

type userService struct {
	db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{db: db}
}

func (s *userService) FindAll() ([]User, error) {
	var users []User
	err := s.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *userService) FindById(id uint) (*User, error) {
	var user User
	err := s.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *userService) AddUser(req UserRequest) error {
	user := User{Username: req.Username, Password: req.Password, Role: req.Role}
	return s.db.Create(&user).Error
}

func (s *userService) TopUpBalanceByUserId(id uint, amount int64) error {
	if amount <= 0 {
		return fmt.Errorf("amount must be positive")
	}
	return s.db.Transaction(func(tx *gorm.DB) error {
		var user User

		if err := tx.First(&user, id).Error; err != nil {
			return err
		}

		if err := tx.Model(&user).Update("balance", user.Balance+amount).Error; err != nil {
			return err
		}
		return nil
	})
}

func (s *userService) WithdrawBalanceByUserId(id uint, amount int64) error {
	if amount <= 0 {
		fmt.Errorf("amount must be positive")
	}
	return s.db.Transaction(func(tx *gorm.DB) error {
		var user User
		if err := tx.First(&user, id).Error; err != nil {
			return err
		}

		if err := tx.Model(&user).Update("balance", user.Balance-amount).Error; err != nil {
			return err
		}
		return nil
	})
}

func (s *userService) CheckBalance(userId uint) (*int64, error) {
	var user User
	err := s.db.First(&user, userId).Error
	if err != nil {
		return nil, err
	}
	return &user.Balance, nil
}

func (u *User) BeforeSave(tx *gorm.DB) (err error) {
	if u.Password != "" {
		hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
		if err != nil {
			return err
		}
		u.Password = string(hash)
	}
	return nil
}

func (u *User) CheckPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}
