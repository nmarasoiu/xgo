// +build ignore

package main

import (
	"log"
	"text/template"

	"../gen"
)

func main() {
	t, err := template.New("").Parse(tmpl)
	if err != nil {
		log.Fatal(err)
	}
	var j js
	if err := gen.Gen("bob.json", &j, t); err != nil {
		log.Fatal(err)
	}
}

// The JSON structure we expect to be able to unmarshal into
type js struct {
	Cases []struct {
		Description string
		Input       string
		Expected    string
	}
}

// template applied to above data structure generates the Go test cases
var tmpl = `package bob

// Source: {{.Ori}}
{{if .Commit}}// Commit: {{.Commit}}
{{end}}

var testCases = []struct {
	description string
	input       string
	expected    string
}{
{{range .J.Cases}}{
	{{printf "%q" .Description}},
	{{printf "%q" .Input}},
	{{printf "%q" .Expected}},
},
{{end}}}
`
