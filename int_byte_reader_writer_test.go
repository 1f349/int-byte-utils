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
	var val int
	n, err, val = ReadIntFromBytes(buff)
	assert.NoError(t, err)
	t.Log(n)
	t.Log(val)
	assert.Equal(t, v, val)
}

func testUint(t *testing.T, v uint) {
	t.Log(v)
	buff := bytes.NewBuffer([]byte{})
	n, err := WriteUintAsBytes(v, buff)
	assert.NoError(t, err)
	t.Log(n)
	t.Log(buff.Bytes())
	var val uint
	n, err, val = ReadUintFromBytes(buff)
	assert.NoError(t, err)
	t.Log(n)
	t.Log(val)
	assert.Equal(t, v, val)
}

func testUintReadInt(t *testing.T, v uint) {
	t.Log(v)
	buff := bytes.NewBuffer([]byte{})
	n, err := WriteUintAsBytes(v, buff)
	assert.NoError(t, err)
	t.Log(n)
	t.Log(buff.Bytes())
	var val int
	n, err, val = ReadIntFromBytes(buff)
	assert.NoError(t, err)
	t.Log(n)
	t.Log(val)
	if v == highestBit {
		assert.Equal(t, math.MinInt, val)
	} else if (v & highestBit) == 0 {
		assert.Equal(t, int(v), val)
	} else {
		assert.Equal(t, int(-(v & ^highestBit)), val)
	}
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

func TestUint(t *testing.T) {
	t.Run("Uint0", func(t *testing.T) {
		testUint(t, 0)
	})
	t.Run("Uint1", func(t *testing.T) {
		testUint(t, 1)
	})
	t.Run("Uint65537", func(t *testing.T) {
		testUint(t, 65537)
	})
	t.Run("IntMax", func(t *testing.T) {
		testUint(t, math.MaxInt)
	})
	t.Run("IntMax+1", func(t *testing.T) {
		testUint(t, math.MaxInt+1)
	})
	t.Run("UintMax", func(t *testing.T) {
		testUint(t, math.MaxUint)
	})
}

func TestUintReadInt(t *testing.T) {
	t.Run("Uint0", func(t *testing.T) {
		testUintReadInt(t, 0)
	})
	t.Run("Uint1", func(t *testing.T) {
		testUintReadInt(t, 1)
	})
	t.Run("Uint65537", func(t *testing.T) {
		testUintReadInt(t, 65537)
	})
	t.Run("IntMax", func(t *testing.T) {
		testUintReadInt(t, math.MaxInt)
	})
	t.Run("IntMax+1", func(t *testing.T) {
		testUintReadInt(t, math.MaxInt+1)
	})
	t.Run("UintMax", func(t *testing.T) {
		testUintReadInt(t, math.MaxUint)
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
