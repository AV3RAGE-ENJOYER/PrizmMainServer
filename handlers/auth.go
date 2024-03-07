package handlers

import (
	"server/database"

	"github.com/gin-gonic/gin"
)

func CheckUserPassword(db database.Database) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data map[string]string
		c.ShouldBindJSON(&data)

		username := data["username"]
		passwordHash := data["passwordHash"]

		if !db.CheckPassword(username, passwordHash) {
			c.IndentedJSON(401, gin.H{"status": "not authenticated"})
			return
		}

		c.JSON(200, gin.H{"status": "authenticated"})
	}
}
