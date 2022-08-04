package app

import (
	"mpbe/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type EmployeeHandlers struct {
	service service.EmployeeService
}

func (ch *EmployeeHandlers) getAllEmployee(c *gin.Context) {

	// status := r.URL.Query().Get("status")

	employees, err := ch.service.GetAllEmployee()

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}

	c.JSON(http.StatusOK, employees)

}

func (ch *EmployeeHandlers) GetEmployeeID(c *gin.Context) {

	id := c.Param("id")
	newId, _ := strconv.Atoi(id)
	employees, err := ch.service.GetEmployeeID(newId)

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	c.JSON(http.StatusOK, employees)
}
