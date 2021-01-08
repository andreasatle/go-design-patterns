package main

import (
	"fmt"

	"github.com/andreasatle/go-design-patterns/02-behavioural-design-patterns/02-command/tv"
)

func main() {

	teve := tv.NewTV()

	onCommand := tv.NewOnCommand(teve)

	offCommand := tv.NewOffCommand(teve)

	onButton := tv.NewButton(onCommand)

	offButton := tv.NewButton(offCommand)

	fmt.Println(teve)
	onButton.Press()
	fmt.Println(teve)
	onButton.Press()
	fmt.Println(teve)
	offButton.Press()
	fmt.Println(teve)
	offButton.Press()
	fmt.Println(teve)

	teve.On()
	fmt.Println(teve)
	teve.Off()
	fmt.Println(teve)
}
