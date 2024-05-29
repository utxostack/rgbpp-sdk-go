package types

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

// HexToBytes decodes a hex string with 0x prefix to byte array.
func HexToBytes(h string) []byte {
	if strings.Contains(h, "0x") {
		data, err := hexutil.Decode(h)
		if err != nil {
			panic(err)
		}
		return data
	}
	data, err := hexutil.Decode(fmt.Sprintf("0x%s", h))
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
