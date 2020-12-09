package utils

import (
	"fmt"
	"path/filepath"
)

func GetFileWithoutExtension(filename string) string {
	var extension = filepath.Ext(filename)
	return filename[0 : len(filename)-len(extension)]
}

func GetCleanBase(path string) (string, error) {
	// NOTE: don't trust the path even if it came from fsnotify
	cleanPath := filepath.Clean(path)
	if cleanPath == "" {
		return "", fmt.Errorf("Invalid path: %s", path)
	}

	// NOTE: Still not trusting that path. Let's just use the base
	// and use our configured base path
	return filepath.Base(cleanPath), nil
}

func GetSafePath(modulePath, path string) (string, error) {
	policyFileBase, err := GetCleanBase(path)
	if err != nil {
		return "", err
	}
	policyPath := filepath.Join(modulePath, policyFileBase)
	return policyPath, nil
}