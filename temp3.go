package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/template"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	templates := map[string]string{
		"Welcome":      "Welcome, {{.name}} ! We are glad that you joined.",
		"Notification": "{{.name}}, you have a new notifcation: {{.notification}}",
		"Error":        "Oops! An Error occured: {{.errorMessage}}",
	}

	parsedTemplates := make(map[string]*template.Template)

	for name, tmpl := range templates {
		parsedTemplates[name] = template.Must(template.New(name).Parse(tmpl))
	}

	for {
		fmt.Println("\nMenu:")
		fmt.Println("1. Join")
		fmt.Println("2. Get Notfication")
		fmt.Println("3. Get Error")
		fmt.Println("4. Exit")
		fmt.Println("Enter Choices: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		var data map[string]interface{}
		var tmpl *template.Template

		switch choice {
		case "1":
			tmpl = parsedTemplates["Welcome"]
			data = map[string]interface{}{"name": name}
		case "2":
			fmt.Println("Enter your notifcation nessage: ")
			notification, _ := reader.ReadString('\n')
			notification = strings.TrimSpace(notification)
			tmpl = parsedTemplates["Notification"]
			data = map[string]interface{}{"name": name, "notification": notification}
		case "3":
			fmt.Println("Enter your error nessage: ")
			errorMessage, _ := reader.ReadString('\n')
			errorMessage = strings.TrimSpace(errorMessage)
			tmpl = parsedTemplates["Error"]
			data = map[string]interface{}{"name": name, "errorMessage": errorMessage}

		case "4":
			fmt.Println("Exiting.........")
			return
		default:
			fmt.Println("Invalid Choice Please select a valid option. ")
			continue

		}

		err := tmpl.Execute(os.Stdout, data)
		if err != nil {
			fmt.Println("Error in Executing template: ", err)
		}
	}
}
