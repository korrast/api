package request

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/gin-gonic/gin"
)

func ParseJSONBody(c *gin.Context, dest interface{}) error {
	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return errors.New("could not read request body: " + err.Error())
	}

	if err := json.Unmarshal(jsonData, dest); err != nil {
		return errors.New("invalid JSON format: " + err.Error())
	}

	return nil
}

func ParseJSONToMap(c *gin.Context) (map[string]string, error) {
	var result map[string]string
	err := ParseJSONBody(c, &result)
	return result, err
}