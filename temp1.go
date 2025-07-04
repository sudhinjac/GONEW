package main

import (
	"os"
	"text/template"
)

func main() {
	// Define template string
	tmpl := `Hello, {{.Name}}! Welcome to {{.City}}.`

	// Parse the template
	t, err := template.New("welcome").Parse(tmpl)
	if err != nil {
		panic(err)
	}

	// Data to inject into template
	data := map[string]string{
		"Name": "Sudhin",
		"City": "Bangalore",
	}

	// Execute template and write output to stdout
	t.Execute(os.Stdout, data)
}
