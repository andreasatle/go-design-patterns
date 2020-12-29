package sport

// Non-visible type that implements the SportsFactory interface.
type adidas struct {
}

// Non-visible type that implements the ShoeInterface.
type adidasShoe struct {
	shoe
}

// Non-visible type that implements the ShortsInterface.
type adidasShorts struct {
	shorts
}

// MakeShoe implements part of SportsFactory interface.
func (a *adidas) MakeShoe(size int) ShoeInterface {
	return &adidasShoe{
		shoe: shoe{
			logo: "adidas",
			size: size,
		},
	}
}

// MakeShorts implements part of SportsFactory interface.
func (a *adidas) MakeShorts(size int) ShortsInterface {
	return &adidasShorts{
		shorts: shorts{
			logo: "adidas",
			size: size,
		},
	}
}
