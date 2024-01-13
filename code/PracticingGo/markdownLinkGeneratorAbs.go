package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func generateREADME(mp interface{}, root string, depth int) {
	switch val := mp.(type) {
	case map[string]interface{}:
		fmt.Printf("%s %s\n", strings.Repeat("#", depth), root)
		for key, value := range val {
			generateREADME(value, key, depth+1)
		}
	case []string:
		fmt.Printf("[%s](https://github.com/adiChoudhary/learningGo/blob/main/code/PracticingGo/%s)\n\n", val[0], val[1])
	}
}

func main() {
	rootDir := "/home/adi/Documents/coding/learningGo/code/PracticingGo/" // Change this to your project root directory
	filePaths, err := getAllGoFiles(rootDir)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	mp := make(map[string]interface{})

	for _, filePath := range filePaths {
		relPath, _ := filepath.Abs(filePath)
		relPath = strings.Replace(relPath, "/home/adi/Documents/coding/learningGo/code/PracticingGo/", "", 1)
		temp := strings.Split(relPath, "/")
		currentMp := mp
		for _, value := range temp {
			if strings.Contains(value, ".go") {
				currentMp[value] = []string{value, relPath}
				continue
			}
			if _, ok := currentMp[value]; !ok {
				currentMp[value] = make(map[string]interface{})
			}
			currentMp = currentMp[value].(map[string]interface{})
		}
	}
	generateREADME(mp, "Go Practice Questions", 2)
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
