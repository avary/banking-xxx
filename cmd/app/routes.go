package app

import (
	"github.com/ashtishad/banking/internal/rest"
	"github.com/gin-gonic/gin"
)

func getRouteMappings(r *gin.Engine, ch rest.CustomerHandlers) {
	r.GET("/customers/:id", ch.GetCustomerByID)
}
