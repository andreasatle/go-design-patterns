package main

import (
	"strconv"

	"github.com/andreasatle/go-design-patterns/01-creational-design-patterns/04-object-pool/pool"
)

// connection implements pool.Object
type connection struct {
	id string
}

// GetId implements interface pool.Object.
func (c *connection) GetID() string {
	return c.id
}

// createConnections creates a list of pool objects.
func createConnections(num int) []pool.Object {
	// Observe that we use []pool.Object containing *connections
	connections := make([]pool.Object, 0)
	for i := 0; i < num; i++ {
		c := &connection{id: strconv.Itoa(i)}
		connections = append(connections, c)
	}
	return connections
}
