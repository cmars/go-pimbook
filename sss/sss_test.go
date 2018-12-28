package sss_test

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cmars/pimbook/poly"
	"github.com/cmars/pimbook/sss"
)

func TestExample(t *testing.T) {
	secret, err := sss.Reveal(poly.NewPointInt(1, 325), poly.NewPointInt(3, 2383), poly.NewPointInt(5, 6609))
	assert.Nil(t, err)
	i := big.NewInt(0).SetBytes(secret)
	assert.Equal(t, i.Int64(), int64(109))
}
