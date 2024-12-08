package controller

import (
	"Salonisaroha/models"
	"Salonisaroha/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ExpenseController struct {
	svc services.ExpenseService
}

func NewExpenseController(svc services.ExpenseService) ExpenseController {
	return ExpenseController{
		svc: svc,
	}
}

func (c ExpenseController) AddExpense(ctx *gin.Context) {
	var expense models.Expense
	if err := ctx.ShouldBindJSON(&expense); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := c.svc.AddExpense(expense)
	if err != nil {
		ctx.JSON(500, ErrorResponse("unable to add expense: "+err.Error()))
		return
	}

	ctx.JSON(200, SuccessResponse("Expense Added"))
}

func (c ExpenseController) GetExpenses(ctx *gin.Context) {
	expenses := c.svc.GetExpenses()
	ctx.JSON(http.StatusOK, SuccessResponse(expenses))
}

func (c ExpenseController) UpdateExpense(ctx *gin.Context) {
	id := ctx.Param("id")

	var newExpense models.Expense

	if err := ctx.ShouldBindJSON(&newExpense); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	err := c.svc.UpdateExpense(id, newExpense)
	if err != nil {
		ctx.JSON(500, ErrorResponse("unable to update expense: "+err.Error()))
		return
	}

	ctx.JSON(200, SuccessResponse("Expense Updated"))
}

func ErrorResponse(error string) map[string]string {
	return map[string]string{
		"error": error,
	}
}

func SuccessResponse(data any) map[string]any {
	return map[string]any{
		"data": data,
	}
}
