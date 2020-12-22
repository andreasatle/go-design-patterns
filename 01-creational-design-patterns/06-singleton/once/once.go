package once

import (
	"fmt"
	"sync"
)

var once sync.Once

type Single struct {
}

var singleInstance *Single

func GetInstance() *Single {
	if singleInstance != nil {
		fmt.Println("Single Instance already created.")
		return singleInstance
	}

	once.Do(
		func() {
			fmt.Println("Creating Single Instance Now.")
			singleInstance = &Single{}
		})
	return singleInstance
}
