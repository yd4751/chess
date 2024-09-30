package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"text/template"
)

const initTemplate = `package {{.}}
 
import "fmt"
 
func init() {
    // 这里可以添加初始化代码
    fmt.Printf("Initialized package: {{.}}\\n")
}
`

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run generate_init.go <directory>")
		os.Exit(1)
	}

	dir := os.Args[1]
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Printf("Error reading directory: %s\n", err)
		os.Exit(1)
	}

	for _, file := range files {
		if !file.IsDir() {
			continue
		}

		pkgPath := path.Join(dir, file.Name())
		initFile := path.Join(pkgPath, "init.go")

		if _, err := os.Stat(initFile); err == nil {
			// 文件已存在，跳过
			continue
		}

		err = os.MkdirAll(pkgPath, 0755)
		if err != nil {
			fmt.Printf("Error creating directory: %s\n", err)
			os.Exit(1)
		}

		f, err := os.Create(initFile)
		if err != nil {
			fmt.Printf("Error creating file: %s\n", err)
			os.Exit(1)
		}
		defer f.Close()

		t := template.Must(template.New("init").Parse(initTemplate))
		err = t.Execute(f, file.Name())
		if err != nil {
			fmt.Printf("Error executing template: %s\n", err)
			os.Exit(1)
		}
	}
}
