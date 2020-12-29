package sport

import "fmt"

// SportsFactory is an interface handling sportswear.
type SportsFactory interface {
	MakeShoe(int) ShoeInterface
	MakeShorts(int) ShortsInterface
}

// GetSportsFactory creates a new instance of a SportsFactory.
func GetSportsFactory(brand string) (SportsFactory, error) {
	switch brand {
	case "adidas":
		return &adidas{}, nil
	case "nike":
		return &nike{}, nil
	default:
		return nil, fmt.Errorf("Wrong brand type passed")
	}
}
