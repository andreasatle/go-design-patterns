// Package sport implements an abstract factory design pattern.
// The package consists of three interfaces:
//
//  ShoeInterface implements
//    SetLogo(logo string)
//    SetSize(size int)
//    GetLogo() string
//    GetSize() int
//
//  ShortsInterface implements
//    SetLogo(logo string)
//    SetSize(size int)
//    GetLogo() string
//    GetSize() int
//
//  SportsFactory implements
//    MakeShoe(int) ShoeInterface
//    MakeShorts(int) ShortsInterface
package sport
