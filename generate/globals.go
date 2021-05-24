package main

import (
	"bytes"
	"fmt"
	"os"
	"text/template"
)

// Globals represent SAM template globals: these are essentially resources but properties as direct children rather than within a "Properties" field.
type Globals struct {

	// Documentation is a link to the AWS SAM User Guide for information about the resource.
	Documentation string `json:"Documentation"`

	// Children of Globals are essentially duplicates of SAM resources with some fields exclude.
	// https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/sam-specification-template-anatomy-globals.html
	Children map[string]Global
}

// A Global is a duplicate of a Serverless resource but with some fields excluded
type Global struct {
	Reference string

	Exclude []string
}

// Schema returns a JSON Schema for the global (as a string)
func (g Global) Schema(name string, r map[string]Resource) string {

	// Open the schema template and setup a counter function that will
	// available in the template to be used to detect when trailing commas
	// are required in the JSON when looping through maps
	tmpl, err := template.New("schema-globals.template").Funcs(template.FuncMap{
		"counter": counter,
	}).ParseFiles("generate/templates/schema-globals.template")

	var buf bytes.Buffer

	properties := make(map[string]Property)
	for k, v := range r[g.Reference].Properties {
		if !g.isExcluded(k) {
			properties[k] = v
		}
	}

	templateData := struct {
		Name       string
		RefName    string
		Properties map[string]Property
	}{
		Name:       name,
		RefName:    g.Reference,
		Properties: properties,
	}

	// Execute the template, writing it to the buffer
	err = tmpl.Execute(&buf, templateData)

	if err != nil {
		fmt.Printf("Error: Failed to generate global %s\n%s\n", name, err)
		os.Exit(1)
	}

	return buf.String()

}

// Determines whether or not a field is excluded from the Global's implementation
func (g Global) isExcluded(f string) bool {
	for _, v := range g.Exclude {
		if v == f {
			return true
		}
	}

	return false
}
