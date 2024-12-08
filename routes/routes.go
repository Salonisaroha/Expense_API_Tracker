package routes

import (
	"Salonisaroha/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	api := router.Group("/api/v1")
	{
		api.POST("/expenses", handlers.AddExpense)
		api.GET("/expenses", handlers.GetExpenses)
		api.PUT("/expenses/:id", handlers.UpdateExpense)
		api.DELETE("/expenses/:id", handlers.DeleteExpense)
		api.GET("/expenses/summary", handlers.GetSummary)
	}
}
