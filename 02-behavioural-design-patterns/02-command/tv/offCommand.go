package tv

type OffCommand struct {
	device Device
}

func (c *OffCommand) Execute() {
	c.device.Off()
}

func NewOffCommand(tv *TV) *OffCommand {
	return &OffCommand{
		device: tv,
	}
}
