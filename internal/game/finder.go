package game

import (
	"errors"
	"os"
	"path/filepath"
)

var cachePath string = ""

// Find CrossCode from some dir to base dir
func Find(dir string) (string, error) {
	if cachePath != "" {
		return cachePath, nil
	}

	hasGame := hasCrossCode(dir)

	if hasGame {

		dir = filepath.Join(dir, "assets", ".")
		cachePath = dir
		return dir, nil
	}

	parentDir := filepath.Dir(dir)

	if parentDir == dir {
		return "", errors.New("game/finder: Could not find Game Path")
	}

	return Find(parentDir)
}

// GetModDirectory h
func GetModDirectory(dir string) string {

	moddir := filepath.Join(dir, "mods", ".")

	if !folderExists(moddir) {
		return ""
	}

	return moddir
}

func hasCrossCode(dir string) bool {

	packagePath := filepath.Join(dir, "package.json")

	if !fileExists(packagePath) {
		return false
	}

	mainHTML := filepath.Join(dir, "assets", "node-webkit.html")
	if !fileExists(mainHTML) {
		return false
	}

	return true
}

func fileExists(filePath string) bool {
	_, err := os.Stat(filePath)

	if os.IsNotExist(err) {
		return false
	}
	return true
}

func folderExists(folderPath string) bool {
	stat, err := os.Stat(folderPath)

	if os.IsNotExist(err) {
		return false
	}

	if !stat.IsDir() {
		return false
	}
	return true
}
