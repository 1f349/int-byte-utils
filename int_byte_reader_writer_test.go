package int_byte_utils

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func testInt(t *testing.T, v int) {
	t.Log(v)
	buff := bytes.NewBuffer([]byte{})
	n, err := WriteIntAsBytes(v, buff)
	assert.NoError(t, err)
	t.Log(n)
	t.Log(buff.Bytes())
	t.Log(LenIntAsBytes(v))
	assert.Equal(t, LenIntAsBytes(v), buff.Len())
	var val int
	n, err, val = ReadIntFromBytes(buff)
	assert.NoError(t, err)
	t.Log(n)
	t.Log(val)
	assert.Equal(t, v, val)
}

func testIntReadUint(t *testing.T, v int) {
	t.Log(v)
	buff := bytes.NewBuffer([]byte{})
	n, err := WriteIntAsBytes(v, buff)
	assert.NoError(t, err)
	t.Log(n)
	t.Log(buff.Bytes())
	var val uint
	n, err, val = ReadUintFromBytes(buff)
	assert.NoError(t, err)
	t.Log(n)
	t.Log(val)
	if v == math.MinInt {
		assert.Equal(t, highestBit, val)
	} else if v < 0 {
		assert.Equal(t, val, uint(-v)|highestBit)
	} else {
		assert.Equal(t, val, uint(v))
	}
}

func TestInt(t *testing.T) {
	t.Run("Int0", func(t *testing.T) {
		testInt(t, 0)
	})
	t.Run("Int1", func(t *testing.T) {
		testInt(t, 1)
	})
	t.Run("Int-1", func(t *testing.T) {
		testInt(t, -1)
	})
	t.Run("Int65537", func(t *testing.T) {
		testInt(t, 65537)
	})
	t.Run("Int-65537", func(t *testing.T) {
		testInt(t, -65537)
	})
	t.Run("IntMax", func(t *testing.T) {
		testInt(t, math.MaxInt)
	})
	t.Run("IntMax-1", func(t *testing.T) {
		testInt(t, math.MaxInt-1)
	})
	t.Run("IntMin+1", func(t *testing.T) {
		testInt(t, math.MinInt+1)
	})
	t.Run("IntMin", func(t *testing.T) {
		testInt(t, math.MinInt)
	})
}

func TestIntReadUint(t *testing.T) {
	t.Run("Int0", func(t *testing.T) {
		testIntReadUint(t, 0)
	})
	t.Run("Int1", func(t *testing.T) {
		testIntReadUint(t, 1)
	})
	t.Run("Int-1", func(t *testing.T) {
		testIntReadUint(t, -1)
	})
	t.Run("Int65537", func(t *testing.T) {
		testIntReadUint(t, 65537)
	})
	t.Run("Int-65537", func(t *testing.T) {
		testIntReadUint(t, -65537)
	})
	t.Run("IntMax", func(t *testing.T) {
		testIntReadUint(t, math.MaxInt)
	})
	t.Run("IntMax-1", func(t *testing.T) {
		testIntReadUint(t, math.MaxInt-1)
	})
	t.Run("IntMin+1", func(t *testing.T) {
		testIntReadUint(t, math.MinInt+1)
	})
	t.Run("IntMin", func(t *testing.T) {
		testIntReadUint(t, math.MinInt)
	})
}
