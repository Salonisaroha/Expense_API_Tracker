package api

import (
	controller "Salonisaroha/controllers"
	"Salonisaroha/repository"
	"Salonisaroha/services"

	"github.com/gin-gonic/gin"
)

func StartServer() {
	router := gin.Default()

	// Expenses API
	api := router.Group("/api/v1/expense")
	{
		repo := repository.NewExpenseRepository()
		svc := services.NewExpenseSvc(repo)
		expenseController := controller.NewExpenseController(svc)

		api.POST("/expenses", expenseController.AddExpense)
		api.GET("/expenses", expenseController.GetExpenses)
		api.PUT("/expenses/:id", expenseController.UpdateExpense)
	}

}
