package rest

import (
	"github.com/ashtishad/banking/internal/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type CustomerHandlers struct {
	Service service.CustomerService
}

// GetCustomerByID is a handler function to get customer by id
func (ch *CustomerHandlers) GetCustomerByID(c *gin.Context) {
	log.Println("Handling GET request on ... /customers/{id}")

	id, err := getCustomerIdFromPath(c.Param("id"))
	if err != nil {
		log.Printf("Error while getting user id : %v", err)
		c.JSON(err.AsStatus(), err)
		return
	}

	customer, err := ch.Service.GetCustomerById(id)
	if err != nil {
		c.JSON(err.AsStatus(), err)
		return
	}

	c.JSON(http.StatusOK, customer)
}

// SearchCustomerByStatus is a handler function to get customer by status
func (ch *CustomerHandlers) SearchCustomerByStatus(c *gin.Context) {
	log.Println("Handling GET request on ... /customers/{status}")

	status, err := validateCustomerStatus(c.Query("status"))

	if err != nil {
		c.JSON(err.AsStatus(), err)
		return
	}

	customers, err := ch.Service.SearchByStatus(status)
	if err != nil {
		c.JSON(err.AsStatus(), err)
		return
	}

	c.JSON(http.StatusOK, customers)
}
