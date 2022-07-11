package utils

import (
	"fmt"
	"os"
)

// CheckPath returns true if the given path exists
func CheckPath(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}

// CheckAndCreatePath if the path does not exist, create it
func CheckAndCreatePath(path ...string) error {
	for _, p := range path {
		if !CheckPath(p) {
			err := os.MkdirAll(p, os.ModePerm)
			if err != nil {
				fmt.Println("创建目录错误:", err)
				return err
			} else {
				fmt.Println("已创建目录:", p)
			}
		}
	}
	return nil
}
