package types

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

// HexToBytes decodes a hexadecimal string to byte array.
func HexToBytes(h string) []byte {
	var str = h
	if !strings.Contains(h, "0x") {
		str = fmt.Sprintf("0x%s", h)
	}
	data, err := hexutil.Decode(str)
	if err != nil {
		panic(err)
	}
	return data
}

// BytesToHex returns the hexadecimal(with 0x) encoding of src.
func BytesToHex(b []byte) string {
	return fmt.Sprintf("0x%s", hex.EncodeToString(b))
}

func PackByte(v byte) *Byte {
	b := NewByte(v)
	return &b
}

func PackBytes(v []byte) *Bytes {
	builder := NewBytesBuilder()
	for _, vv := range v {
		builder.Push(*PackByte(vv))
	}
	b := builder.Build()
	return &b
}

func PackUint16(v uint16) *Uint16 {
	b := make([]byte, 2)
	binary.LittleEndian.PutUint16(b, v)
	return Uint16FromSliceUnchecked(b)
}
