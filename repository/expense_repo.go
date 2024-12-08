package repository

import (
	db "Salonisaroha/database"
	"Salonisaroha/models"
	"errors"

	"gorm.io/gorm"
)

type ExpenseRepository struct {
	db *gorm.DB
}

func NewExpenseRepository() ExpenseRepository {
	return ExpenseRepository{
		db: db.Connection,
	}
}

func (repo ExpenseRepository) AddExpense(expense models.Expense) error {
	if result := repo.db.Create(&expense); result.Error != nil {
		return errors.New("unable to create expense: " + result.Error.Error())
	}
	return nil
}

func (repo ExpenseRepository) GetExpenses() []models.Expense {
	var expenses []models.Expense
	repo.db.Find(&expenses)

	return expenses
}

func (repo ExpenseRepository) GetOneExpense(id string) (models.Expense, error) {
	var expense models.Expense

	if result := repo.db.First(&expense, id); result.Error != nil {
		return models.Expense{}, result.Error
	}

	return expense, nil
}

func (repo ExpenseRepository) SaveExpense(expense *models.Expense) {
	repo.db.Save(&expense)
}
