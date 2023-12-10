package help

import (
	"RedisDurabilityDB/internal/commands"
	"fmt"
)

// DBHelpCommand is a struct representing the "help" command.
type DBHelpCommand struct {
	AvailableCommands map[string]commands.CommandInterface // A map of available commands
	Name              string
	Description       string
	Required          map[string]struct{}
}

// NewHelpCommand creates a new instance of the DBHelpCommand.
func NewHelpCommand(availableCommands map[string]commands.CommandInterface) *DBHelpCommand {
	return &DBHelpCommand{
		AvailableCommands: availableCommands,
		Name:              "help",
		Description:       "Display a list of available commands\n",
		Required:          map[string]struct{}{},
	}
}

// GetName returns the name of the command.
func (c *DBHelpCommand) GetName() string {
	return c.Name
}

// GetDescription returns a description of the command.
func (c *DBHelpCommand) GetDescription() string {
	return c.Description
}

// GetRequired returns an empty list since "help" has no required arguments.
func (c *DBHelpCommand) GetRequired() map[string]struct{} {
	return c.Required
}

// DoAction executes the "help" command's action.
func (c *DBHelpCommand) DoAction(args map[string]string) error {
	if len(args) > 0 {
		return fmt.Errorf("the 'help' command does not accept any arguments")
	}

	fmt.Println("Available commands:")

	for name, cmd := range c.AvailableCommands {
		fmt.Printf("%s: %s\n", name, cmd.GetDescription())
	}

	return nil
}
