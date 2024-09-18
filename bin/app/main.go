package main

import (
	"fmt"
	"net/http"
	"rest-api/bin/config"
	"rest-api/bin/pkg/utils"

	employeeHTTP "rest-api/bin/modules/employee/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()
	e.Validator = utils.NewValidationUtil()
	e.Use(middleware.CORS())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "This service is running properly")
	})

	employeeGroup := e.Group("/api")

	employeeHTTP.New().Mount(employeeGroup)

	listenerPort := fmt.Sprintf(":%d", config.GlobalEnv.HTTPPort)
	e.Logger.Fatal(e.Start(listenerPort))
}
