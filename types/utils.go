package types

import (
	"encoding/hex"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

func HexToBytes(h string) []byte {
	if strings.Contains(h, "0x") {
		data, _ := hexutil.Decode(h)
		return data
	}
	data, _ := hexutil.Decode(fmt.Sprintf("0x%s", h))
	return data
}

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
