package int_byte_utils

import (
	"io"
	"math"
)

const bits6 = 63
const bit7 = 64

// WriteIntAsBytesReduced writes an integer to an io.Writer supporting negative integers
// with the sign bit as the first bit of the first byte
func WriteIntAsBytesReduced(i int, writer io.Writer) (n int, err error) {
	if i == 0 {
		bw, err := writer.Write([]byte{0})
		return bw, err
	}
	var ui uint
	if i < 0 {
		ui = uint(-i)
		return WriteUintAsBytes(ui&bits6|(ui>>6)<<7|bit7, writer)
	} else {
		ui = uint(i)
		return WriteUintAsBytes(ui&bits6|(ui>>6)<<7, writer)
	}
}

// ReadIntFromBytesReduced reads an integer from an io.Reader supporting negative numbers
// with the sign bit as the first bit of the first byte
func ReadIntFromBytesReduced(r io.Reader) (n int, err error, val int) {
	var uval uint
	n, err, uval = ReadUintFromBytes(r)
	ui := (uval & bits6) | (uval>>7)<<6
	if uval == bit7 { // Negative 0 aka math.MinInt aka bit7
		return n, err, math.MinInt
	} else if (uval & bit7) == 0 {
		return n, err, int(ui)
	} else {
		return n, err, -int(ui)
	}
}
