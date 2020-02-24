package utils

import (
	"crypto/sha1"
	"fmt"
	"io"
)

func ToSha1(s string) string {
	h := sha1.New()
	_, err := io.WriteString(h, s)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", h.Sum(nil))
}
