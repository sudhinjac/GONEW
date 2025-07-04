package main

import (
	"os"
	"text/template"
)

func main() {

	tmpl := template.New("example")
	tmpl, err := tmpl.Parse("Welcome,{{.name}}! how are you doing?\n")
	if err != nil {
		panic(err)
	}

	data := map[string]interface{}{
		"name": "john",
	}

	err = tmpl.Execute(os.Stdout, data)
	if err != nil {
		panic(err)
	}

}
