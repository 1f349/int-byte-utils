package int_byte_utils

import (
	"io"
	"math"
)

const highestBit = uint((math.MaxUint + 1) >> 1)

// WriteIntAsBytes writes an integer to an io.Writer supporting negative integers
// with the sign bit as the first bit of the last byte
func WriteIntAsBytes(i int, writer io.Writer) (n int, err error) {
	if i == 0 {
		bw, err := writer.Write([]byte{0})
		return bw, err
	}
	if i < 0 {
		return WriteUintAsBytes(uint(-i)|highestBit, writer)
	} else {
		return WriteUintAsBytes(uint(i), writer)
	}
}

// ReadIntFromBytes reads an integer from an io.Reader supporting negative numbers
// with the sign bit as the first bit of the last byte
func ReadIntFromBytes(r io.Reader) (n int, err error, val int) {
	var uval uint
	n, err, uval = ReadUintFromBytes(r)
	if uval == highestBit { // Negative 0 aka math.MinInt aka -highestBit
		return n, err, math.MinInt
	} else if (uval & highestBit) == 0 {
		return n, err, int(uval)
	} else {
		return n, err, int(-(uval & ^highestBit))
	}
}
