package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// MarkdownGenerator is an interface that defines a GenerateMarkdown method.
type MarkdownGenerator interface {
	GenerateMarkdown(depth int) string
}

// Section represents a section in the Markdown document.
type Section map[string][]MarkdownGenerator

// Link represents a Markdown link.
type Link struct {
	Text string
	URL  string
}

// GenerateMarkdown is a method associated with the Section type.
func (s Section) GenerateMarkdown(depth int) string {
	var result string
	for title, content := range s {
		result += fmt.Sprintf("%s %s\n", strings.Repeat("#", depth), title)
		for _, item := range content {
			result += item.GenerateMarkdown(depth + 1)
		}
	}
	return result
}

// GenerateMarkdown is a method associated with the Link type.
func (l Link) GenerateMarkdown(depth int) string {
	return fmt.Sprintf("[%s](%s)\n\n", l.Text, l.URL)
}

// ParseFileLinks converts a list of file links to the hierarchical structure.
func ParseFileLinks(fileLinks map[string]string) Section {
	result := make(Section)
	for title, link := range fileLinks {
		components := strings.Split(title, "/")
		result = mergeSections(result, createHierarchy(components, Link{Text: title, URL: link}))
	}
	return result
}

func createHierarchy(components []string, generator MarkdownGenerator) Section {
	if len(components) == 1 {
		return Section{components[0]: []MarkdownGenerator{generator}}
	}
	return Section{components[0]: []MarkdownGenerator{createHierarchy(components[1:], generator)}}
}

func mergeSections(existing, newSection Section) Section {
	for title, content := range newSection {
		if existingContent, ok := existing[title]; ok {
			// Title already exists, merge the content.
			// TODO correct this logic to resolve incorrect merging issue
			existing[title] = append(existingContent, content...)
		} else {
			// Title doesn't exist, add it.
			existing[title] = content
		}
	}
	return existing
}

func generateFileLinks(filePaths []string, baseUrl string) map[string]string {
	mp := make(map[string]string)
	for _, value := range filePaths {
		mp[value] = baseUrl + value
	}
	return mp
}

func main() {
	rootDir := "/home/adi/Documents/coding/learningGo/code/PracticingGo/basics" // Change this to your project root directory
	filePaths, err := getAllGoFiles(rootDir)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fileLinks := generateFileLinks(filePaths, "https://github.com/adiChoudhary/learningGo/blob/main/code/PracticingGo/")
	data := ParseFileLinks(fileLinks)
	fmt.Println(data)
	// Generate Markdown content for the provided data.
	//result := data.GenerateMarkdown(2)

	// Print the result.
	//fmt.Println(result)
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
	for index, value := range goFiles {
		goFiles[index] = strings.Replace(value, "/home/adi/Documents/coding/learningGo/code/PracticingGo/", "", 1)
	}

	return goFiles, err
}
