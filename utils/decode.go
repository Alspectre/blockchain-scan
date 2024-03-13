package utils

import (
	"encoding/base64"
	"fmt"
)

func Base64Decode(encrypt string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(encrypt)
	if err != nil {
		fmt.Println("Error decoding base64:", err)
		return "", err
	}
	return string(decoded), nil
}
