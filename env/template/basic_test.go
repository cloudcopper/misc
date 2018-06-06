package template

import (
	"testing"
)

func TestSimplest(t *testing.T) {
	s := `
  Hello {{ .USER }} at {{ call .SHELL "hostname" }}!!!
  Your working directory name is {{ call .SHELL "basename $(pwd)" }}.
  Your IP addresses are:
  {{ call .SHELL "ifconfig|grep 'inet addr'"}}
  Now is {{ call .SHELL "date" }}
  `

	tpl, err := Parse(s)
	if err != nil {
		t.Fatal(err)
	}
	str, err := tpl.Execute()
	if err != nil {
		t.Fatal(err)
	}

	t.Log(str)
}
