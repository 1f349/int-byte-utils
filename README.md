# Int Byte Utils

Utilities for converting (u)integers to and from a slice of bytes; supports the standard io.Writer and io.Reader.
This is almost the same as varint provided in the standard encoding/binary package with different interface targets
and that only up to 64 bits is supported allowing for 9 byte uint max storage rather than 10 bytes.

## Format
All byte data is treated as unsigned integers by default, 
the most significant bit of each byte signifies another byte follows, 
the rest of the bits are in order within each byte, representing the most significant to the least significant bit, 
however, the bytes are in the reverse order where the least significant bits are in the first byte and the most significant in the last byte.
This allows for a full uinteger to be stored without explicitly recording the length of the structure.
However, when the most significant bit of a uint64 is in use, rather than extend for another byte, the most significant bit is used for storage
allowing for a 9 byte max size as opposed to 10 bytes sacrificing support for 128 bit uints but saving space.

### EG:
65537 = [129 128 4] = [~~1~~0000001 ~~1~~0000000 ~~0~~0000100] -> [1 0 4]

The first byte is 64-1 so 2^0=1

The second byte is 8192-128 so 0

The third byte is 1048576-16384 so 2^16=65536

1 + 0 + 65536 = 65537

## Integer Format
The standard format utilizes the most significant bit of the uinteger to represent the sign byte, 
where it being set means negative, this is the standard overflow conversion behaviour in C between signed and unsigned.
This bit would be located in the last byte of the written data as the second most significant byte.

There is also a reduced format that moves the sign bit to the least significant bit of the first byte, 
the 1st bit of the uinteger. This allows for the negative numbers to have a smaller footprint similar to their positive counterparts 
as the sign bit does not have to be at the most significant bit of a uint.

## License
BSD 3-Clause - (C) 1f349 2025
