package libs

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func Unique(target []string) (unique []string) {
	m := map[string]bool{}

	for _, v := range target {
		if !m[v] {
			m[v] = true
			unique = append(unique, v)
		}
	}

	return unique
}

func ConvertToDirectory(namespace string) string {
	return strings.ReplaceAll(strings.Trim(namespace, "\\"), "\\", "/")
}

func ConvertToNamespace(directory string) string {
	return strings.ReplaceAll(strings.Trim(strings.Trim(directory, "."), "/"), "/", "\\")
}

func Indent(i int) string {
	return strings.Repeat(" ", i)
}

func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}

func EchoDoubleLine() {
	fmt.Println(strings.Repeat("=", 30))
}

func EchoSingleLine() {
	fmt.Println(strings.Repeat("-", 30))
}

func EchoNewLine() {
	fmt.Println()
}

func EchoSection(sectionName string) {
	EchoDoubleLine()
	fmt.Println(sectionName)
	EchoDoubleLine()
	EchoNewLine()
}

func ConvertMapForDisplay(targetMap map[string]string) (forDisplay []string) {
	for key := range targetMap {
		forDisplay = append(forDisplay, key)
	}
	sort.Strings(forDisplay)

	return forDisplay
}

func Ask(ask string) (answer string) {
	fmt.Print(strings.Trim(ask, " ") + " ")
	fmt.Scanln(&answer)
	EchoNewLine()

	return answer
}

func AskedRepetition(ask string) (answer string) {
	trimedAsk := strings.Trim(ask, " ")
	for answer == "" {
		fmt.Print(trimedAsk + " ")
		fmt.Scanln(&answer)
		EchoNewLine()
	}

	return answer
}
