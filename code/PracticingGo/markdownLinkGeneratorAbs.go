package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	rootDir := "." // Change this to your project root directory
	filePaths, err := getAllGoFiles(rootDir)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, filePath := range filePaths {
		absPath, _ := filepath.Abs(filePath)
		linkText := strings.TrimPrefix(filePath, rootDir)
		linkText = strings.TrimPrefix(linkText, "/")
		linkText = strings.TrimSuffix(linkText, ".go")
		linkText = strings.Replace(linkText, "/", " / ", -1)
		linkText = strings.Replace(linkText, "_", "\\_", -1)

		link := fmt.Sprintf("[%s](<file:%s>)", linkText, absPath)
		fmt.Println(link)
	}
}

func getAllGoFiles(rootDir string) ([]string, error) {
	var goFiles []string

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") {
			goFiles = append(goFiles, path)
		}
		return nil
	})

	return goFiles, err
}
