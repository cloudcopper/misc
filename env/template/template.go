package template

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	text "text/template"
)

// Data map is the template data storage type.
// During execution has captured variables
// and the SHELL is replaced by function.
type Data map[string]interface{}

// Template struct is object which keeps parsed text/template,
// and captured environment.
type Template struct {
	Template *text.Template
	Data     Data
}

// MustExec is a simplest one-liner to create, execute and get out the result of template.
// It panic in case of any error.
func MustExec(s string) string {
	t, err := Parse(s)
	if err != nil {
		panic(err)
	}
	r, err := t.Execute()
	if err != nil {
		panic(err)
	}

	return r
}

// MustExecFile is a simplest one-liner to create from file, execute and get out the result of template.
// It panic in case of any error.
func MustExecFile(filename string) string {
	s, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	return MustExec(string(s))
}

// Execute applies to parsed template the captured environment.
func (t *Template) Execute() (string, error) {
	buf := &bytes.Buffer{}
	err := t.Template.Execute(buf, t.Data)
	return buf.String(), err

}

// Parse create new Template struct out of string.
// The set of environment variables captured at the moment of creation.
func Parse(s string) (*Template, error) {
	// create template
	t := text.New("")
	t, err := t.Parse(s)
	if err != nil {
		return nil, err
	}

	return create(t)
}

// ParseFiles parses the named files and return new Template struct.
// The set of environment variables captured at the moment of creation.
func ParseFiles(filenames ...string) (*Template, error) {
	// create template
	t, err := text.ParseFiles(filenames...)
	if err != nil {
		return nil, err
	}

	return create(t)
}

func create(t *text.Template) (*Template, error) {
	// create data out of environment and shell function
	data := Data{}
	env := os.Environ()
	for _, e := range env {
		a := strings.SplitN(e, "=", 2)
		if len(a) < 2 {
			continue
		}

		k := a[0]
		if k == "" {
			continue
		}
		v := a[1]
		data[k] = v
	}

	shell := os.Getenv("SHELL")
	data["SHELL"] = func(cmd string) (string, error) {
		out, err := exec.Command(shell, "-c", cmd).Output()
		s := string(out)
		str := strings.TrimSuffix(s, "\n")
		return str, err
	}

	return &Template{t, data}, nil
}
