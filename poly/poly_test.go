package poly_test

import (
	"math/big"
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/google/go-cmp/cmp"

	"github.com/cmars/pimbook/poly"
)

var ratSliceEquals = qt.CmpEquals(cmp.Transformer("ratstring", func(in *big.Rat) string { return in.String() }))

func TestAdd(t *testing.T) {
	t.Run("add 2 polys, both degree=2", func(t *testing.T) {
		c := qt.New(t)
		p1 := poly.NewInt64(1, 2, 1)
		p2 := poly.NewInt64(3, 5, 7)
		sum := poly.New().Add(p1, p2)
		c.Assert(sum.Coeff(), ratSliceEquals, poly.NewInt64(4, 7, 8).Coeff())
	})
	t.Run("add 2 polys, unequal degree, smaller first", func(t *testing.T) {
		c := qt.New(t)
		p1 := poly.NewInt64(1, 2)
		p2 := poly.NewInt64(3, 5, 7)
		sum := poly.New().Add(p1, p2)
		c.Assert(sum.Coeff(), ratSliceEquals, []*big.Rat{
			big.NewRat(4, 1), big.NewRat(7, 1), big.NewRat(7, 1)})
	})
	t.Run("add 2 polys, unequal degree, larger first", func(t *testing.T) {
		c := qt.New(t)
		p1 := poly.NewInt64(1, 2, 1)
		p2 := poly.NewInt64(3, 5)
		sum := poly.New().Add(p1, p2)
		c.Assert(sum.Coeff(), ratSliceEquals, []*big.Rat{
			big.NewRat(4, 1), big.NewRat(7, 1), big.NewRat(1, 1)})
	})
	t.Run("add 0 poly first", func(t *testing.T) {
		c := qt.New(t)
		p1 := poly.New()
		p2 := poly.NewInt64(3, 5, 7)
		sum := poly.New().Add(p1, p2)
		c.Assert(sum.Coeff(), ratSliceEquals, []*big.Rat{
			big.NewRat(3, 1), big.NewRat(5, 1), big.NewRat(7, 1)})
	})
	t.Run("add 0 poly last", func(t *testing.T) {
		c := qt.New(t)
		p1 := poly.NewInt64(3, 5, 7)
		p2 := poly.New()
		sum := poly.New().Add(p1, p2)
		c.Assert(sum.Coeff(), ratSliceEquals, []*big.Rat{
			big.NewRat(3, 1), big.NewRat(5, 1), big.NewRat(7, 1)})
	})
	t.Run("add to zero, self", func(t *testing.T) {
		c := qt.New(t)
		p1 := poly.New()
		sum := p1.Add(p1, poly.NewInt64(3, 5, 7))
		c.Assert(sum.Coeff(), ratSliceEquals, []*big.Rat{
			big.NewRat(3, 1), big.NewRat(5, 1), big.NewRat(7, 1)})
	})
	t.Run("add to nonzero, self", func(t *testing.T) {
		c := qt.New(t)
		p1 := poly.NewInt64(1, 2, 1)
		sum := p1.Add(p1, poly.NewInt64(3, 5, 7))
		c.Assert(sum.Coeff(), ratSliceEquals, poly.NewInt64(4, 7, 8).Coeff())
	})
}

func TestString(t *testing.T) {
	t.Run("strings", func(t *testing.T) {
		c := qt.New(t)
		c.Assert(poly.NewInt64(7).String(), qt.Equals, "7")
		c.Assert(poly.NewInt64(7, 2).String(), qt.Equals, "7+2x")
		c.Assert(poly.NewInt64(7, 2, 9).String(), qt.Equals, "7+2x+9x^2")
		c.Assert(poly.NewInt64(-7, -2, -9).String(), qt.Equals, "-7-2x-9x^2")
		c.Assert(poly.NewInt64(7, 0, 9).String(), qt.Equals, "7+9x^2")
		c.Assert(poly.NewInt64(0, 0, 9).String(), qt.Equals, "9x^2")
	})
}

func TestMul(t *testing.T) {
	t.Run("multiply 2 polys, both degree=2", func(t *testing.T) {
		c := qt.New(t)
		p1 := poly.NewInt64(1, 2, 1)
		p2 := poly.NewInt64(7, 5, 3)
		prod := poly.New().Mul(p1, p2)
		c.Assert(prod.Coeff(), ratSliceEquals, poly.NewInt64(7, 19, 20, 11, 3).Coeff())
	})
}

func TestEval(t *testing.T) {
	t.Run("evaluate poly at x", func(t *testing.T) {
		c := qt.New(t)
		p := poly.NewInt64(1, 2, 1)
		f9 := p.Eval(big.NewRat(9, 1))
		c.Assert(f9.String(), qt.Equals, "100/1")

		p = poly.New()
		f9 = p.Eval(big.NewRat(9, 1))
		c.Assert(f9.String(), qt.Equals, "0/1")

		p = poly.NewInt64(0, 0, 0, 1)
		f9 = p.Eval(big.NewRat(9, 1))
		c.Assert(f9.String(), qt.Equals, "729/1")
	})
}
