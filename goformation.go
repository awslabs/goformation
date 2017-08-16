package goformation

import (
	"encoding/json"
	"io/ioutil"

	"github.com/paulmaddox/goformation/intrinsics"
	"github.com/paulmaddox/goformation/resources"
)

//go:generate generate/generate.sh

// Open and parse a AWS CloudFormation template from file
func Open(filename string) (*resources.Template, error) {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return Parse(data)

}

// Parse an AWS CloudFormation template
func Parse(data []byte) (*resources.Template, error) {

	// Process all AWS CloudFormation intrinsic functions (e.g. Fn::Join)
	intrinsified, err := intrinsics.Process(data, nil)
	if err != nil {
		return nil, err
	}

	template := &resources.Template{}
	if err := json.Unmarshal(intrinsified, template); err != nil {
		return nil, err
	}

	return template, nil

}
