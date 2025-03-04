package pkg

import (
	"bufio"
	"fmt"
	"os"
)

// FunctionData stores extracted function details
type FunctionData struct {
	Name        string
	Parameters  string
	ReturnType  string
	Description string
}

// GenerateMarkdown creates an API documentation file.
func GenerateMarkdown(functions []FunctionData, outputFile string) {
	file, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	writer.WriteString("# API Documentation\n\n")

	for _, fn := range functions {
		writer.WriteString(fmt.Sprintf("## Function: `%s(%s) -> %s `\n", fn.Name, fn.Parameters, fn.ReturnType))
		writer.WriteString("- **Description:** _To be updated_\n")
		writer.WriteString("- **Parameters:**\n")
		if fn.Parameters == "" {
			writer.WriteString("  - None\n")
		} else {

			writer.WriteString(fmt.Sprintf("- **Returns:** `%s`\n\n", fn.ReturnType))
		}
		writer.Flush()
		fmt.Println("📝 API documentation generated successfully: ", outputFile)

	}
}
