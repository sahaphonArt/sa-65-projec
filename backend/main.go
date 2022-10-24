package main

import (
	"github.com/sahaphonArt/sa-65-project/controller"
	"github.com/sahaphonArt/sa-65-project/entity"

	"github.com/gin-gonic/gin"
)

const PORT = "8080"

func main() {
	entity.SetupDatabase()

	r := gin.Default()
	r.Use(CORSMiddleware())
	router := r.Group("/")
	{
		{
			//Admin Router
			router.GET("/admins", controller.ListAdmin)
			router.GET("/admin/:id", controller.GetAdmin)
			router.PATCH("/admins", controller.UpdateAdmin)
			router.DELETE("/admin/:id", controller.DeleteAdmin)

			//Organization Router
			router.GET("/Organizations", controller.ListOrganization)
			router.GET("/Organization/:id", controller.GetOrganization)
			router.POST("/Organizations", controller.CreateOrganization)
			router.PATCH("/Organizations", controller.UpdateOrganization)
			router.DELETE("/Organization/:id", controller.DeleteOrganization)

			//TypeFund Router
			router.GET("/TypeFunds", controller.ListTypeFund)
			router.GET("/TypeFund/:id", controller.GetTypeFund)
			router.POST("/TypeFunds", controller.CreateTypeFund)
			router.PATCH("/TypeFunds", controller.UpdateTypeFund)
			router.DELETE("/TypeFund/:id", controller.DeleteTypeFund)

			//Donator Router
			router.GET("/donators", controller.ListDonators)
			router.GET("/donator/:id", controller.GetDonator)
			router.POST("/donators", controller.CreateDonator)
			router.PATCH("/donators", controller.UpdateDonator)
			router.DELETE("/donator/:id", controller.CreateDonator)

		}
	}

	// Signup User Route
	// r.POST("/signup", controller.CreateAdmins)
	// // login User Route
	// r.POST("/login", controller.Login)

	// Run the server go run main.go
	r.Run("localhost: " + PORT)

}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
