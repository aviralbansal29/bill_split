package responseDeserializers

import (
	"net/http"
	"time"

	commonUtils "github.com/aviralbansal29/bill_split/app/common/utils"
	"github.com/aviralbansal29/bill_split/app/models"
	"github.com/labstack/echo/v4"
)

// DeserializedTransaction defines structure for returning transaction
type DeserializedTransaction struct {
	Name      string
	Amount    uint
	GroupName string `json:"group_name"`
}

// GetDeserialisedTransaction deserializes response format
func GetDeserialisedTransaction(ctx echo.Context, transaction models.Transaction) error {
	resp := DeserializedTransaction{
		Name: transaction.Name, Amount: transaction.Amount, GroupName: transaction.Group.Name,
	}
	return ctx.JSON(http.StatusCreated, resp)
}

type DeserializedTransactionList struct {
	Date    time.Time
	Group   string
	Expense string
	Amount  uint
	Pending string
}

func GetDeserialisedTransactionList(ctx echo.Context, dbTransactions []map[string]interface{}) error {
	resp := make([]DeserializedTransactionList, 0)
	commonUtils.ConvertType(dbTransactions, &resp)
	return ctx.JSON(http.StatusOK, resp)
}
