package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"io/ioutil"
	"net/http"
	"os"
	"text/template"
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
		generateResources(name, resource, spec)
	}

	// Write all of the custom properties, using a template
	for name, property := range spec.Properties {
		fmt.Printf("Generating Custom Property: %s\n", name)
		generateResources(name, property, spec)
	}

	// Generate the JSON-Schema
	generateSchema(spec)

}

// generateResources generates Go structs for all of the resources and custom property types
// found in a CloudformationResourceSpecification
func generateResources(name string, resource Resource, spec *CloudFormationResourceSpecification) {

	// Open the resource template
	tmpl, err := template.ParseFiles("resource.template")
	if err != nil {
		fmt.Printf("Error: Failed to load resource template\n%s\n", err)
		os.Exit(1)
	}

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
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, templateData)
	if err != nil {
		fmt.Printf("Error: Failed to generate resource %s\n%s\n", name, err)
		os.Exit(1)
	}

	// Format the generated Go file with gofmt
	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Printf("Error: Failed to format Go file for resource %s\n%s\n", name, err)
		os.Exit(1)
	}

	// Write the file out
	if err := ioutil.WriteFile("resources/"+filename(name), formatted, 0644); err != nil {
		fmt.Printf("Error: Failed to write JSON Schema\n%s\n", err)
		os.Exit(1)
	}

}

// generateResources generates a JSON Schema for all of the resources and custom property types
// found in a CloudformationResourceSpecification
func generateSchema(spec *CloudFormationResourceSpecification) {

	// Open the schema template and setup a counter function that will
	// available in the template to be used to detect when trailing commas
	// are required in the JSON when looping through maps
	tmpl, err := template.New("schema.template").Funcs(template.FuncMap{
		"counter": counter,
	}).ParseFiles("schema.template")

	var buf bytes.Buffer

	// Execute the template, writing it to file
	err = tmpl.Execute(&buf, spec)
	if err != nil {
		fmt.Printf("Error: Failed to generate JSON Schema\n%s\n", err)
		os.Exit(1)
	}

	// Parse it to JSON objects and back again to format it
	var j interface{}
	if err := json.Unmarshal(buf.Bytes(), &j); err != nil {
		fmt.Printf("Error: Failed to unmarhal JSON Schema\n%s\n", err)
		os.Exit(1)
	}

	formatted, err := json.MarshalIndent(j, "", "    ")
	if err != nil {
		fmt.Printf("Error: Failed to marshal JSON Schema\n%s\n", err)
		os.Exit(1)
	}

	if err := ioutil.WriteFile("resources/json.schema", formatted, 0644); err != nil {
		fmt.Printf("Error: Failed to write JSON Schema\n%s\n", err)
		os.Exit(1)
	}

	fmt.Printf("Generated JSON Schema: resources/json.schema\n")

}

// counter is used within the JSON Schema template to determin whether or not
// to put a comma after a JSON resource (i.e. if it's the last element, then no comma)
// see: http://android.wekeepcoding.com/article/10126058/Go+template+remove+the+last+comma+in+range+loop
func counter(length int) func() int {
	i := length
	return func() int {
		i--
		return i
	}
}
