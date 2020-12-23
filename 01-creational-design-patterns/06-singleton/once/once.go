// Package mutex contains a thread-safe implementation of
// a singleton instance using sync.Mutex.
//
// We use the package variable once (sync.Once) to guarantee
// that the singleton instance is created exactly once.
package once

import (
	"fmt"
	"sync"
)

var once sync.Once

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

	once.Do(
		func() {
			fmt.Println("Creating Single Instance Now.")
			singleInstance = &Single{}
		})
	return singleInstance
}
