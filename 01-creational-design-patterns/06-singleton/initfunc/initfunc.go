// Package initfunc contains a thread-safe implementation of
// a singleton instance using init().
package initfunc

import "log"

// Single is an empty struct, but could be filled with content.
type Single struct {
}

var singleInstance *Single

func init() {
	log.Println("Create singleInstance.")
	singleInstance = &Single{}
}

// GetInstance returns a pointer to the singleton instance.
// The instance is assumed to already be initialized by init().
func GetInstance() *Single {
	if singleInstance == nil {
		log.Panic("singleInstance is not initialized.")
	}

	log.Println("singleInstance already created.")
	return singleInstance
}
