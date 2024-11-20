package handlers

import (
	"api-gateway/proto"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handlers struct {
	employeeClient proto.EmployeeServiceClient
}

func NewHandler(employeeClient proto.EmployeeServiceClient) *Handlers {
	return &Handlers{employeeClient: employeeClient}
}

func (h *Handlers) AddEmployee(c *gin.Context) {
	var AddEmployeeRequest proto.AddEmployeeRequest
	if err := c.BindJSON(&AddEmployeeRequest); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"gw_handler: add employee: bind json": err.Error()})
		return
	}

	id, err := h.employeeClient.AddEmployee(context.Background(), &AddEmployeeRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"gw_handler: add employee: client": err.Error()})
		return
	}

	c.JSON(http.StatusOK, id)
}

func (h *Handlers) RemoveEmployee(c *gin.Context) {
	var removeRequest *proto.DeleteEmployeeRequest
	if err := c.BindJSON(&removeRequest); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"gw_handlers: remove employee: bind: ": err.Error()})
		return
	}

	success, err := h.employeeClient.DeleteEmployee(context.Background(), removeRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"gw_handlers: remove employee: client: ": err.Error()})
		return
	}

	c.JSON(http.StatusOK, success)
}

func (h *Handlers) GetEmployees(c *gin.Context) {
	var companyRequest *proto.CompanyEmployeesRequest

	if err := c.BindJSON(&companyRequest); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"gw_handlers: get employee: bind: ": err.Error()})
		return
	}

	companyResponse, err := h.employeeClient.ShowCompanyEmployees(context.Background(), companyRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"gw_handlers: get employee: show company": err.Error()})
		return
	}

	c.JSON(http.StatusOK, companyResponse)
}

func (h *Handlers) UpdateEmployee(c *gin.Context) {
	var updateRequest *proto.UpdateEmployeeRequest
	if err := c.BindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusBadRequest, map[string]interface{}{"gw_handlers: update employee: bind: ": err.Error()})
		return
	}

	success, err := h.employeeClient.UpdateEmployee(context.Background(), updateRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, map[string]interface{}{"gw_handlers: update employee: client:": err.Error()})
		return
	}

	c.JSON(http.StatusOK, success)
}
