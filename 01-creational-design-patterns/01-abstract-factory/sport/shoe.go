package sport

// ShoeInterface is a part of the SportsFactory interface.
type ShoeInterface interface {
	SetLogo(logo string)
	SetSize(size int)
	GetLogo() string
	GetSize() int
}

// Non-visible type that implements the ShoeInterface.
type shoe struct {
	logo string
	size int
}

// SetLogo is a part of ShoeInterface.
func (s *shoe) SetLogo(logo string) {
	s.logo = logo
}

// GetLogo is a part of ShoeInterface.
func (s *shoe) GetLogo() string {
	return s.logo
}

// SetSize is a part of ShoeInterface.
func (s *shoe) SetSize(size int) {
	s.size = size
}

// GetSize is a part of ShoeInterface.
func (s *shoe) GetSize() int {
	return s.size
}
