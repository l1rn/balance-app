package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/l1rn/balance-app/services"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(s services.UserService) *UserController {
	return &UserController{userService: s}
}

func (ctrl *UserController) GetAllUsers(ctx *gin.Context) {
	users, err := ctrl.userService.FindAll()
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(200, users)
}

func (ctrl *UserController) TopUpBalance(ctx *gin.Context) {
	var req services.TopUpRequest
	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	err := ctrl.userService.TopUpBalanceByUserId(req.UserId, req.Amount)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err,
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Balance was updated",
	})
}

func (ctrl *UserController) CreateUser(ctx *gin.Context) {
	var input services.UserRequest
	if err := ctx.ShouldBindBodyWithJSON(&input); err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := ctrl.userService.AddUser(input)
	if err != nil {
		ctx.JSON(500, gin.H{
			"error": err,
		})
	}
	ctx.JSON(200, gin.H{"message": "User was created!"})
}

func (ctrl *UserController) CheckBalance(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.ParseUint(idParam, 10, 64)

	if err != nil {
		ctx.JSON(403, gin.H{"error": "Invalid ID Format"})
		return
	}

	balance, err := ctrl.userService.CheckBalance(uint(id))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err})
		return
	}

	ctx.JSON(200, gin.H{"balance": balance})
}
