package main

import (
	"log"

	"github.com/andreasatle/go-design-patterns/01-creational-design-patterns/04-object-pool/pool"
)

func main() {

	var err error

	pool, err := pool.InitPool(createConnections(2))
	if err != nil {
		log.Fatalf("Init Pool Error: %v", err)
	}

	conn1, err := pool.Get()
	if err != nil {
		log.Printf("Get Pool Error: %v", err)
	}

	conn2, err := pool.Get()
	if err != nil {
		log.Printf("Get Pool Error: %v", err)
	}

	conn3, err := pool.Get()
	if err != nil {
		log.Printf("Get Pool Error: %v", err)
	}

	err = pool.Return(conn1)
	if err != nil {
		log.Printf("Return Pool Error: %v", err)
	}

	err = pool.Return(conn2)
	if err != nil {
		log.Printf("Return Pool Error: %v", err)
	}

	err = pool.Return(conn2)
	if err != nil {
		log.Printf("Return Pool Error: %v", err)
	}

	err = pool.Return(conn3)
	if err != nil {
		log.Printf("Return Pool Error: %v", err)
	}
}
