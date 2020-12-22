package mutex

import (
	"fmt"
	"sync"
)

var mutex = &sync.Mutex{}

type Single struct {
}

var singleInstance *Single

func GetInstance() *Single {
	if singleInstance != nil {
		fmt.Println("Single Instance already created.")
		return singleInstance
	}

	mutex.Lock()
	defer mutex.Unlock()

	if singleInstance != nil {
		fmt.Println("Single Instance already created during lock.")
		return singleInstance
	}

	fmt.Println("Creating Single Instance Now.")
	singleInstance = &Single{}
	return singleInstance
}
