package Hex

import (
	"encoding/hex"
	"fmt"
	"strings"
)

func Normal2Hex(input string) string {
	encoded := hex.EncodeToString([]byte(input))
	fmt.Println(encoded)
	return encoded
}

func Hex2Normal(input string) string {
	input = strings.TrimSpace(input)
	decoded, err := hex.DecodeString(input)
	if err != nil {
		error, _ := fmt.Println("decode error: ", err)
		return string(error)
	} else if err == nil {
		fmt.Println(string(decoded))
		return string(decoded)
	}
	return "hi"
}
