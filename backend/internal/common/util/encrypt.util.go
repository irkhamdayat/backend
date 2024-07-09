package util

import (
	"crypto/sha512"
	"fmt"
)

func EncryptWithSalt(plainText, salt string) string {
	h := sha512.New()
	h.Write([]byte(plainText + salt))
	bs := h.Sum(nil)

	return fmt.Sprintf("%x", bs)
}
