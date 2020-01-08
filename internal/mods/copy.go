package mods

import (
	"io"
	"os"
	"path/filepath"
	"strings"
)

// Copy folder from src to dest
func Copy(src string, dest string) error {
	e := filepath.Walk(src, func(currentPath string, info os.FileInfo, err error) error {

		relativePath := strings.Replace(currentPath, src, "", -1)
		destPath := filepath.Join(dest, relativePath)

		if info.IsDir() {
			return copyDir(info, destPath)
		}
		return copyFile(info, currentPath, destPath)
	})

	if e != nil {
		return e
	}
	return nil
}

func copyFile(info os.FileInfo, src string, dest string) error {
	file, err := os.Open(src)

	if err != nil {
		return err
	}
	defer file.Close()

	newFile, err := os.Create(dest)

	newFile.Chmod(info.Mode())

	if err != nil {
		return err
	}

	defer newFile.Close()

	_, err = io.Copy(newFile, file)

	if err != nil {
		return err
	}

	return newFile.Sync()
}

func copyDir(info os.FileInfo, dest string) error {
	err := os.MkdirAll(dest, info.Mode())

	if err != nil {
		return err
	}
	return nil
}
