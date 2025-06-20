package Base

import (
	"encoding/base32"
	"encoding/base64"
	"fmt"
)

func Normal2base64(input string) string {
	encoded := base64.StdEncoding.EncodeToString([]byte(input))

	return encoded
}

func Base642Normal(input string) string {
	decoded, err := base64.StdEncoding.DecodeString(input)
	if err != nil {
		error, _ := fmt.Println("decode error: ", err)
		return string(error)
	} else if err == nil {

		
		return string(decoded)
	}
	return "hi"

}

func Normal2base32(input string) string {
	str := base32.StdEncoding.EncodeToString([]byte(input))
	return str
}

func Base322Normal(input string) string {
	decoded, err := base32.StdEncoding.DecodeString(input)

	if err != nil {
		error, _ := fmt.Println("decode error: ", err)
		return string(error)
	} else if err == nil {
	
		return string(decoded)
	}
	return "hi"
}
