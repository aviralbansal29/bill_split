package services

import (
	"net/http"

	requestSerializers "github.com/aviralbansal29/bill_split/app/api/request_serializers"
	"github.com/aviralbansal29/bill_split/app/models"
	"github.com/aviralbansal29/bill_split/config"
	"github.com/labstack/echo/v4"
)

// CreateTransaction creates transaction and related user transactions
func CreateTransaction(ctx echo.Context, formData requestSerializers.CreateTransactionForm) (models.Transaction, error) {
	transaction := models.Transaction{GroupID: formData.GroupID, Amount: formData.Amount, Name: formData.Name}
	if err := validate100Percent(ctx, formData); err != nil {
		return transaction, err
	}
	config.DatabaseHandler().Create(&transaction)
	createUserTransactions(ctx, transaction, formData)
	return transaction, nil
}

func GetTransactions(ctx echo.Context, params requestSerializers.GetTransactionParam) ([]map[string]interface{}, error) {
	transactions := make([]map[string]interface{}, 0)
	t := models.Transaction{}

	query := t.GetListQuery(params.UserID)
	if params.GroupID != 0 {
		query = t.AddGroupFilter(query, params.GroupID)
	}

	if !params.FromDate.IsZero() && !params.TillDate.IsZero() {
		query = t.AddTimePeriodFilter(query, params.FromDate, params.TillDate)
	}

	query.Find(&transactions)
	return transactions, nil
}

func validate100Percent(ctx echo.Context, formData requestSerializers.CreateTransactionForm) error {
	totalPercent := uint(0)
	for _, entry := range formData.SplitMap {
		totalPercent += entry.Percent
	}
	if totalPercent != 100 {
		return echo.NewHTTPError(http.StatusBadRequest, "Split not reached 100%")
	}
	return nil
}

func createUserTransactions(
	ctx echo.Context, transaction models.Transaction, formData requestSerializers.CreateTransactionForm,
) {
	userTransactions := make([]models.UserTransaction, 0)
	for _, entry := range formData.SplitMap {
		if entry.UserID != formData.UserID {
			userTransactions = append(userTransactions, models.UserTransaction{
				TransactionID: transaction.ID, FromID: entry.UserID, ToID: formData.UserID,
				Amount: entry.Percent * formData.Amount / 100,
			})
		}
	}
	models.BulkCreateUserTransaction(userTransactions)
}
