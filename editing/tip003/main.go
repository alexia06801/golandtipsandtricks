package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	// Code completion works in language injections too!

	// Step 1. Inject (Alt+Enter | Inject Language | JSON) below
	json := ""
	fmt.Println(json)

	// Step 2. Inject Go Template below
	//  and type something like "各位 {{.Message}} 晚上好"
	message := ``
	printMessage(message, "Gopher")
}

type output struct {
	Message string
}

func printMessage(msg, name string) {
	t := template.Must(template.New("output").Parse(msg))

	o := output{Message: name}

	err := t.Execute(os.Stdout, o)
	if err != nil {
		log.Fatalln(err)
	}
}
