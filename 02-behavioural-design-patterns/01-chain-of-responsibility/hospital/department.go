package hospital

// Department is an interface, looks like a linked list to me.
type Department interface {
	Execute(*Patient)
	SetNext(Department)
}
