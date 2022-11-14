package responseDeserializers

import (
	"net/http"

	commonUtils "github.com/aviralbansal29/bill_split/app/common/utils"
	"github.com/aviralbansal29/bill_split/app/models"
	"github.com/aviralbansal29/bill_split/config"
	"github.com/labstack/echo/v4"
)

// DeserializedUserFormat defines structure for user details
type DeserializedUserFormat struct {
	Name          string
	PhoneNumber   string              `json:"phone_number"`
	ActiveGroups  []DeserializedGroup `json:"active_groups"`
	SettledGroups []DeserializedGroup `json:"settled_groups"`
}

// GetDeserialisedUser deserializes response format
func GetDeserialisedUser(ctx echo.Context, user models.User) error {
	resp := DeserializedUserFormat{}
	commonUtils.ConvertType(user, &resp)
	config.DatabaseHandler().Model(&models.Group{}).Joins(
		"left join user_groups on groups.id = user_groups.group_id",
	).Where("user_groups.user_id = ? and groups.is_settled = false", user.ID).Find(&resp.ActiveGroups)
	config.DatabaseHandler().Model(&models.Group{}).Joins(
		"left join user_groups on groups.id = user_groups.group_id",
	).Where("user_groups.user_id = ? and groups.is_settled = true", user.ID).Find(&resp.SettledGroups)
	return ctx.JSON(http.StatusCreated, resp)
}
