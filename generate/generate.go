package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"text/template"
)

// ResourceGenerator takes AWS CloudFormation Resource Specification
// documents, and generates Go structs and a JSON Schema from them.
type ResourceGenerator struct {
	filenames map[string]string
	urls      map[string]string
	Results   *ResourceGeneratorResults
}

// ResourceGeneratorResults contains a summary of the items generated
type ResourceGeneratorResults struct {
	UpdatedResources map[string]string
	UpdatedSchemas   map[string]string
	ProcessedCount   int
}

// NewResourceGenerator an array of AWS CloudFormation Resource Specification
// documents, and generates Go structs and a JSON Schema from them.
// The input can be a mix of URLs (https://) or files (file://).
func NewResourceGenerator(urls map[string]string) (*ResourceGenerator, error) {

	rg := &ResourceGenerator{
		filenames: map[string]string{},
		urls:      map[string]string{},
		Results: &ResourceGeneratorResults{
			UpdatedResources: map[string]string{},
			UpdatedSchemas:   map[string]string{},
			ProcessedCount:   0,
		},
	}

	for name, found := range urls {

		uri, err := url.Parse(found)
		if err != nil {
			return nil, err
		}

		switch uri.Scheme {
		case "https":
			rg.urls[name] = uri.String()
		case "http":
			rg.urls[name] = uri.String()
		case "file":
			rg.filenames[name] = uri.String()
		default:
			return nil, fmt.Errorf("invalid URL scheme %s", uri.Scheme)
		}

	}

	return rg, nil

}

// Generate generates Go structs and a JSON Schema from the AWS CloudFormation
// Resource Specifications provided to NewResourceGenerator
func (rg *ResourceGenerator) Generate() error {

	for name, uri := range rg.urls {
		fmt.Printf("Downloading %s specification from %s\n", name, uri)
		data, err := rg.downloadSpec(uri)
		if err != nil {
			return err
		}
		spec, err := rg.processSpec(name, data)
		if err != nil {
			return err
		}
		if err := rg.generateJSONSchema(name, spec); err != nil {
			return err
		}
	}

	for name, filename := range rg.filenames {
		fmt.Printf("Downloading %s specification from %s\n", name, filename)
		data, err := ioutil.ReadFile(strings.Replace(filename, "file://", "", -1))
		if err != nil {
			return err
		}
		spec, err := rg.processSpec(name, data)
		if err != nil {
			return err
		}
		if err := rg.generateJSONSchema(name, spec); err != nil {
			return err
		}

	}

	return nil

}

func (rg *ResourceGenerator) downloadSpec(uri string) ([]byte, error) {

	response, err := http.Get(uri)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return data, nil

}

func (rg *ResourceGenerator) processSpec(specname string, data []byte) (*CloudFormationResourceSpecification, error) {

	// Unmarshall the JSON specification
	spec := &CloudFormationResourceSpecification{}
	if err := json.Unmarshal(data, spec); err != nil {
		return nil, err
	}

	// Write all of the resources in the spec file
	for name, resource := range spec.Resources {
		if err := rg.generateResources(name, resource, false, spec); err != nil {
			return nil, err
		}
	}

	// Write all of the custom types in the spec file
	for name, resource := range spec.Properties {
		if err := rg.generateResources(name, resource, true, spec); err != nil {
			return nil, err
		}
	}

	return spec, nil

}

func (rg *ResourceGenerator) generateResources(name string, resource Resource, isCustomProperty bool, spec *CloudFormationResourceSpecification) error {

	// Open the resource template
	tmpl, err := template.ParseFiles("generate/templates/resource.template")
	if err != nil {
		return fmt.Errorf("failed to load resource template: %s", err)
	}

	// Pass in the following information into the template
	sname := structName(name)
	structNameParts := strings.Split(name, ".")
	basename := structName(structNameParts[0])

	templateData := struct {
		Name             string
		StructName       string
		Basename         string
		Resource         Resource
		IsCustomProperty bool
		Version          string
	}{
		Name:             name,
		StructName:       sname,
		Basename:         basename,
		Resource:         resource,
		IsCustomProperty: isCustomProperty,
		Version:          spec.ResourceSpecificationVersion,
	}

	// Execute the template, writing it to a buffer
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, templateData)
	if err != nil {
		return fmt.Errorf("failed to generate resource %s: %s", name, err)
	}

	// Format the generated Go code with gofmt
	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		return fmt.Errorf("failed to format Go file for resource %s: %s", name, err)
	}

	// Check if the file has changed since the last time generate ran
	fn := "cloudformation/" + filename(name)
	current, err := ioutil.ReadFile(fn)

	if err != nil || bytes.Compare(formatted, current) != 0 {

		// Write the file contents out
		if err := ioutil.WriteFile(fn, formatted, 0644); err != nil {
			return fmt.Errorf("failed to write resource file %s: %s", fn, err)
		}

		// Log the updated class name to the results
		rg.Results.UpdatedResources[fn] = templateData.Basename + "_" + templateData.StructName

	}

	rg.Results.ProcessedCount++

	return nil

}

func (rg *ResourceGenerator) generateJSONSchema(specname string, spec *CloudFormationResourceSpecification) error {

	// Open the schema template and setup a counter function that will
	// available in the template to be used to detect when trailing commas
	// are required in the JSON when looping through maps
	tmpl, err := template.New("schema.template").Funcs(template.FuncMap{
		"counter": counter,
	}).ParseFiles("generate/templates/schema.template")

	var buf bytes.Buffer

	// Execute the template, writing it to file
	err = tmpl.Execute(&buf, spec)
	if err != nil {
		return fmt.Errorf("failed to generate JSON Schema: %s", err)
	}

	// Parse it to JSON objects and back again to format it
	var j interface{}
	if err := json.Unmarshal(buf.Bytes(), &j); err != nil {
		return fmt.Errorf("failed to unmarshal JSON Schema: %s", err)
	}

	formatted, err := json.MarshalIndent(j, "", "    ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON Schema: %s", err)

	}

	filename := fmt.Sprintf("schema/%s.schema.json", specname)
	if err := ioutil.WriteFile(filename, formatted, 0644); err != nil {
		return fmt.Errorf("failed to write JSON Schema: %s", err)
	}

	rg.Results.UpdatedSchemas[filename] = specname

	return nil

}

func generatePolymorphicProperty(name string, property Property) {

	// Open the polymorphic property template
	tmpl, err := template.New("polymorphic-property.template").Funcs(template.FuncMap{
		"convertToGoType": convertTypeToGo,
	}).ParseFiles("generate/templates/polymorphic-property.template")

	nameParts := strings.Split(name, "_")

	types := append([]string{}, property.PrimitiveTypes...)
	types = append(types, property.PrimitiveItemTypes...)
	types = append(types, property.ItemTypes...)
	types = append(types, property.Types...)

	templateData := struct {
		Name        string
		Basename    string
		Property    Property
		Types       []string
		TypesJoined string
	}{
		Name:        name,
		Basename:    nameParts[0],
		Property:    property,
		Types:       types,
		TypesJoined: conjoin("or", types),
	}

	// Execute the template, writing it to file
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, templateData)
	if err != nil {
		fmt.Printf("Error: Failed to generate polymorphic property %s\n%s\n", name, err)
		os.Exit(1)
	}

	// Format the generated Go file with gofmt
	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Printf("Error: Failed to format Go file for resource %s\n%s\n", name, err)
		os.Exit(1)
	}

	// Write the file out
	if err := ioutil.WriteFile("cloudformation/"+filename(name), formatted, 0644); err != nil {
		fmt.Printf("Error: Failed to write JSON Schema\n%s\n", err)
		os.Exit(1)
	}

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

func conjoin(conj string, items []string) string {
	if len(items) == 0 {
		return ""
	}
	if len(items) == 1 {
		return items[0]
	}
	if len(items) == 2 { // "a and b" not "a, and b"
		return items[0] + " " + conj + " " + items[1]
	}

	sep := ", "
	pieces := []string{items[0]}
	for _, item := range items[1 : len(items)-1] {
		pieces = append(pieces, sep, item)
	}
	pieces = append(pieces, sep, conj, " ", items[len(items)-1])

	return strings.Join(pieces, "")
}
