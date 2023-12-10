package get

import (
	"RedisDurabilityDB/internal/datasource"
	"context"
	"fmt"
)

type DBGetCommand struct {
	Name        string
	Description string
	Required    map[string]struct{}
	db          datasource.Datasource
}

func NewGet(db datasource.Datasource) *DBGetCommand {
	return &DBGetCommand{
		Name:        "get",
		Description: "get --table: [string] --keyID: [string]\n",
		Required:    map[string]struct{}{"table:": {}, "keyID:": {}},
		db:          db,
	}
}

func (cmd *DBGetCommand) GetName() string {
	return cmd.Name
}

func (cmd *DBGetCommand) GetDescription() string {
	return cmd.Description
}

func (cmd *DBGetCommand) GetRequired() map[string]struct{} {
	return cmd.Required
}

func (cmd *DBGetCommand) DoAction(args map[string]string) error {
	keyDB := ""
	for key, val := range args {
		keyDB += fmt.Sprintf("%s %s, ", key, val)
	}

	got, err := cmd.db.Get(context.Background(), keyDB)
	if err != nil {
		return err
	}

	fmt.Println(got)

	return nil
}
