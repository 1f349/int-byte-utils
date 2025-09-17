package int_byte_utils

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

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

func testUintReadOverread(t *testing.T) {
	const btsLen = 255
	bts := make([]byte, btsLen)
	for i := 0; i < btsLen; i++ {
		bts[i] = byte(bit8)
	}
	buff := bytes.NewBuffer(bts)
	var n, err, val = ReadUintFromBytes(buff)
	t.Log(n)
	t.Log(err)
	t.Log(val)
	assert.NoError(t, err)
	assert.Equal(t, uint(math.MaxInt)+1, val)
	assert.Equal(t, 9, n)
}

func testUintReadOverflow(t *testing.T) {
	buff := bytes.NewBuffer([]byte{128, 128, 128, 128, 128,
		128, 128, 128, 128, 3})
	var n, err, val = ReadUintFromBytes(buff)
	t.Log(n)
	t.Log(err)
	t.Log(val)
	assert.NoError(t, err)
	assert.Equal(t, uint(math.MaxInt)+1, val)
	assert.Equal(t, 9, n)
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
	t.Run("UintOverread", testUintReadOverread)
	t.Run("UintOverflow", testUintReadOverflow)
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
