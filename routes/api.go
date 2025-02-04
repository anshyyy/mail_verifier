package routes

import (
	v1 "github.com/anshyyy/mail_verifier/routes/v1"
	"github.com/gin-gonic/gin"
)

func InitalizeRoutes(r *gin.Engine) {
	v1Group := r.Group("/api/v1")
	{
		v1.InitalizeMailingRoutes(v1Group)
	}
}