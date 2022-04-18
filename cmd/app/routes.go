package app

import (
	"github.com/ashtishad/banking/internal/rest"
	"github.com/gin-gonic/gin"
)

func getRouteMappings(r *gin.Engine, ch rest.CustomerHandlers) {
	r.GET("/customers/:id", ch.GetCustomerByID)
	r.GET("/customers/search", ch.SearchCustomerByStatus)
	r.POST("/customers", ch.CreateCustomer)
	r.PUT("/customers/:id", ch.UpdateCustomer)
}
