package int_byte_utils

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func testIntReduced(t *testing.T, v int) {
	t.Log(v)
	buff := bytes.NewBuffer([]byte{})
	n, err := WriteIntAsBytesReduced(v, buff)
	assert.NoError(t, err)
	t.Log(n)
	t.Log(buff.Bytes())
	t.Log(LenIntAsBytesReduced(v))
	assert.Equal(t, LenIntAsBytesReduced(v), buff.Len())
	var val int
	n, err, val = ReadIntFromBytesReduced(buff)
	assert.NoError(t, err)
	t.Log(n)
	t.Log(val)
	assert.Equal(t, v, val)
}

func TestIntReduced(t *testing.T) {
	t.Run("Int0", func(t *testing.T) {
		testIntReduced(t, 0)
	})
	t.Run("Int1", func(t *testing.T) {
		testIntReduced(t, 1)
	})
	t.Run("Int-1", func(t *testing.T) {
		testIntReduced(t, -1)
	})
	t.Run("Int65537", func(t *testing.T) {
		testIntReduced(t, 65537)
	})
	t.Run("Int-65537", func(t *testing.T) {
		testIntReduced(t, -65537)
	})
	t.Run("IntMax", func(t *testing.T) {
		testIntReduced(t, math.MaxInt)
	})
	t.Run("IntMax-1", func(t *testing.T) {
		testIntReduced(t, math.MaxInt-1)
	})
	t.Run("IntMin+1", func(t *testing.T) {
		testIntReduced(t, math.MinInt+1)
	})
	t.Run("IntMin", func(t *testing.T) {
		testIntReduced(t, math.MinInt)
	})
}
