package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func main() {
	projectDir := "../p5js-wasm-go"
	outputFile := "all/combined.go"

	functions := make(map[string][]string)
	imports := make(map[string]struct{})

	// Recursively walk all files in a project directory
	err := filepath.Walk(projectDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the file is .go
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".go") && info.Name() != "combined.go" && info.Name() != "combine.go" && !strings.Contains(path, "examples") && !strings.Contains(path, "utils") {
			// Parse file
			fs := token.NewFileSet()
			node, err := parser.ParseFile(fs, path, nil, parser.ParseComments)
			if err != nil {
				return err
			}

			// Get package name
			pkgName := node.Name.Name
			importPath := getImportPath(projectDir, path)

			// Add package to imports
			imports[importPath] = struct{}{}

			// Look for functions in AST
			ast.Inspect(node, func(n ast.Node) bool {
				if fn, ok := n.(*ast.FuncDecl); ok {
					if fn.Recv == nil && fn.Name.IsExported() { // Исключаем методы и неэкспортируемые функции
						funcName := fn.Name.Name
						functions[pkgName] = append(functions[pkgName], funcName)
					}
				}
				return true
			})
		}
		return nil
	})
	if err != nil {
		fmt.Println("Ошибка при обходе директорий:", err)
		return
	}

	// Write everything to a file
	var output []string
	output = append(output, "// Code generated by a script to simplify the import process.")
	output = append(output, "// This file aggregates all the functions from various packages in the library")
	output = append(output, "// into a single package-level file. This allows users to import only this")
	output = append(output, "// package instead of importing multiple sub-packages.")
	output = append(output, "// Authors of source code: Sanya239, Nosorozhek")
	output = append(output, "// This file and comment are generated by ChatGPT")
	output = append(output, "")
	output = append(output, "package all")
	output = append(output, "")

	output = append(output, "import (")
	sortedImp := make([]string, 0, len(imports))
	for imp := range imports {
		sortedImp = append(sortedImp, imp)
	}
	sort.Strings(sortedImp)
	for _, imp := range sortedImp {
		output = append(output, fmt.Sprintf("\t\"%s\"", imp))
	}
	output = append(output, ")")
	output = append(output, "")

	for pkg, funcs := range functions {
		for _, fn := range funcs {
			output = append(output, fmt.Sprintf("var %s = %s.%s", fn, pkg, fn))
		}
		output = append(output, "")
	}

	err = ioutil.WriteFile(outputFile, []byte(strings.Join(output, "\n")), 0644)
	if err != nil {
		fmt.Println("Ошибка при записи в файл:", err)
		return
	}

	fmt.Println("Функции успешно собраны в файл:", outputFile)
}

// getImportPath returns package's path
func getImportPath(baseDir, filePath string) string {
	relPath, err := filepath.Rel(baseDir, filePath)
	if err != nil {
		return ""
	}
	dir := filepath.Dir(relPath)
	fmt.Println(baseDir, filePath, dir)
	return "github.com/aHaHaTeam/p5js-wasm-go/" + filepath.ToSlash(dir)
}
