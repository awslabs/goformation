package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
)

// SpecURL is the HTTP URL of the latest AWS CloudFormation Resource Specification
const SpecURL = "https://d3teyb21fexa9r.cloudfront.net/latest/gzip/CloudFormationResourceSpecification.json"

func main() {

	// Fetch the latest CloudFormation Resource Specification
	response, err := http.Get(SpecURL)
	if err != nil {
		fmt.Printf("Error: Failed to fetch AWS CloudFormation Resource Specification\n%s\n", err)
		os.Exit(1)
	}

	// Read all of the retrieved data at once (~70KB)
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error: Failed to read AWS CloudFormation Resource Specification\n%s\n", err)
		os.Exit(1)
	}

	// Unmarshall the JSON specification data to objects
	spec := &CloudFormationResourceSpecification{}
	if err := json.Unmarshal(data, spec); err != nil {
		fmt.Printf("Error: Failed to parse AWS CloudFormation Resource Specification\n%s\n", err)
		os.Exit(1)
	}

	// Write all of the resources, using a template
	for name, resource := range spec.Resources {
		fmt.Printf("Generating Resource: %s\n", name)
		generate(name, resource, spec)
	}

	// Write all of the custom properties, using a template
	for name, property := range spec.Properties {
		fmt.Printf("Generating Custom Property: %s\n", name)
		generate(name, property, spec)
	}

}

func generate(name string, resource Resource, spec *CloudFormationResourceSpecification) {

	// Open the resource template
	tmpl, err := template.ParseFiles("resource.template")
	if err != nil {
		fmt.Printf("Error: Failed to load resource template\n%s\n", err)
		os.Exit(1)
	}

	// Create the file for the generated struct
	fn := "generated/" + filename(name)
	f, err := os.Create(fn)
	if err != nil {
		fmt.Printf("Error: Failed to create %s\n%s\n", fn, err)
		os.Exit(1)
	}
	defer f.Close()

	// Pass in the following information into the template
	templateData := struct {
		Name       string
		StructName string
		Resource   Resource
		Version    string
	}{
		Name:       name,
		StructName: structName(name),
		Resource:   resource,
		Version:    spec.ResourceSpecificationVersion,
	}

	// Execute the template, writing it to file
	err = tmpl.Execute(f, templateData)
	if err != nil {
		fmt.Printf("Error: Failed to generate resource %s\n%s\n", fn, err)
		os.Exit(1)
	}

}
