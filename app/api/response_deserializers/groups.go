package responseDeserializers

import (
	"net/http"

	"github.com/aviralbansal29/bill_split/app/models"
	"github.com/labstack/echo/v4"
)

// DeserializedGroup defines structure for returning group
type DeserializedGroup struct {
	Name      string
	IsSettled bool `json:"is_settled"`
}

// GetDeserialisedGroup deserializes response format
func GetDeserialisedGroup(ctx echo.Context, group models.Group) error {
	return ctx.JSON(http.StatusCreated, group)
}
