package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	models "rest-api/bin/modules/employee/models/domain"
	"rest-api/bin/modules/employee/repositories/commands"
	"rest-api/bin/modules/employee/repositories/queries"
	"rest-api/bin/modules/employee/services"
	"rest-api/bin/pkg/database"
	"rest-api/bin/pkg/utils"

	"github.com/labstack/echo/v4"
)

// HTTPHandler struct
type HTTPHandler struct {
	queryService   services.QueryService
	commandService services.CommandService
}

// New initiation
func New() *HTTPHandler {
	postgreDb := database.InitPostgre(context.Background())

	postgreQuery := queries.NewPostgreQuery(postgreDb)
	queryService := services.NewQueryService(postgreQuery)

	postgreCommand := commands.NewPostgreCommand(postgreDb)
	commandService := services.NewCommandService(postgreQuery, postgreCommand)

	return &HTTPHandler{
		queryService:   queryService,
		commandService: commandService,
	}
}

// Mount function
func (u *HTTPHandler) Mount(g *echo.Group) {
	g.GET("/v1/employees", u.getListEmployees)
	g.GET("/v1/employees/:id", u.getEmployees)
	g.POST("/v1/employees", u.createEmployee)
	g.PUT("/v1/employees/:id", u.updateEmployee)
	g.DELETE("/v1/employees/:id", u.deleteEmployee)
}

func (u *HTTPHandler) getListEmployees(c echo.Context) error {
	result := u.queryService.GetListEmployees(c.Request().Context())
	if result.Error != nil {
		return utils.ResponseError(result.Error, c)
	}

	return utils.ResponseData(result.Data, http.StatusOK, c)
}

func (u *HTTPHandler) getEmployees(c echo.Context) error {
	query := make(map[string]interface{})
	for key, value := range c.QueryParams() {
		query[key] = value[0]
	}

	payload, _ := json.Marshal(query)
	var data models.EmployeeRequest
	json.Unmarshal(payload, &data)

	data.ID = c.Param("id")

	result := u.queryService.GetEmployee(c.Request().Context(), &data)
	if result.Error != nil {
		return utils.ResponseError(result.Error, c)
	}

	return utils.ResponseData(result.Data, http.StatusOK, c)
}

func (u *HTTPHandler) createEmployee(c echo.Context) error {
	payload := new(models.EmployeeRequest)
	if err := utils.BindValidate(c, payload); err != nil {
		return utils.ResponseData(map[string]string{"error": err.Error()}, http.StatusBadRequest, c)
	}

	result := u.commandService.CreateEmployee(c.Request().Context(), payload)
	if result.Error != nil {
		return utils.ResponseError(result.Error, c)
	}

	return utils.ResponseData(result.Data, http.StatusCreated, c)
}

func (u *HTTPHandler) updateEmployee(c echo.Context) error {
	payload := new(models.EmployeeRequest)

	if err := utils.BindValidate(c, payload); err != nil {
		return utils.ResponseData(map[string]string{"error": err.Error()}, http.StatusBadRequest, c)
	}

	payload.ID = c.Param("id")
	
	result := u.commandService.UpdateEmployee(c.Request().Context(), payload)

	if result.Error != nil {
		return utils.ResponseError(result.Error, c)
	}

	return utils.ResponseData(result.Data, http.StatusOK, c)
}

func (u *HTTPHandler) deleteEmployee(c echo.Context) error {
	result := u.commandService.DeleteEmployee(c.Request().Context(), c.Param("id"))

	if result.Error != nil {
		return utils.ResponseError(result.Error, c)
	}

	return utils.ResponseData(result.Data, http.StatusOK, c)
}