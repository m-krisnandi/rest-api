package commands

import (
	"rest-api/bin/pkg/utils"

	"gorm.io/gorm"
)

type PostgreCommand struct {
	db *gorm.DB
}

type CommandPayload struct {
	Table string
	Query map[string]interface{}
	Parameter map[string]interface{}
	Document interface{}
	Join string
	Raw string
	Where map[string]interface{}
}

func NewPostgreCommand(db *gorm.DB) *PostgreCommand {
	return &PostgreCommand{
		db: db,
	}
}

func (c *PostgreCommand) InsertOne(payload * CommandPayload) <-chan utils.Result {
	output := make(chan utils.Result)

	go func() {
		defer close(output)

		result := c.db.Table(payload.Table).Create(payload.Document)
		if result.Error != nil {
			output <- utils.Result{Error: result}
		}

		output <- utils.Result{Data: payload.Document}
	}()

	return output
}

func (c *PostgreCommand) Update(payload *CommandPayload) <-chan utils.Result {
	output := make(chan utils.Result)

	go func()  {
		defer close(output)

		result := c.db.Table(payload.Table).Where(payload.Query).Updates(payload.Document)
		if result.Error != nil {
			output <- utils.Result{Error: result}
		}

		output <- utils.Result{Data: payload.Document}
	}()

	return output
}

func (c *PostgreCommand) Delete(payload *CommandPayload) <-chan utils.Result {
	output := make(chan utils.Result)

	go func() {
		defer close(output)

		var data map[string]interface{}
		result := c.db.Table(payload.Table).Where(payload.Query).Delete(payload.Query)
		if result.Error != nil {
			output <- utils.Result{Error: result}
		}

		output <- utils.Result{Data: data}
	}()

	return output
}