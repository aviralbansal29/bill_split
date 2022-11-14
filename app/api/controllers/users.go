package controllers

import (
	requestSerializers "github.com/aviralbansal29/bill_split/app/api/request_serializers"
	responseDeserializers "github.com/aviralbansal29/bill_split/app/api/response_deserializers"
	"github.com/aviralbansal29/bill_split/app/api/services"
	"github.com/labstack/echo/v4"
)

// CreateUser godoc
// @Summary       Creates User
// @Description   Creates new user and returns created user
// @Tags          users
// @Accept        json
// @Produce       json
// @Param         data  body  requestSerializers.CreateUserForm true  "Body"
// @Success       200  {object}  responseDeserializers.DeserializedUserFormat
// @Failure       400
// @Failure       404
// @Failure       500
// @Router        /users [post]
func CreateUser(ctx echo.Context) error {
	formData, err := requestSerializers.BindCreateUser(ctx)
	if err != nil {
		return err
	}

	user, err := services.CreateUser(ctx, formData)
	if err != nil {
		return err
	}

	return responseDeserializers.GetDeserialisedUser(ctx, user)
}

// RetrieveUser godoc
// @Summary       Retrieves User
// @Description   Retrieves User data
// @Tags          users
// @Accept        json
// @Produce       json
// @Param         id      path    int     true  "User ID"
// @Success       200  {object}  responseDeserializers.DeserializedUserFormat
// @Success       200
// @Failure       400
// @Failure       404
// @Failure       500
// @Router        /users/{id} [get]
func RetrieveUser(ctx echo.Context) error {
	formData, err := requestSerializers.BindGetUser(ctx)
	if err != nil {
		return err
	}

	user, err := services.RetrieveUser(ctx, formData)
	if err != nil {
		return err
	}

	return responseDeserializers.GetDeserialisedUser(ctx, user)
}
