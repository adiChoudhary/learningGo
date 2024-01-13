package main

import (
	"fmt"
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
			existing[title] = append(existingContent, content...)
		} else {
			// Title doesn't exist, add it.
			existing[title] = content
		}
	}
	return existing
}

func main() {
	// Example usage:
	fileLinks := map[string]string{
		"ch2/temp/q10calculatingEnergy.go": "https://github.com/adiChoudhary/learningGo/blob/main/code/PracticingGo/danielLangQuestions/ch2/q10calculatingEnergy.go",
		"ch2/q13compundValue.go":           "https://github.com/adiChoudhary/learningGo/blob/main/code/PracticingGo/danielLangQuestions/ch2/q13compundValue.go",
		"ch2/q14compundValue.go":           "https://github.com/adiChoudhary/learningGo/blob/main/code/PracticingGo/danielLangQuestions/ch2/q13compundValue.go",
		// Add more file links as needed...
	}

	data := ParseFileLinks(fileLinks)

	// Generate Markdown content for the provided data.
	result := data.GenerateMarkdown(2)

	// Print the result.
	fmt.Println(result)
}
