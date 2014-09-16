package util

import (
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/mitchellh/go-homedir"
)

// sha1 hash func.
// func Hash(content string) (hash string) {
// 	h := sha1.New()
// 	h.Write([]byte(content))
// 	sum := h.Sum(nil)
//
// 	hash = fmt.Sprintf("%x", sum)
// 	return
// }

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

func Reverse(list []string) []string {
	var newList []string

	for i := len(list) - 1; i >= 0; i-- {
		newList = append(newList, list[i])
	}

	return newList
}
