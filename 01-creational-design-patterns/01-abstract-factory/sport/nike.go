package sport

// Non-visible type that implements the SportsFactory interface.
type nike struct {
}

// Non-visible type that implements the ShoeInterface.
type nikeShoe struct {
	shoe
}

// Non-visible type that implements the ShortsInterface.
type nikeShorts struct {
	shorts
}

// MakeShoe implements part of SportsFactory interface.
func (a *nike) MakeShoe(size int) ShoeInterface {
	return &nikeShoe{
		shoe: shoe{
			logo: "nike",
			size: size,
		},
	}
}

// MakeShorts implements part of SportsFactory interface.
func (a *nike) MakeShorts(size int) ShortsInterface {
	return &nikeShorts{
		shorts: shorts{
			logo: "nike",
			size: size,
		},
	}
}
