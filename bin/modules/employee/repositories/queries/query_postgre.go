package queries

import (
	"rest-api/bin/pkg/utils"

	"gorm.io/gorm"
)

type PostgreQuery struct {
	db    *gorm.DB
	table string
}

type QueryPayload struct {
	Select    string
	Query     string
	Parameter map[string]interface{}
	Where     map[string]interface{}
	Offset    int
	Limit     int
	Join      string
	Group     string
	Table     string
	Order     string
	WhereRaw  string
	Output    interface{}
}

func NewPostgreQuery(db *gorm.DB) *PostgreQuery {
	return &PostgreQuery{
		db: db,
	}
}

func (c *PostgreQuery) FindOne(payload *QueryPayload) <-chan utils.Result {
	output := make(chan utils.Result)

	go func() {
		defer close(output)

		var data map[string]interface{}
		result := c.db.Table(payload.Table).Select(payload.Select).Where(payload.Where).Joins(payload.Join).
			Order(payload.Order).Limit(1).Find(&data)
		if result.Error != nil || result.RowsAffected == 0 {
			output <- utils.Result{
				Error: "Data Not Found",
			}
		}

		output <- utils.Result{Data: data}
	}()

	return output
}

func (c *PostgreQuery) FindMany(payload *QueryPayload) <-chan utils.Result {
	output := make(chan utils.Result)

	go func() {
		defer close(output)

		result := c.db.Table(payload.Table).Select(payload.Select).Where(payload.Where).Where(payload.WhereRaw).
			Joins(payload.Join).Order(payload.Order).Group(payload.Group).Find(&payload.Output)
		if result.Error != nil || result.RowsAffected == 0 {
			output <- utils.Result{
				Error: "Data Not Found",
			}
		}

		output <- utils.Result{Data: payload.Output}
	}()

	return output
}

func (c *PostgreQuery) FindManyBasic(payload *QueryPayload) <-chan utils.Result {
	output := make(chan utils.Result)

	go func() {
		defer close(output)

		result := c.db.Table(payload.Table).Select(payload.Select).Where(payload.Where).
			Joins(payload.Join).Find(&payload.Output)
		if result.Error != nil || result.RowsAffected == 0 {
			output <- utils.Result{
				Error: "Data Not Found",
			}
		}

		output <- utils.Result{Data: payload.Output}
	}()

	return output
}
