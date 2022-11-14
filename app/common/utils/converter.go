package commonUtils

import (
	"encoding/json"
)

// ConvertType Converts one data type to another using JSON
func ConvertType(src interface{}, des interface{}) error {
	tempJSON, err := json.Marshal(src)
	if err != nil {
		return err
	}
	err = json.Unmarshal(tempJSON, des)
	return err
}
