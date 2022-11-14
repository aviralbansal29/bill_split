package controllers

import (
	requestSerializers "github.com/aviralbansal29/bill_split/app/api/request_serializers"
	responseDeserializers "github.com/aviralbansal29/bill_split/app/api/response_deserializers"
	"github.com/aviralbansal29/bill_split/app/api/services"
	"github.com/labstack/echo/v4"
)

// CreateGroup godoc
// @Summary       Creates Group
// @Description   Creates new group and associated user groups
// @Tags          groups
// @Accept        json
// @Produce       json
// @Param         data  body  requestSerializers.CreateGroupForm true  "Body"
// @Success       200
// @Failure       400
// @Failure       404
// @Failure       500
// @Router        /groups [post]
func CreateGroup(ctx echo.Context) error {
	formData, err := requestSerializers.BindCreateGroup(ctx)
	if err != nil {
		return err
	}

	group, err := services.CreateGroup(ctx, formData)
	if err != nil {
		return err
	}

	return responseDeserializers.GetDeserialisedGroup(ctx, group)
}

// AddUser godoc
// @Summary       Add User
// @Description   Adds user to specified group
// @Tags          groups
// @Accept        json
// @Produce       json
// @Param         group_id      path    int     true  "Group ID"
// @Param         user_id       path    int     true  "User ID"
// @Success       200
// @Failure       400
// @Failure       404
// @Failure       500
// @Router        /groups/{group_id}/add-user/{user_id} [post]
func AddUser(ctx echo.Context) error {
	groupID, userID, err := requestSerializers.BindAddUser(ctx)
	if err != nil {
		return err
	}

	group, err := services.AddUser(ctx, userID, groupID)
	if err != nil {
		return err
	}

	return responseDeserializers.GetDeserialisedGroup(ctx, group)
}
