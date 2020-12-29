// Package house contains a design pattern to build a new house
// where that house is created after all pieces of the house
// are already built. The House type is a struct which the Builder
// will fill with values. Onces the Builder has filled the House with
// values, it will return a completed House.
//
// A Director contains a Builder which can be changed after creation.
// A Director can build a house and set a Builder.
//
// In this example we have two realizations of a Builder, namely
// normalBuilder and iglooBuilder. The realizations of Builder does not
// have to be exported.
//
// All the routines in the Builder interface are un-exported and hidden
// from the user.
package house
