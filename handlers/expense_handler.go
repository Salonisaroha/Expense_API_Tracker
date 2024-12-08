package handlers

import (
    "Salonisaroha/config"
    "Salonisaroha/models"
    "net/http"
    "github.com/gin-gonic/gin"
)

func AddExpense(c *gin.Context) {
    var expense models.Expense
    if err := c.ShouldBindJSON(&expense); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }
    if result := config.DB.Create(&expense); result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not save expense"})
        return
    }
    c.JSON(http.StatusCreated, expense)
}

func GetExpenses(c *gin.Context) {
    var expenses []models.Expense
    config.DB.Find(&expenses)
    c.JSON(http.StatusOK, expenses)
}

func UpdateExpense(c *gin.Context) {
    id := c.Param("id")
    var expense models.Expense
    if result := config.DB.First(&expense, id); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Expense not found"})
        return
    }
    if err := c.ShouldBindJSON(&expense); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
        return
    }
    config.DB.Save(&expense)
    c.JSON(http.StatusOK, expense)
}

func DeleteExpense(c *gin.Context) {
    id := c.Param("id")
    var expense models.Expense
    if result := config.DB.First(&expense, id); result.Error != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Expense not found"})
        return
    }
    config.DB.Delete(&expense)
    c.JSON(http.StatusOK, gin.H{"message": "Expense deleted successfully"})
}

func GetSummary(c *gin.Context) {
    groupBy := c.Query("group_by")
    var result []map[string]interface{}
    if groupBy == "category" {
        config.DB.Model(&models.Expense{}).Select("category, SUM(amount) as total").Group("category").Scan(&result)
    } else if groupBy == "date" {
        config.DB.Model(&models.Expense{}).Select("DATE(date) as date, SUM(amount) as total").Group("DATE(date)").Scan(&result)
    }
    c.JSON(http.StatusOK, result)
}
