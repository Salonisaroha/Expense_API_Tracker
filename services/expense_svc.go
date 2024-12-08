package services

import (
	"Salonisaroha/models"
	"Salonisaroha/repository"
	"errors"
)

type ExpenseService struct {
	expenseRepo repository.ExpenseRepository
}

func NewExpenseSvc(expenseRepo repository.ExpenseRepository) ExpenseService {
	return ExpenseService{
		expenseRepo: expenseRepo,
	}
}

func (svc ExpenseService) AddExpense(expense models.Expense) error {
	return svc.expenseRepo.AddExpense(expense)
}

func (svc ExpenseService) GetExpenses() []models.Expense {
	return svc.expenseRepo.GetExpenses()
}

func (svc ExpenseService) UpdateExpense(id string, newExpense models.Expense) error {
	expense, err := svc.expenseRepo.GetOneExpense(id)
	if err != nil {
		return errors.New("unable to get expense: " + err.Error())
	}

	expense = newExpense
	svc.expenseRepo.SaveExpense(&expense)

	return nil
}
