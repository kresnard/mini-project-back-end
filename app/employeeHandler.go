package app

import (
	"mpbe/input"
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

func (ch *EmployeeHandlers) getEmployeeID(c *gin.Context) {

	id := c.Param("id")
	newId, _ := strconv.Atoi(id)
	employees, err := ch.service.GetEmployeeID(newId)

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	c.JSON(http.StatusOK, employees)
}

func (ch *EmployeeHandlers) createEmployee(c *gin.Context) {
	var input input.EmployeeInput
	err := c.ShouldBindJSON(&input)
	employees, _ := ch.service.CreateEmployee(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	c.JSON(http.StatusOK, employees)
}

func (ch *EmployeeHandlers) deleteEmployee(c *gin.Context) {

	id := c.Param("id")
	newId, _ := strconv.Atoi(id)
	employees, err := ch.service.DltEmployee(newId)

	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
	}

	c.JSON(http.StatusOK, employees)
}

func (ch *EmployeeHandlers) updateEmployee(c *gin.Context) {
	id := c.Param("id")
	newId, _ := strconv.Atoi(id)
	var input input.EmployeeInput
	err := c.ShouldBindJSON(&input)
	employees, _ := ch.service.UpdtEmployee(newId, input)
	if err != nil {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	c.JSON(http.StatusOK, employees)
}
