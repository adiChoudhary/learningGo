package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func generateREADME(mp interface{}, depth int) {
	if ele, ok := mp.(map[string]interface{}); ok {
		keys := make([]string, len(ele))
		i := 0
		for k := range ele {
			keys[i] = k
			i++
		}
		for _, value := range keys {
			fmt.Printf("%s %s\n", strings.Repeat("#", depth), value)
			generateREADME(ele[value], depth+1)
		}
	} else if ele, ok := mp.([][]string); ok {
		for _, value := range ele {
			fmt.Printf("[%s](https://github.com/adiChoudhary/learningGo/blob/main/code/PracticingGo/%s)\n\n", value[0], value[1])
		}
	}
}

func main() {
	rootDir := "/home/adi/Documents/coding/learningGo/code/PracticingGo" // Change this to your project root directory
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
		for index, value := range temp {
			if strings.Contains(value, ".go") {
				continue
			}
			if _, ok := currentMp[value]; (!ok) && (len(temp)-2) != index {
				currentMp[value] = make(map[string]interface{})
			} else if (!ok) && (len(temp)-2) == index {
				currentMp[value] = [][]string{[]string{temp[len(temp)-1], relPath}}
			} else if (ok) && (len(temp)-2) == index {
				currentMp[value] = append(currentMp[value].([][]string), []string{temp[len(temp)-1], relPath})
			}
			// Here type assertion is used to tell go a specific type since
			// interface isn't a type
			if (len(temp) - 2) != index {
				currentMp = currentMp[value].(map[string]interface{})
			}
		}
	}
	generateREADME(mp, 2)
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
