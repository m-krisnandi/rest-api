package commands

import "rest-api/bin/pkg/utils"

type CommandPostgre interface {
	InsertOne(payload *CommandPayload) <-chan utils.Result
	Update(payload *CommandPayload) <-chan utils.Result
	Delete(payload *CommandPayload) <-chan utils.Result
}