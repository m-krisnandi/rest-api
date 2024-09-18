package services

import (
	"context"
	"rest-api/bin/modules/employee/models/domain"
	"rest-api/bin/modules/employee/repositories/queries"
	"rest-api/bin/pkg/utils"
)

type queryService struct {
	postgreQuery queries.QueryPostgre
}

func NewQueryService(postgreQuery queries.QueryPostgre) *queryService {
	return &queryService{
		postgreQuery: postgreQuery,
	}
}

func (q *queryService) GetListEmployees(ctx context.Context) utils.Result {
	var result utils.Result

	queryPayload := queries.QueryPayload{
		Table:  "employees",
		Select: `*`,
		Output: []models.EmployeeResponse{},
	}

	queryRes := <-q.postgreQuery.FindManyBasic(&queryPayload)
	if queryRes.Error != nil {
		queryRes.Data = []models.EmployeeResponse{}
	}

	result.Data = queryRes.Data
	return result
}

func (q *queryService) GetEmployee(ctx context.Context, payload *models.EmployeeRequest) utils.Result {
	var result utils.Result

	parameter := map[string]interface{}{
		"id": payload.ID,
	}

	queryPayload := queries.QueryPayload{
		Table:  "employees",
		Select: `*`,
		Where:  parameter,
		Output: models.EmployeeResponse{},
	}

	queryRes := <-q.postgreQuery.FindOne(&queryPayload)
	if queryRes.Error != nil {
		queryRes.Data = models.EmployeeResponse{}
	}

	result.Data = queryRes.Data
	return result
}