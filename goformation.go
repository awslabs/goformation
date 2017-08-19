package goformation

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/awslabs/goformation/cloudformation"
	"github.com/awslabs/goformation/intrinsics"

	yamlwrapper "github.com/ghodss/yaml"

	"reflect"
	"log"
	"gopkg.in/yaml.v2"
)

//go:generate generate/generate.sh

type tagUnmarshalerType struct {
}

func (t *tagUnmarshalerType) UnmarshalYAMLTag(tag string, fieldValue reflect.Value) reflect.Value {

	log.Printf("%s %s", tag, fieldValue.Interface())
	output := reflect.ValueOf(make(map[string]interface{}))
	key := reflect.ValueOf(tag)

	output.SetMapIndex(key, fieldValue)

	return output
}
var tagUnmarshaller = &tagUnmarshalerType{}
var allTags = []string{"Ref", "GetAtt", "Base64", "FindInMap", "GetAZs",
	"ImportValue", "Join", "Select", "Split", "Sub",
}

func registerTagMarshallers() {
	for _, tag := range allTags {
		yaml.RegisterTagUnmarshaler("!"+tag, tagUnmarshaller)
	}
}

func unregisterTagMarshallers() {
	for _, tag := range allTags {
		yaml.RegisterTagUnmarshaler("!"+tag, tagUnmarshaller)
	}
}

// Open and parse a AWS CloudFormation template from file.
// Works with either JSON or YAML formatted templates.
func Open(filename string) (*cloudformation.Template, error) {

	registerTagMarshallers()

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	if strings.HasSuffix(filename, ".yaml") || strings.HasSuffix(filename, ".yml") {
		data, err = yamlwrapper.YAMLToJSON(data)
		if err != nil {
			return nil, fmt.Errorf("invalid YAML template: %s", err)
		}

	}

	return Parse(data)

}

// Parse an AWS CloudFormation template (expects a []byte of valid JSON)
func Parse(data []byte) (*cloudformation.Template, error) {

	// Process all AWS CloudFormation intrinsic functions (e.g. Fn::Join)
	intrinsified, err := intrinsics.Process(data, nil)
	if err != nil {
		return nil, err
	}

	template := &cloudformation.Template{}
	if err := json.Unmarshal(intrinsified, template); err != nil {
		return nil, err
	}

	return template, nil

}
