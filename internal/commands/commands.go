package commands

type CommandInterface interface {
	GetName() string
	GetDescription() string
	GetRequired() map[string]struct{}
	DoAction(args map[string]string) error
}
