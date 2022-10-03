//go:build ignore

package main

import (
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	f, err := os.Create("./generated.go")
	check(err)

	defer f.Close()
	f.WriteString("package main\n")
	f.WriteString("const compNum = " + strconv.Itoa(len(os.Args)-1) + "\n")

	for _, name := range os.Args[1:] {
		f.WriteString("var " + strings.ToLower(name) + "s [entitySize]" + name + "\n")
	}

	f.WriteString("const (\n")
	if len(os.Args) >= 2 {
		f.WriteString(`	` + strings.ToLower(os.Args[1]) + " = iota\n")
	}
	for _, name := range os.Args[2:] {
		f.WriteString(`	` + strings.ToLower(name) + "\n")
	}
	f.WriteString(")")

	f.Sync()
}
