package pkg

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Extracts Go function definitions
func ExtractGoFunctions(filePath string) map[string]string {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	// Regex to detect function definitions
	funcRegex := regexp.MustCompile(`func\s+(\w+)\s*\(.*\)\s*{`)

	functions := make(map[string]string)
	var functionName, functionBody string
	var inFunction bool
	openBraces := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Detect function start
		if match := funcRegex.FindStringSubmatch(line); match != nil {
			functionName = match[1]
			functionBody = line + "\n"
			inFunction = true
			openBraces = 1
			continue
		}

		// Collect function body
		if inFunction {
			functionBody += line + "\n"
			openBraces += strings.Count(line, "{") - strings.Count(line, "}")

			// Function ends when all braces are closed
			if openBraces == 0 {
				functions[functionName] = functionBody
				inFunction = false
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	return functions
}

func ExtractPythonFunctions(filePath string) map[string]string {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	funcRegex := regexp.MustCompile(`^def\s+(\w+)\s*\(([^)]*)\):`)
	functions := make(map[string]string)

	var functionName, functionBody string
	var inFunction bool
	var indentLevel int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if match := funcRegex.FindStringSubmatch(line); match != nil {
			functionName = match[1]
			functionBody = line + "\n"
			inFunction = true
			indentLevel = strings.Index(line, "def")
			continue
		}

		if inFunction {
			// Check indentation to determine end of function
			currentIndent := len(line) - len(strings.TrimLeft(line, " "))
			if currentIndent > indentLevel || line == "" {
				functionBody += line + "\n"
			} else {
				functions[functionName] = functionBody
				inFunction = false
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	return functions
}

func ExtractJSFunctions(filePath string) map[string]string {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer file.Close()

	funcRegex := regexp.MustCompile(`function\s+(\w+)\s*\(.*\)\s*{|\s*(\w+)\s*=\s*\(.*\)\s*=>`)
	functions := make(map[string]string)

	var functionName, functionBody string
	var inFunction bool
	openBraces := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if match := funcRegex.FindStringSubmatch(line); match != nil {
			if match[1] != "" {
				functionName = match[1]
			} else {
				functionName = match[2]
			}
			functionBody = line + "\n"
			inFunction = true
			openBraces = 1
			continue
		}

		if inFunction {
			functionBody += line + "\n"
			openBraces += strings.Count(line, "{") - strings.Count(line, "}")

			if openBraces == 0 {
				functions[functionName] = functionBody
				inFunction = false
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
	return functions
}
