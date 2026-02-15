package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/l1rn/balance-app/services"
)

type TransactionController struct {
	s services.TransactionService
}

func NewTransactionController(s services.TransactionService) *TransactionController {
	return &TransactionController{s: s}
}

func (ctrl *TransactionController) GetUserHistory(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)
	if err != nil {
		ctx.JSON(403, gin.H{
			"error": "Invalid ID Format",
		})
		return
	}
	response, err := ctrl.s.GetUserWithHistory(uint(id))
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": "Failed to get user history",
		})
		return
	}
	ctx.JSON(200, response)
}
