package v1

import (
	"log"
	"net/http"

	"github.com/anshyyy/mail_verifier/services/mailverifier"
	"github.com/gin-gonic/gin"
)

func InitalizeMailingRoutes(r *gin.RouterGroup) {
	mailGroup := r.Group("/")
	{
		mailGroup.GET("/", func(ctx *gin.Context) {

			ctx.JSON(http.StatusOK, gin.H{
				"message": "Success",
				"data":    "yo!!",
				"success": true,
			})
		})

		mailGroup.POST("/verify", func(ctx *gin.Context) {
			type RequestBody struct {
				Email    string `json:"email" binding:"required,email"`
				UseProxy bool   `json:"use_proxy"`
			}

			var reqBody RequestBody
			

			if err := ctx.ShouldBindJSON(&reqBody); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{
					"data":    nil,
					"success": false,
					"error":   err.Error(),
				})
				return
			}
			log.Println("req body ", reqBody)

			result, err := mailService.VerifySingleMail(reqBody.Email, reqBody.UseProxy)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"data":    nil,
					"success": false,
					"error":   err.Error()})
				return
			}

			ctx.JSON(http.StatusOK, gin.H{"result": result})
		})

		mailGroup.POST("/verify-bulk", func(ctx *gin.Context) {
			type RequestBody struct {
				Emails   []string `json:"emails" binding:"required,dive,email"`
				UseProxy bool     `json:"use_proxy"`
			}

			var reqBody RequestBody
			if err := ctx.ShouldBindJSON(&reqBody); err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			results, err := mailService.VerifyBulkMail(reqBody.Emails, reqBody.UseProxy)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			ctx.JSON(http.StatusOK, gin.H{"results": results})
		})

	}
}
