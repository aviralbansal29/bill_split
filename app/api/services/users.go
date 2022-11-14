package services

import (
	"net/http"

	requestSerializers "github.com/aviralbansal29/bill_split/app/api/request_serializers"
	"github.com/aviralbansal29/bill_split/app/models"
	"github.com/aviralbansal29/bill_split/config"
	"github.com/labstack/echo/v4"
)

// CreateUser is the main logic for creating data
func CreateUser(ctx echo.Context, userForm requestSerializers.CreateUserForm) (models.User, error) {
	user := models.User{
		Name: userForm.Name, PhoneNumber: userForm.PhoneNumber,
	}
	err := user.CreateEntry()
	if err != nil {
		return user, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return user, nil
}

// RetrieveUser retrieves the user info
func RetrieveUser(ctx echo.Context, id uint) (models.User, error) {
	user := models.User{}
	err := config.DatabaseHandler().First(&user, id).Error
	if user.ID == 0 {
		return user, echo.NewHTTPError(http.StatusBadRequest, "Invalid User ID")
	}
	return user, err
}
