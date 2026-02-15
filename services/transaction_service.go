package services

import (
	"time"

	"gorm.io/gorm"
)

type TransactionType string

const (
	TopUp    TransactionType = "top-up"
	Withdraw TransactionType = "withdraw"
)

type TransactionItem struct {
	gorm.Model
	UserID uint
	Amount int64
	Type   TransactionType `gorm:"type:varchar(20)"`
	User   User            `gorm:"foreignKey:UserID"`
}

type TransactionRequest struct {
	UserID uint            `json:"user_id"`
	Amount int64           `json:"amount"`
	Type   TransactionType `json:"type"`
}

type TransactionResponse struct {
	Username  string          `json:"username"`
	Amount    int64           `json:amount`
	Type      TransactionType `json:"type"`
	CreatedAt time.Time       `json:"created_at"`
}

type TransactionService interface {
	GetUserWithHistory(id uint) (UserResponse, error)
}
type transactionService struct {
	db *gorm.DB
}

func NewTransactionService(db *gorm.DB) TransactionService {
	return &transactionService{db: db}
}

func (s *transactionService) AddTransaction(req TransactionRequest) error {
	transaction := TransactionItem{
		UserID: req.UserID,
		Amount: req.Amount,
		Type:   req.Type,
	}

	return s.db.Create(&transaction).Error
}

func (s *transactionService) GetUserWithHistory(id uint) (UserResponse, error) {
	var user User
	err := s.db.Preload("Transactions").First(&user, id).Error
	var tResponse []TransactionResponse

	for _, t := range user.Transactions {
		tResponse = append(tResponse, TransactionResponse{
			Username:  user.Username,
			Amount:    t.Amount,
			Type:      t.Type,
			CreatedAt: t.CreatedAt,
		})
	}
	response := UserResponse{
		ID:           user.ID,
		Username:     user.Username,
		Role:         user.Role,
		Transactions: tResponse,
	}
	return response, err
}
