// Package template is the wrapper around text/template to support access
// to environment variables and piping shell commands.
//
// Template example:
//    Hello {{ .USER }} at {{ call .SHELL "hostname" }}!!!
//    Your working directory name is {{ call .SHELL "basename $(pwd)" }}.
//    Your IP addresses are:
//    {{ call .SHELL "ifconfig|grep 'inet addr'"}}
//    Now is {{ call .SHELL "date" }}
//
// For simplest usage see functions Exec and ExecFile.
//
// Basic usage:
//
//    import tpl "github.com/cloudcopper/misc/env/template"
//    ...
//    t, err := tpl.Parse(`Hello {{ .USER }} !!! Now is {{ call .SHELL "date" }}`)
//    if err != nil {
//       return err
//    }
//
//    s, err := t.Execute()
//    if err != nil {
//       return err
//    }
//
//    fmt.Printf(str)
//
package template
