package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ExampleHandler function handler example
func ExampleHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"foo": "bar",
	})
}
