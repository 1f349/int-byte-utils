package int_byte_utils

import (
	"errors"
	"io"
)

// OverflowError the read value has overflowed the container.
var OverflowError = errors.New("overflow")

const bits7 = 127
const bit8 = 128
const platformBits = 32 << (^uint(0) >> 63) // from go math package

// LenUintAsBytes gets the length of the data representing the passed value
func LenUintAsBytes(i uint) (n int) {
	if i == 0 {
		return 1
	}
	n = 0
	currentI := i
	for currentI > 0 {
		if n == 8 { // (8+1)*8 = 56, last byte for uint64 uses most significant as storage rather than an extension signifier
			currentI = 0
		} else {
			currentI = currentI >> 7
		}
		n++
	}
	return
}

// WriteUintAsBytes writes an integer to an io.Writer
func WriteUintAsBytes(i uint, writer io.Writer) (n int, err error) {
	if i == 0 {
		bw, err := writer.Write([]byte{0})
		return bw, err
	}
	n = 0
	currentI := i
	for currentI > 0 {
		var bt byte
		if n == 8 { // (8+1)*8 = 56, last byte for uint64 uses most significant as storage rather than an extension signifier
			bt = byte(currentI)
			currentI = 0
		} else {
			bt = byte(currentI & bits7)
			currentI = currentI >> 7
			if currentI > 0 {
				bt |= bit8
			}
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
	for cBitSize <= platformBits-1 {
		br, err := io.ReadFull(r, cBuff)
		n += br
		if err != nil {
			return n, err, val
		}
		// Allow for storing a 64 bit uint within 9 bytes by sacrificing the use of the extender bit and limiting support to 64 bits
		if cBuff[0] < bit8 || cBitSize >= 56 {
			return n, nil, val | uint(uint(cBuff[0])<<cBitSize)
		}
		val |= uint(cBuff[0]&bits7) << cBitSize
		cBitSize += 7
	}
	return n, OverflowError, val
}
