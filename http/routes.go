package http

import (
	"github.com/banwire/api-exam/handlers"
	"github.com/gin-gonic/gin"
)

// Router define all routes http
func Router(router *gin.Engine) {

	v1 := router.Group("/comerce")
	{
		v1.POST("/addOne", handlers.AddnewComerce)
		v1.POST("/addMany", handlers.AddnewComerces)
		v1.PUT("/updateOne/:id", handlers.UpdateComerce)
	}

	tr := router.Group("/tran")
	{
		tr.POST("/addOne/:id", handlers.AddTransaccion)
		tr.GET("/profits", handlers.GetProfits)
		tr.GET("/profits/:id", handlers.GetEspecificProfit)
	}
}
