package tv

type Button struct {
	command Command
}

func (b *Button) Press() {
	b.command.Execute()
}

func NewButton(command Command) *Button {
	return &Button{
		command: command,
	}
}
