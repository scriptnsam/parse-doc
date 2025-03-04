/*
Copyright © 2025 SCRIPTNSAM SCRIPTNSAM@DUCK.COM
*/

package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/scriptnsam/parse-doc/pkg"
	"github.com/scriptnsam/parse-doc/utils"
	"github.com/spf13/cobra"
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate API documentation from source code",
	Long:  `This is the part for the longer description`,
	Run: func(cmd *cobra.Command, args []string) {
		// prevent index out-of-range error
		if len(args) < 1 {
			fmt.Println("Error: Missing source directory. Usage: parse-doc generate <source>")
			os.Exit(1)
		}

		source := args[0]
		if _, err := os.Stat(source); os.IsNotExist(err) {
			fmt.Println("Error: Source directory not found!")
			os.Exit(1)
		}

		fmt.Printf("Parsing code from %s\n", source)

		files := []string{}

		err := filepath.Walk(source, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			//only add .go, .py, .js files
			if !info.IsDir() && (strings.HasSuffix(path, ".go") || strings.HasSuffix(path, ".py") || strings.HasSuffix(path, ".js")) {
				files = append(files, path)
			}
			return nil
		})

		if err != nil {
			fmt.Println("Error reaading directory:", err)
			os.Exit(1)
		}

		if len(files) == 0 {
			fmt.Println("No source files found!")
			return
		}

		fmt.Println("Found files:")

		for _, file := range files {
			fmt.Println(" -", file)
		}

		// TODD: Call specific parsers for each file type
		// for _, file := range files {
		// 	if strings.HasSuffix(file, ".go") {
		// 		GoFunctions := pkg.ParseGoFile(file)
		// 		pkg.GenerateMarkdown(GoFunctions, "API_DOCS_GO.md")
		// 	} else if strings.HasSuffix(file, ".py") {
		// 		PyFunctions := pkg.ParsePythonFile(file)
		// 		pkg.GenerateMarkdown(PyFunctions, "API_DOCS_PY.md")
		// 		pkg.ParsePythonFile(file)
		// 	} else if strings.HasSuffix(file, ".js") {
		// 		pkg.ParseJSFile(file)
		// 	}
		// }

		for _, file := range files {
			if strings.HasSuffix(file, ".go") {
				GoFunctions := pkg.ExtractGoFunctions(file)

				for fName, fCode := range GoFunctions {
					err := utils.AIImplementation(fName, fCode, "API_DOCS_GO.md")
					if err != nil {
						fmt.Printf("Error generating docs for %s: %v\n", fName, err)
						continue
					}
				}

			} else if strings.HasSuffix(file, ".py") {
				PyFunctions := pkg.ExtractPythonFunctions(file)
				fmt.Printf("This is the PYFunctions: %s", PyFunctions)

				for fName, fCode := range PyFunctions {
					err := utils.AIImplementation(fName, fCode, "API_DOCS_PY.md")
					if err != nil {
						fmt.Printf("Error generating docs for %s: %v\n", fName, err)
						continue
					}
				}
			} else if strings.HasSuffix(file, ".js") {
				JSFunctions := pkg.ExtractJSFunctions(file)

				for fName, fCode := range JSFunctions {
					err := utils.AIImplementation(fName, fCode, "API_DOCS_JS.md")
					if err != nil {
						fmt.Printf("Error generating docs for %s: %v\n", fName, err)
						continue
					}
				}

			}
		}

	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
