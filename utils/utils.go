package utils

import (
	"encoding/binary"
)

type Endianness uint
type WordOrder uint

const (
	// endianness of 16-bit registers
	BIG_ENDIAN    Endianness = 1
	LITTLE_ENDIAN Endianness = 2

	// word order of 32-bit registers
	HIGH_WORD_FIRST WordOrder = 1
	LOW_WORD_FIRST  WordOrder = 2
)

func Uint16ToBytes(endianness Endianness, in uint16) (out []byte) {
	out = make([]byte, 2)
	switch endianness {
	case BIG_ENDIAN:
		binary.BigEndian.PutUint16(out, in)
	case LITTLE_ENDIAN:
		binary.LittleEndian.PutUint16(out, in)
	}

	return
}

func Uint16sToBytes(endianness Endianness, in []uint16) (out []byte) {
	for i := range in {
		out = append(out, Uint16ToBytes(endianness, in[i])...)
	}

	return
}

func BytesToUint16(endianness Endianness, in []byte) (out uint16) {
	switch endianness {
	case BIG_ENDIAN:
		out = binary.BigEndian.Uint16(in)
	case LITTLE_ENDIAN:
		out = binary.LittleEndian.Uint16(in)
	}

	return
}

func BytesToUint16s(endianness Endianness, in []byte) (out []uint16) {
	for i := 0; i < len(in); i += 2 {
		out = append(out, BytesToUint16(endianness, in[i:i+2]))
	}

	return
}
