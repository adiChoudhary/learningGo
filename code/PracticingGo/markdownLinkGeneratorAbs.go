package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	rootDir := "/home/adi/Documents/coding/learningGo/code/PracticingGo" // Change this to your project root directory
	filePaths, err := getAllGoFiles(rootDir)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, filePath := range filePaths {
		relPath, _ := filepath.Abs(filePath)
		relPath = strings.Replace(relPath, "/home/adi/Documents/coding/learningGo/code/PracticingGo/", "", 1)
		temp := strings.Split(relPath, "/")
		linkText := temp[len(temp)-1]
		link := fmt.Sprintf("[%s](https://github.com/adiChoudhary/learningGo/blob/main/code/PracticingGo/%s)", linkText, relPath)
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
