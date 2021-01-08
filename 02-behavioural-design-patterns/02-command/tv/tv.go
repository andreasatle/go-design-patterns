package tv

import "fmt"

type TV struct {
	isRunning bool
}

func (t *TV) On() {
	t.isRunning = true
	fmt.Println("Turning TV ON")
}

func (t *TV) Off() {
	t.isRunning = false
	fmt.Println("Turning TV OFF")
}

func NewTV() *TV {
	return &TV{}
}

func (t *TV) IsRunning() bool {
	return t.isRunning
}

func (t *TV) String() string {
	if t.isRunning {
		return fmt.Sprintln("TV is on!")
	}
	return fmt.Sprintln("TV is off!")
}
