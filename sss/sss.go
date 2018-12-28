package sss

import (
	"crypto/rand"
	"errors"
	"math/big"

	"github.com/cmars/pimbook/poly"
)

func Split(secret []byte, n, k int) ([]*poly.Point, error) {
	coeffs := []*big.Rat{newRatBytes(secret)}
	for i := 0; i < k-1; i++ {
		coeffBytes := make([]byte, len(secret))
		_, err := rand.Reader.Read(coeffBytes)
		if err != nil {
			return nil, err
		}
		coeffs = append(coeffs, newRatBytes(coeffBytes))
	}
	p := poly.New(coeffs...)
	points := make([]*poly.Point, n)
	for i := 0; i < n; i++ {
		x := big.NewInt(int64(i + 1))
		points[i-1] = poly.NewPoint(x, p.Eval(big.NewRat(0, 1).SetFrac(x, big.NewInt(1))).Num())
	}
	return points, nil
}

func Reveal(points []*poly.Point) ([]byte, error) {
	p, err := poly.Interpolate(points...)
	if err != nil {
		return nil, err
	}
	secret := p.Coeff()[0]
	if !secret.IsInt() {
		return nil, errors.New("invalid")
	}
	return p.Coeff()[0].Num().Bytes(), nil
}

func newRatBytes(b []byte) *big.Rat {
	return big.NewRat(0, 1).SetFrac(big.NewInt(0).SetBytes(b), big.NewInt(1))
}
