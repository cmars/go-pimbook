package poly_test

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cmars/pimbook/poly"
)

func TestAdd2PolyDegree2(t *testing.T) {
	p1 := poly.NewInt(1, 2, 1)
	p2 := poly.NewInt(3, 5, 7)
	sum := poly.New().Add(p1, p2)
	assert.Equal(t, sum, poly.NewInt(4, 7, 8))
}

func TestAdd2PolyUnequalDegree(t *testing.T) {
	p1 := poly.NewInt(1, 2)
	p2 := poly.NewInt(3, 5, 7)
	sum := poly.New().Add(p1, p2)
	assert.Equal(t, sum, poly.NewInt(4, 7, 7))
}

func TestAdd2PolyUnequalDegreeLargerFirst(t *testing.T) {
	p1 := poly.NewInt(1, 2, 1)
	p2 := poly.NewInt(3, 5)
	sum := poly.New().Add(p1, p2)
	assert.Equal(t, sum, poly.NewInt(4, 7, 1))
}

func TestAddZeroPolyFirst(t *testing.T) {
	p1 := poly.New()
	p2 := poly.NewInt(3, 5, 7)
	sum := poly.New().Add(p1, p2)
	assert.Equal(t, sum, poly.NewInt(3, 5, 7))
}

func TestAddZeroPolyLast(t *testing.T) {
	p1 := poly.NewInt(3, 5, 7)
	p2 := poly.New()
	sum := poly.New().Add(p1, p2)
	assert.Equal(t, sum, poly.NewInt(3, 5, 7))
}

func TestAddToZeroSelf(t *testing.T) {
	p1 := poly.New()
	sum := p1.Add(p1, poly.NewInt(3, 5, 7))
	assert.Equal(t, sum, poly.NewInt(3, 5, 7))
}

func TestAddToNonzeroSelf(t *testing.T) {
	p1 := poly.NewInt(1, 2, 1)
	sum := p1.Add(p1, poly.NewInt(3, 5, 7))
	assert.Equal(t, sum, poly.NewInt(4, 7, 8))
}

func TestString(t *testing.T) {
	assert.Equal(t, poly.NewInt(7).String(), "7")
	assert.Equal(t, poly.NewInt(7, 2).String(), "7+2x")
	assert.Equal(t, poly.NewInt(7, 2, 9).String(), "7+2x+9x^2")
	assert.Equal(t, poly.NewInt(-7, -2, -9).String(), "-7-2x-9x^2")
	assert.Equal(t, poly.NewInt(7, 0, 9).String(), "7+9x^2")
	assert.Equal(t, poly.NewInt(0, 0, 9).String(), "9x^2")
}

func TestMul(t *testing.T) {
	p1 := poly.NewInt(1, 2, 1)
	p2 := poly.NewInt(7, 5, 3)
	prod := poly.New().Mul(p1, p2)
	assert.Equal(t, prod, poly.NewInt(7, 19, 20, 11, 3))
}

func TestEval(t *testing.T) {
	p := poly.NewInt(1, 2, 1)
	f9 := p.Eval(big.NewRat(9, 1))
	assert.Equal(t, f9.String(), "100/1")

	p = poly.New()
	f9 = p.Eval(big.NewRat(9, 1))
	assert.Equal(t, f9.String(), "0/1")

	p = poly.NewInt(0, 0, 0, 1)
	f9 = p.Eval(big.NewRat(9, 1))
	assert.Equal(t, f9.String(), "729/1")
}

func TestInterpolate(t *testing.T) {
	p, err := poly.Interpolate(poly.NewPointInt(1, 1))
	assert.Nil(t, err)
	assert.Equal(t, p, poly.NewInt(1))

	p, err = poly.Interpolate(poly.NewPointInt(1, 1), poly.NewPointInt(2, 0))
	assert.Nil(t, err)
	assert.Equal(t, p, poly.NewInt(2, -1))

	p, err = poly.Interpolate(poly.NewPointInt(1, 1), poly.NewPointInt(2, 4), poly.NewPointInt(7, 9))
	assert.Nil(t, err)
	assert.Equal(t, p, poly.New(
		big.NewRat(-8, 3), big.NewRat(4, 1), big.NewRat(-1, 3)))
}
