package poly

import (
	"errors"
	"fmt"
	"math/big"
	"sort"
)

// Point represents the evaluation of a polynomial at a given value, where X is
// that value, and Y is the result of evaluating the polynomial at that value.
type Point struct {
	X, Y *big.Int
}

// NewPoint returns a new Point for given (x, y).
func NewPoint(x, y *big.Int) *Point {
	return &Point{X: x, Y: y}
}

// NewPoint returns a new Point for given (x, y) as int64 values.
func NewPointInt(x, y int64) *Point {
	return NewPoint(big.NewInt(x), big.NewInt(y))
}

// Cmp returns 1 if p > q, -1 if p < q, 0 if equal.
func (p *Point) Cmp(q *Point) int {
	if cmp := p.X.Cmp(q.X); cmp != 0 {
		return cmp
	}
	return p.Y.Cmp(q.Y)
}

type pointSlice []*Point

func (ps pointSlice) Len() int           { return len(ps) }
func (ps pointSlice) Less(i, j int) bool { return ps[i].Cmp(ps[j]) < 0 }
func (ps pointSlice) Swap(i, j int)      { ps[i], ps[j] = ps[j], ps[i] }

// Interpolate returns a polynomial that fits all of the given points,
// so long as the x values of each point is unique.
func Interpolate(points ...*Point) (*Poly, error) {
	if len(points) == 0 {
		return nil, errors.New("no points given")
	}
	sort.Sort(pointSlice(points))
	for i := range points {
		if i > 0 && points[i].Cmp(points[i-1]) == 0 {
			return nil, fmt.Errorf("duplicate x %q", points[i].X.String())
		}
	}
	result := New()
	for i := range points {
		term := singleTerm(points, i)
		result.Add(result, term)
	}
	return result, nil
}

func singleTerm(points []*Point, i int) *Poly {
	result := NewInt(1)
	xi, yi := points[i].X, points[i].Y
	for j := range points {
		if i != j {
			xj := points[j].X
			xijDiff := big.NewInt(0).Sub(xi, xj)
			a0 := big.NewRat(0, 1).SetFrac(big.NewInt(0).Neg(xj), xijDiff)
			a1 := big.NewRat(0, 1).SetFrac(big.NewInt(1), xijDiff)
			result.Mul(result, New(a0, a1))
		}
	}
	return result.Mul(result, New(big.NewRat(0, 1).SetFrac(yi, big.NewInt(1))))
}
