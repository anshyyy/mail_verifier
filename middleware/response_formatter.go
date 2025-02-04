package middleware

import (
	"log"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error"`
	Success bool        `json:"success"`
}

// Custom writer function that will format the response
func WriteJSON(c *gin.Context, code int, obj gin.H) {
	// Create formatted response

	log.Print(obj)
	response := Response{
		Message: obj["message"].(string),
		Data:    obj["data"],
		Error:   "",
		Success: true,
	}

	// If there's an error field
	if errMsg, exists := obj["error"]; exists && errMsg != nil {
		response.Error = errMsg.(string)
		response.Data = nil
		response.Success = false
	}

	// Send the formatted response
	c.JSON(code, response)
}

func ResponseFormatter() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
