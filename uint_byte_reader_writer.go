package int_byte_utils

import (
	"errors"
	"io"
	"math"
)

// OverflowError the read value has overflowed the container.
var OverflowError = errors.New("overflow")

const bits7 = 127
const bit8 = 128

// WriteUintAsBytes writes an integer to an io.Writer
func WriteUintAsBytes(i uint, writer io.Writer) (n int, err error) {
	if i == 0 {
		bw, err := writer.Write([]byte{0})
		return bw, err
	}
	n = 0
	currentI := i
	for currentI > 0 {
		var bt = byte(currentI & bits7)
		currentI = currentI >> 7
		if currentI > 0 {
			bt |= bit8
		}
		cbw, err := writer.Write([]byte{bt})
		n += cbw
		if err != nil {
			return n, err
		}
	}
	return n, nil
}

// ReadUintFromBytes reads an integer from an io.Reader
func ReadUintFromBytes(r io.Reader) (n int, err error, val uint) {
	cBuff := make([]byte, 1)
	val = 0
	cBitSize := 0
	n = 0
	for val < math.MaxUint {
		br, err := io.ReadFull(r, cBuff)
		n += br
		if err != nil {
			return n, err, val
		}
		if cBuff[0] < bit8 {
			if val > val|uint(uint(cBuff[0])<<cBitSize) {
				return n, OverflowError, math.MaxUint
			}
			return n, nil, val | uint(uint(cBuff[0])<<cBitSize)
		}
		val |= uint(cBuff[0]&bits7) << cBitSize
		cBitSize += 7
	}
	return n, OverflowError, val
}
