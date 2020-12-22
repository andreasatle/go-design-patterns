package initfunc

import "log"

type Single struct {
}

var singleInstance *Single

func init() {
	log.Println("Create singleInstance.")
	singleInstance = &Single{}
}

func GetInstance() *Single {
	if singleInstance == nil {
		log.Panic("singleInstance is not initialized.")
	}

	log.Println("singleInstance already created.")
	return singleInstance
}
