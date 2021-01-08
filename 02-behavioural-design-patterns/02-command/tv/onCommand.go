package tv

type OnCommand struct {
	device Device
}

func (c *OnCommand) Execute() {
	c.device.On()
}

func NewOnCommand(tv *TV) *OnCommand {
	return &OnCommand{
		device: tv,
	}
}
