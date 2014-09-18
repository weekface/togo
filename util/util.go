package util

import (
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/mitchellh/go-homedir"
)

// timestramp hash func.
func Hash() (hash string) {
	return strconv.Itoa(int(time.Now().Unix()))
}

// initialize the togo data dir.
func InitializeTogoDir() {
	dir, _ := homedir.Dir()
	togoDir := filepath.Join(dir, ".togo")
	os.MkdirAll(filepath.Join(togoDir, "new"), 0755)
	os.MkdirAll(filepath.Join(togoDir, "old"), 0755)
}
