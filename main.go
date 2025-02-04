package main

import (
	"log"
	"os"

	//"github.com/anshyyy/mail_verifier/middleware"
	"github.com/anshyyy/mail_verifier/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)



func main() {
	gin.SetMode(gin.ReleaseMode)
	log.Println("Starting the application...")
	r := gin.Default()
	// r.Use(middleware.ResponseFormatter())
	routes.InitalizeRoutes(r)
	if err := godotenv.Load(); err != nil {
		log.Panic("Error loading .env file")
	}
	port := os.Getenv("PORT")
	func() {
		if err := r.Run(":" + port); err != nil {
			log.Fatal("Error occured starting the server", err)
		}
	}()
}
