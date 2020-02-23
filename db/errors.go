package db

type ErrBudgetNotFound struct {}

func (e *ErrBudgetNotFound) Error() string {
	return "Budget with given name not found"
}