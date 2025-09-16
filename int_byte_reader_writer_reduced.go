package int_byte_utils

import (
	"io"
	"math"
)

// WriteIntAsBytesReduced writes an integer to an io.Writer supporting negative integers
// with the sign bit as the last bit of the first byte
func WriteIntAsBytesReduced(i int, writer io.Writer) (n int, err error) {
	if i == 0 {
		bw, err := writer.Write([]byte{0})
		return bw, err
	}
	var ui uint
	if i < 0 {
		ui = uint(-i)
		return WriteUintAsBytes(ui<<1|1, writer)
	} else {
		ui = uint(i)
		return WriteUintAsBytes(ui<<1, writer)
	}
}

// ReadIntFromBytesReduced reads an integer from an io.Reader supporting negative numbers
// with the sign bit as the last bit of the first byte
func ReadIntFromBytesReduced(r io.Reader) (n int, err error, val int) {
	var uval uint
	n, err, uval = ReadUintFromBytes(r)
	if uval == 1 { // Negative 0 aka math.MinInt aka 1
		return n, err, math.MinInt
	} else if (uval & 1) == 0 {
		return n, err, int(uval >> 1)
	} else {
		return n, err, -int(uval >> 1)
	}
}
