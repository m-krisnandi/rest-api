package services

import (
	"context"
	models "rest-api/bin/modules/employee/models/domain"
	"rest-api/bin/pkg/utils"
)

type QueryService interface {
	GetListEmployees(ctx context.Context) utils.Result
	GetEmployee(ctx context.Context, payload *models.EmployeeRequest) utils.Result
}

type CommandService interface {
	CreateEmployee(ctx context.Context, payload *models.EmployeeRequest) utils.Result
	UpdateEmployee(ctx context.Context, payload *models.EmployeeRequest) utils.Result
	DeleteEmployee(ctx context.Context, payload string) utils.Result
}
