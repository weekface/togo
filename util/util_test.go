package util

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/bmizerany/assert"
	"github.com/mitchellh/go-homedir"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

// exists returns whether the given file or directory exists or not
func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func TestHash(t *testing.T) {
	hash := Hash()
	assert.Equal(t, len(hash), 10)
}

func TestInitializeTogoDir(t *testing.T) {
	InitializeTogoDir()
	dir, _ := homedir.Dir()

	exist, _ := exists(filepath.Join(dir, ".togo/new"))

	assert.Equal(t, exist, true)
}
