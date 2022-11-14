package requestSerializers

import (
	"net/http"
	"reflect"
	"strconv"

	commonUtils "github.com/aviralbansal29/bill_split/app/common/utils"
	"github.com/aviralbansal29/bill_split/config"
	"github.com/labstack/echo/v4"
)

// CreatGroupForm defines parameters for POST /groups API
type CreateGroupForm struct {
	Name    string   `json:"name" validate:"required"`
	Members []string `json:"members" validate:"required"`
}

// BindCreateGroup binds form data for creating group
func BindCreateGroup(ctx echo.Context) (CreateGroupForm, error) {
	var validForm CreateGroupForm
	if err := (&echo.DefaultBinder{}).BindBody(ctx, &validForm); err != nil {
		return validForm, echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{"errors": err})
	}
	if err := config.Validator().Struct(validForm); err != nil {
		errorMap := commonUtils.GetErrorMap(reflect.TypeOf(&validForm), err, "json")
		return validForm, echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{"errors": errorMap})
	}
	return validForm, nil
}

// BindAddUser binds params for adding user to group
func BindAddUser(ctx echo.Context) (uint, uint, error) {
	groupID, err := strconv.ParseUint(ctx.Param("group_id"), 10, 64)
	if err != nil {
		return 0, 0, echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{"errors": "Invalid ID"})
	}
	userID, err := strconv.ParseUint(ctx.Param("user_id"), 10, 64)
	if err != nil {
		return 0, 0, echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{"errors": "Invalid ID"})
	}
	return uint(groupID), uint(userID), nil
}
