package services

import (
	"context"
	models "rest-api/bin/modules/employee/models/domain"
	"rest-api/bin/modules/employee/repositories/commands"
	"rest-api/bin/modules/employee/repositories/queries"
	httpError "rest-api/bin/pkg/http-error"
	"rest-api/bin/pkg/utils"

	"github.com/google/uuid"
)

type commandService struct {
	postgreQuery   queries.QueryPostgre
	postgreCommand commands.CommandPostgre
}

func NewCommandService(postgreQuery queries.QueryPostgre, postgreCommand commands.CommandPostgre) *commandService {
	return &commandService{
		postgreQuery:   postgreQuery,
		postgreCommand: postgreCommand,
	}
}

func (u commandService) CreateEmployee(ctx context.Context, payload *models.EmployeeRequest) utils.Result {
	var result utils.Result

	commandPayload := commands.CommandPayload{
		Table: "employees",
		Document: map[string]interface{}{
			"id":        uuid.NewString(),
			"name":      payload.Name,
			"age":       payload.Age,
			"job_title": payload.JobTitle,
			"company":   payload.Company,
		},
	}

	insertEmployee := <-u.postgreCommand.InsertOne(&commandPayload)
	if insertEmployee.Error != nil {
		errObj := httpError.NewInternalServerError()
		errObj.Message = "Gagal menambahkan karyawan"
		result.Error = errObj
		return result
	}

	result.Data = insertEmployee.Data

	return result
}

func (u commandService) UpdateEmployee(ctx context.Context, payload *models.EmployeeRequest) utils.Result {
	var result utils.Result

	queryRes := <-u.postgreQuery.FindOne(&queries.QueryPayload{
		Table:  "employees",
		Select: "id",
		Where: map[string]interface{}{
			"id": payload.ID,
		},
	})

	if queryRes.Error != nil {
		errObj := httpError.NewNotFound()
		errObj.Message = "Data karyawan tidak ditemukan"
		result.Error = errObj
		return result
	}

	data := queryRes.Data.(map[string]interface{})

	commandPayload := commands.CommandPayload{
		Table: "employees",
		Query: map[string]interface{}{
			"id": payload.ID,
		},
		Document: map[string]interface{}{
			"name":      payload.Name,
			"age":       payload.Age,
			"job_title": payload.JobTitle,
			"company":   payload.Company,
		},
	}

	updateEmployee := <-u.postgreCommand.Update(&commandPayload)
	if updateEmployee.Error != nil {
		errObj := httpError.NewInternalServerError()
		errObj.Message = "Gagal mengubah data karyawan"
		result.Error = errObj
		return result
	}

	result.Data = map[string]interface{}{
		"id":        data["id"],
		"name":      payload.Name,
		"age":       payload.Age,
		"job_title": payload.JobTitle,
		"company":   payload.Company,
	}

	return result
}

func (u commandService) DeleteEmployee(ctx context.Context, payload string) utils.Result {
	var result utils.Result

	queryRes := <-u.postgreQuery.FindOne(&queries.QueryPayload{
		Table: "employees",
		Where: map[string]interface{}{
			"id": payload,
		},
	})

	if queryRes.Error != nil {
		errObj := httpError.NewNotFound()
		errObj.Message = "Data karyawan tidak ditemukan"
		result.Error = errObj
		return result
	}

	commandPayload := commands.CommandPayload{
		Table: "employees",
		Query: map[string]interface{}{
			"id": payload,
		},
	}

	deleteEmployee := <-u.postgreCommand.Delete(&commandPayload)

	if deleteEmployee.Error != nil {
		errObj := httpError.NewInternalServerError()
		errObj.Message = "Gagal menghapus data karyawan"
		result.Error = errObj
		return result
	}

	result.Data = map[string]interface{}{}
	return result
}