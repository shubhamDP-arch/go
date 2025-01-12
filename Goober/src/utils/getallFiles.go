package utils

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func GetAllfilesInDirectory(dir string) []string {
	var files []string
	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		fmt.Println("Error walking the directory")
	}
	return files

}

func IsDirectory(file string) bool {
	fileInfo, err := os.Stat(file)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}
