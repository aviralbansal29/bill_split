package requestSerializers

import (
	"net/http"
	"reflect"
	"time"

	commonUtils "github.com/aviralbansal29/bill_split/app/common/utils"
	"github.com/aviralbansal29/bill_split/config"
	"github.com/labstack/echo/v4"
)

type splitMapFormat struct {
	UserID  uint `json:"user_id"`
	Percent uint `validate:"min=0,max=100"`
}

// CreatTransactionForm defines body for POST /transactions API
type CreateTransactionForm struct {
	UserID   uint             `json:"user_id" validate:"required"`
	GroupID  uint             `json:"group_id" validate:"required"`
	Amount   uint             `json:"amount" validate:"required"`
	Name     string           `json:"name" validate:"required"`
	SplitMap []splitMapFormat `json:"split_map" validate:"required"`
}

// BindCreateTransaction binds form data for creating transaction
func BindCreateTransaction(ctx echo.Context) (CreateTransactionForm, error) {
	var validForm CreateTransactionForm
	if err := (&echo.DefaultBinder{}).BindBody(ctx, &validForm); err != nil {
		return validForm, echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{"errors": err})
	}
	if err := config.Validator().Struct(validForm); err != nil {
		errorMap := commonUtils.GetErrorMap(reflect.TypeOf(&validForm), err, "json")
		return validForm, echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{"errors": errorMap})
	}
	return validForm, nil
}

// GetTransactionParam defines params for GET /transactions request
type GetTransactionParam struct {
	UserID   uint      `query:"user_id" json:"user_id" validate:"required"`
	GroupID  uint      `query:"group_id" json:"group_id" validate:"omitempty"`
	FromDate time.Time `query:"from" json:"from" validate:"omitempty"`
	TillDate time.Time `query:"till" json:"till" validate:"omitempty"`
}

// BindGetTransaction binds params for GET /transactions request
func BindGetTransaction(ctx echo.Context) (GetTransactionParam, error) {
	var validForm GetTransactionParam
	if err := (&echo.DefaultBinder{}).BindQueryParams(ctx, &validForm); err != nil {
		return validForm, echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{"errors": err})
	}
	if err := config.Validator().Struct(validForm); err != nil {
		errorMap := commonUtils.GetErrorMap(reflect.TypeOf(&validForm), err)
		return validForm, echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{"errors": errorMap})
	}
	return validForm, nil
}
