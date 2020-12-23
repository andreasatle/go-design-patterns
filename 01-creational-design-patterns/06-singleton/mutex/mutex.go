// Package mutex contains a thread-safe implementation of
// a singleton instance using sync.Mutex.
//
// Extra care has to be taken to check if the instance is
// initialized both before and after the mutex-lock. This is to
// ensure that the instance is not created by another go-routine
// inbetween the first check and the mutex-lock.
package mutex

import (
	"fmt"
	"sync"
)

var mutex = &sync.Mutex{}

// Single is an empty struct, but could be filled with content.
type Single struct {
}

var singleInstance *Single

// GetInstance returns a pointer to the singleton instance.
// The instance is created if it is not already initialized.
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
