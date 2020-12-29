package house

import "fmt"

// House contains the (fully built) house.
type House struct {
	HouseType  string
	WindowType string
	DoorType   string
	Floor      int
}

// Builder is an interface building a house.
type Builder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

// GetBuilder returns a (house) Builder.
func GetBuilder(builderType string) Builder {
	switch builderType {
	case "normal":
		return &normalBuilder{}
	case "igloo":
		return &iglooBuilder{}
	default:
		return nil
	}
}

// String implements the Stringer interface.
func (h House) String() string {
	return fmt.Sprintf("%s House Door Type: %s\n%s House Window Type: %s\n%s House Num Floor: %d\n",
		h.HouseType, h.DoorType, h.HouseType, h.WindowType, h.HouseType, h.Floor)
}
