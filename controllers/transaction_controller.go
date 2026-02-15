package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/l1rn/balance-app/services"
)

type TransactionController struct {
	s services.TransactionService
}

func NewTransactionController(s service.TransactionService) *TransactionController {
	return &TransactionController{s: s}
}

func (ctrl *TransactionController) GetUserHistory(ctx *gin.Context) {
	idParam := ctx.Param("id")
	response, err := ctrl.s.GetUserWithHistory()
}
