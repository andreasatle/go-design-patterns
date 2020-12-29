package sport

// ShortsInterface is a part of the SportsFactory interface.
type ShortsInterface interface {
	SetLogo(logo string)
	SetSize(size int)
	GetLogo() string
	GetSize() int
}

// Non-visible type that implements the ShortsInterface.
type shorts struct {
	logo string
	size int
}

// SetLogo is a part of ShortsInterface.
func (s *shorts) SetLogo(logo string) {
	s.logo = logo
}

// GetLogo is a part of ShortsInterface.
func (s *shorts) GetLogo() string {
	return s.logo
}

// SetSize is a part of ShortsInterface.
func (s *shorts) SetSize(size int) {
	s.size = size
}

// GetSize is a part of ShortsInterface.
func (s *shorts) GetSize() int {
	return s.size
}
