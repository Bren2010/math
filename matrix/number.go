package matrix

import (
	"math/big"
)

// Number is a field element.
type Number interface {
	Neg(x Number) Number
	Inv(x Number) Number

	Add(x, y Number) Number
	Mul(x, y Number) Number

	Cmp(x Number) int
	Sign() int

	Set(x Number) Number
	SetZero() Number
	SetOne() Number

	String() string
}

// NewNumber is called to allocate a new number. It can be replaced with any other implementation of this function.
var NewNumber func() Number = NewRat

// Rat is a wrapper around big.Rat which implements Number.
type Rat big.Rat

// NewRat returns a new, zero rational number.
func NewRat() Number {
	return (*Rat)(new(big.Rat))
}

// NewRatI returns the new rational number a/1.
func NewRatI(a int64) Number {
	return (*Rat)(big.NewRat(a, 1))
}

func (r *Rat) Neg(x Number) Number { return (*Rat)((*big.Rat)(r).Neg((*big.Rat)(x.(*Rat)))) }
func (r *Rat) Inv(x Number) Number { return (*Rat)((*big.Rat)(r).Inv((*big.Rat)(x.(*Rat)))) }

func (r *Rat) Add(x, y Number) Number {
	return (*Rat)((*big.Rat)(r).Add((*big.Rat)(x.(*Rat)), (*big.Rat)(y.(*Rat))))
}
func (r *Rat) Mul(x, y Number) Number {
	return (*Rat)((*big.Rat)(r).Mul((*big.Rat)(x.(*Rat)), (*big.Rat)(y.(*Rat))))
}

func (r *Rat) Cmp(x Number) int { return (*big.Rat)(r).Cmp((*big.Rat)(x.(*Rat))) }
func (r *Rat) Sign() int        { return (*big.Rat)(r).Sign() }

func (r *Rat) Set(x Number) Number { return (*Rat)((*big.Rat)(r).Set((*big.Rat)(x.(*Rat)))) }
func (r *Rat) SetZero() Number     { return (*Rat)((*big.Rat)(r).Set(big.NewRat(0, 1))) }
func (r *Rat) SetOne() Number      { return (*Rat)((*big.Rat)(r).Set(big.NewRat(1, 1))) }

func (r *Rat) String() string { return (*big.Rat)(r).String() }
