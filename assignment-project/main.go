package main

import (
	"assignment-project/controllers"
	"assignment-project/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var PORT = ":8080"

func main() {
	r := gin.Default()
	models.ConnectDatabase()

	r.GET("/orders", controllers.AllOrders)
	r.POST("/orders", controllers.CreateOrder)
	r.GET("orders/:id", controllers.OrderById)
	r.PUT("orders/:id", controllers.UpdateOrder)
	r.DELETE("orders/:id", controllers.DeleteOrder)

	r.Run()
	fmt.Println("Application is listening on port", PORT)
	http.ListenAndServe(PORT, nil)
}
