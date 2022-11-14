package controllers

import (
	requestSerializers "github.com/aviralbansal29/bill_split/app/api/request_serializers"
	responseDeserializers "github.com/aviralbansal29/bill_split/app/api/response_deserializers"
	"github.com/aviralbansal29/bill_split/app/api/services"
	"github.com/labstack/echo/v4"
)

// CreateTransaction godoc
// @Summary       Creates Transaction
// @Description   Creates new Transaction and related user transactions
// @Tags          transactions
// @Accept        json
// @Produce       json
// @Param         data  body  requestSerializers.CreateTransactionForm true  "Body"
// @Success       200  {object}  responseDeserializers.DeserializedTransaction
// @Failure       400
// @Failure       404
// @Failure       500
// @Router        /transactions [post]
func CreateTransaction(ctx echo.Context) error {
	formData, err := requestSerializers.BindCreateTransaction(ctx)
	if err != nil {
		return err
	}

	transaction, err := services.CreateTransaction(ctx, formData)
	if err != nil {
		return err
	}

	return responseDeserializers.GetDeserialisedTransaction(ctx, transaction)
}

// GetTransactions godoc
// @Summary       Get Transactions
// @Description   Gets list of transactions
// @Tags          transactions
// @Accept        json
// @Produce       json
// @Param         user_id      query    int     true  "User ID"
// @Param         group_id     query    int     false  "Group ID"
// @Param         from     query    time.Time     false  "From"
// @Param         till     query    time.Time     false  "Till"
// @Success       200  {object}  responseDeserializers.DeserializedTransactionList
// @Failure       400
// @Failure       404
// @Failure       500
// @Router        /transactions [get]
func GetTransaction(ctx echo.Context) error {
	params, err := requestSerializers.BindGetTransaction(ctx)
	if err != nil {
		return err
	}

	transactions, err := services.GetTransactions(ctx, params)
	if err != nil {
		return err
	}

	return responseDeserializers.GetDeserialisedTransactionList(ctx, transactions)
}
