package util

import (
	"crypto/sha1"
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-homedir"
)

// sha1 hash func.
func Hash(content string) (hash string) {
	h := sha1.New()
	h.Write([]byte(content))
	sum := h.Sum(nil)

	hash = fmt.Sprintf("%x", sum)
	return
}

// initialize the togo data dir.
func InitializeTogoDir() {
	dir, _ := homedir.Dir()
	togoDir := filepath.Join(dir, ".togo")
	os.MkdirAll(filepath.Join(togoDir, "new"), 0755)
	os.MkdirAll(filepath.Join(togoDir, "old"), 0755)
}
