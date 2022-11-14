package requestSerializers

import (
	"net/http"
	"reflect"
	"strconv"

	commonUtils "github.com/aviralbansal29/bill_split/app/common/utils"
	"github.com/aviralbansal29/bill_split/config"
	"github.com/labstack/echo/v4"
)

// CreateUserForm defines parameters for POST /users API
type CreateUserForm struct {
	Name        string `json:"name" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required,numeric"`
}

// BindCreateUser binds form data for creating user
func BindCreateUser(ctx echo.Context) (CreateUserForm, error) {
	var validForm CreateUserForm
	if err := (&echo.DefaultBinder{}).BindBody(ctx, &validForm); err != nil {
		return validForm, echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{"errors": err})
	}
	if err := config.Validator().Struct(validForm); err != nil {
		errorMap := commonUtils.GetErrorMap(reflect.TypeOf(&validForm), err, "json")
		return validForm, echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{"errors": errorMap})
	}
	return validForm, nil
}

// BindGetUser defines parameters for getting user details
func BindGetUser(ctx echo.Context) (uint, error) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return uint(id), echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{"errors": "Invalid ID"})
	}
	return uint(id), nil
}
