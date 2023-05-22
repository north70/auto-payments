package tg_client

type Command interface {
	Name() string
	Description() string
	Handle(update Update) error
}

type Commands struct {
	Commands map[string]Command
}

func (cmds *Commands) List() map[string]string {
	var list map[string]string

	for _, cmd := range cmds.Commands {
		list[cmd.Name()] = cmd.Description()
	}

	return list
}

func (cmds *Commands) Add(cmd Command) {
	cmds.Commands[cmd.Name()] = cmd
}

func (cmds *Commands) AddMany(cmdList []Command) {
	for _, cmd := range cmdList {
		cmds.Commands[cmd.Name()] = cmd
	}
}

func (cmds *Commands) has(name string) bool {
	_, ok := cmds.Commands[name]

	return ok
}
