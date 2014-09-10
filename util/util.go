package util

import (
	"crypto/sha1"
	"fmt"
)

// sha1 hash func.
func Hash(content string) (hash string) {
	h := sha1.New()
	h.Write([]byte(content))
	sum := h.Sum(nil)

	hash = fmt.Sprintf("%x", sum)
	return
}
