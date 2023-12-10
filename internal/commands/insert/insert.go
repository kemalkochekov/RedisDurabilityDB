package insert

import (
	"RedisDurabilityDB/internal/datasource"
	"context"
	"errors"
	"fmt"
	"time"
)

const cacheExpiration = 10 * time.Second

type DBInsertCommand struct {
	Name        string
	Description string
	Required    map[string]struct{}
	db          datasource.Datasource
}

func NewInsert(db datasource.Datasource) *DBInsertCommand {
	return &DBInsertCommand{
		Name:        "insert",
		Description: "insert --table: [string] --keyID: [string] --value: [string]\n",
		Required:    map[string]struct{}{"table:": {}, "keyID:": {}, "value:": {}},
		db:          db,
	}
}

func (cmd *DBInsertCommand) GetName() string {
	return cmd.Name
}

func (cmd *DBInsertCommand) GetDescription() string {
	return cmd.Description
}

func (cmd *DBInsertCommand) GetRequired() map[string]struct{} {
	return cmd.Required
}

func (cmd *DBInsertCommand) DoAction(args map[string]string) error {
	keyDB := ""

	for key, val := range args {
		if key == "value:" {
			continue
		}
		keyDB += fmt.Sprintf("%s %s, ", key, val)
	}

	valDB, ok := args["value:"]
	if !ok {
		return errors.New("not found")
	}

	err := cmd.db.Set(context.Background(), keyDB, valDB, cacheExpiration)
	if err != nil {
		return err
	}

	return nil
}
