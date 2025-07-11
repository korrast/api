package api

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"main/database"
	"main/model"
)

func createTable(c *gin.Context) {
	var newTable model.Table

	rawUserId, err := getUserIdFromContext(c)
	if err != nil {
		msg := gin.H{
			"error": err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, msg)
	}

	userId, ok := rawUserId.(uuid.UUID)
	if !ok {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "invalid user ID in context"})
		return
	}

	jsonData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		msg := gin.H{
			"error": "Could not read request body " + err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, msg)
		return
	}

	var d map[string]string
	if err := json.Unmarshal(jsonData, &d); err != nil {
		msg := gin.H{
			"error": "invalid JSON format " + err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, msg)
		return
	}

	var title string = d["title"]
	newTable.Init(title)

	conn, err := database.StartConn()
	if err != nil {
		msg := gin.H{
			"error": "can't start db : " + err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, msg)
		return
	}
	defer database.CloseConn(conn)

	if err := database.InsertTable(conn, newTable); err != nil {
		msg := gin.H{
			"error": "can't insert in database : " + err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, msg)
		return
	}

	if err := database.InsertLinkUserTable(conn, userId.String(), newTable.Id.String()); err != nil {
		msg := gin.H{
			"errors": "can't insert in database : " + err.Error(),
		}
		c.AbortWithStatusJSON(http.StatusInternalServerError, msg)
		return
	}

	c.JSON(http.StatusCreated, "New table created")
}
