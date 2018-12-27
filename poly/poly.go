// Package poly provides polynomial arithmetic.
package poly

import (
	"bytes"
	"fmt"
	"math/big"
)

// Poly represents a polynomial.
type Poly struct {
	coeff []*big.Rat
}

// New creates a new polynomial from rational coefficients.
func New(coeff ...*big.Rat) *Poly {
	return &Poly{coeff: coeff}
}

// New creates a new polynomial from integer coefficients.
func NewInt64(coeff ...int64) *Poly {
	ratCoeff := make([]*big.Rat, len(coeff))
	for i := range coeff {
		ratCoeff[i] = big.NewRat(coeff[i], 1)
	}
	return New(ratCoeff...)
}

// Degree returns the degree of the polynomial.
func (p *Poly) Degree() int {
	for i := len(p.coeff) - 1; i > 0; i-- {
		if !isZeroRat(p.coeff[i]) {
			return i
		}
	}
	return 0
}

// Add adds two polynomials (a, b), storing the result in c.
func (c *Poly) Add(a, b *Poly) *Poly {
	maxLen := len(a.coeff)
	if bLen := len(b.coeff); bLen > maxLen {
		maxLen = bLen
	}
	var coeff []*big.Rat
	for i := 0; i < maxLen; i++ {
		var ac, bc *big.Rat
		if i < len(a.coeff) {
			ac = a.coeff[i]
		} else {
			ac = big.NewRat(0, 1)
		}
		if i < len(b.coeff) {
			bc = b.coeff[i]
		} else {
			bc = big.NewRat(0, 1)
		}
		coeff = append(coeff, big.NewRat(0, 1).Add(ac, bc))
	}
	c.coeff = coeff
	return c
}

// Mul multiplies two polynomials (a, b), storing the result in c.
func (c *Poly) Mul(a, b *Poly) *Poly {
	var sum *Poly
	for ai := range a.coeff {
		var coeff []*big.Rat
		for i := 0; i < ai; i++ {
			coeff = append(coeff, big.NewRat(0, 1))
		}
		for bi := range b.coeff {
			coeff = append(coeff, big.NewRat(0, 1).Mul(a.coeff[ai], b.coeff[bi]))
		}
		if sum == nil {
			sum = New(coeff...)
		} else {
			sum.Add(sum, New(coeff...))
		}
	}
	c.coeff = sum.coeff
	return c
}

// String returns a string representation of the polynomial.
func (c *Poly) String() string {
	var buf bytes.Buffer
	terms := 0
	for i := range c.coeff {
		if cmp := c.coeff[i].Cmp(zeroRat); cmp == 0 {
			continue
		} else if cmp < 0 {
			fmt.Fprintf(&buf, "-")
		} else if terms > 0 {
			fmt.Fprintf(&buf, "+")
		}
		a := big.NewRat(0, 1).Abs(c.coeff[i])
		if c.coeff[i].IsInt() {
			fmt.Fprintf(&buf, "%s", a.Num().String())
		} else {
			fmt.Fprintf(&buf, "%s", a)
		}
		if i > 0 {
			fmt.Fprintf(&buf, "x")
			if i > 1 {
				fmt.Fprintf(&buf, "^%d", i)
			}
		}
		terms++
	}
	return buf.String()
}

// Eval evaluates the polynomial at a given value for x.
func (c *Poly) Eval(x *big.Rat) *big.Rat {
	sum := big.NewRat(0, 1)
	for i := range c.coeff {
		var xn *big.Rat
		if i == 0 {
			xn = big.NewRat(1, 1)
		} else if i == 1 {
			xn = x
		} else {
			xn = ratExp(x, int64(i))
		}
		sum.Add(sum, big.NewRat(0, 1).Mul(c.coeff[i], xn))
	}
	return sum
}

func ratExp(x *big.Rat, y int64) *big.Rat {
	bigY := big.NewInt(y)
	if x.IsInt() {
		a := big.NewInt(0).Exp(x.Num(), bigY, nil)
		return big.NewRat(0, 1).SetFrac(a, big.NewInt(1))
	}
	a := big.NewInt(0).Exp(x.Num(), bigY, nil)
	b := big.NewInt(0).Exp(x.Denom(), bigY, nil)
	return big.NewRat(0, 1).SetFrac(a, b)
}

// Coeff returns the coefficients of the polynomial.
func (c *Poly) Coeff() []*big.Rat {
	return c.coeff
}

var (
	zeroInt = big.NewInt(0)
	zeroRat = big.NewRat(0, 1)
)

func isZeroRat(n *big.Rat) bool {
	return isZeroInt(n.Num())
}

func isZeroInt(n *big.Int) bool {
	return n.Cmp(zeroInt) == 0
}
