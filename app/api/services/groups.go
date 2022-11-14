package services

import (
	"fmt"
	"net/http"
	"strconv"

	requestSerializers "github.com/aviralbansal29/bill_split/app/api/request_serializers"
	"github.com/aviralbansal29/bill_split/app/models"
	"github.com/aviralbansal29/bill_split/config"
	"github.com/labstack/echo/v4"
)

// CreateGroup is the main logic for creating data
func CreateGroup(ctx echo.Context, formData requestSerializers.CreateGroupForm) (models.Group, error) {
	group := models.Group{Name: formData.Name}

	var count int64
	config.DatabaseHandler().Model(&models.User{}).Where("id IN ?", formData.Members).Count(&count)
	if count != int64(len(formData.Members)) {
		return group, echo.NewHTTPError(http.StatusBadRequest, "Invalid member ids present")
	}

	err := group.CreateEntry()
	if err != nil {
		return group, echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = createUserGroups(ctx, group, formData.Members)
	if err != nil {
		group.DeleteEntry()
	}
	return group, err
}

// AddUser adds user to specified group
func AddUser(ctx echo.Context, userID uint, groupID uint) (models.Group, error) {
	group := models.Group{}
	config.DatabaseHandler().First(&group, groupID)
	err := createUserGroups(ctx, group, []string{fmt.Sprint(userID)})
	return group, err

}

func createUserGroups(ctx echo.Context, group models.Group, members []string) error {
	groupMembers := make([]models.UserGroup, 0)
	for _, member := range members {
		memberID, err := strconv.ParseUint(member, 10, 64)
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Invalid member ids present")
		}
		groupMembers = append(groupMembers, models.UserGroup{
			GroupID: group.ID, UserID: uint(memberID),
		})
	}
	return models.BulkCreateGroups(groupMembers)
}
