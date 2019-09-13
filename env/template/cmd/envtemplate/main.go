package main

import (
	"os"

	"github.com/cloudcopper/misc/env/template"
)

const helpText = `Usage: envtemplate <file>

Process text template according to env/template.
The env/template is golang text/template supporting access to
environment variables and able to call shell.

See @https://github.com/cloudcopper/misc/tree/master/env/template.
`

func main() {
	args := os.Args[1:]
	if len(args) != 1 {
		print(helpText)
		os.Exit(1)
	}

	fileName := args[0]
	out := template.MustExecFile(fileName)
	print(out)
}

func print(s string) {
	os.Stdout.WriteString(s)
}
