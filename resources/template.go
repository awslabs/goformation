package resources

import (
	"encoding/json"

	yaml "gopkg.in/yaml.v2"
)

// CloudFormationTemplate represents an AWS CloudFormation template
// see: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/template-anatomy.html
type CloudFormationTemplate struct {
	AWSTemplateFormatVersion string                 `json:"AWSTemplateFormatVersion,omitempty"`
	Description              string                 `json:"Description,omitempty"`
	Metadata                 map[string]interface{} `json:"Metadata,omitempty"`
	Parameters               map[string]interface{} `json:"Parameters,omitempty"`
	Mappings                 map[string]interface{} `json:"Mappings,omitempty"`
	Conditions               map[string]interface{} `json:"Conditions,omitempty"`
	Resources                map[string]interface{} `json:"Resources,omitempty"`
	Outputs                  map[string]interface{} `json:"Outputs,omitempty"`
}

// NewCloudFormationTemplate creates a new AWS CloudFormation template struct
func NewCloudFormationTemplate() *CloudFormationTemplate {
	return &CloudFormationTemplate{
		AWSTemplateFormatVersion: "2010-09-09",
		Description:              "",
		Metadata:                 map[string]interface{}{},
		Parameters:               map[string]interface{}{},
		Mappings:                 map[string]interface{}{},
		Conditions:               map[string]interface{}{},
		Resources:                map[string]interface{}{},
		Outputs:                  map[string]interface{}{},
	}
}

// JSON converts an AWS CloudFormation template object to JSON
func (t *CloudFormationTemplate) JSON() ([]byte, error) {
	return json.Marshal(t)
}

// YAML converts an AWS CloudFormation template object to YAML
func (t *CloudFormationTemplate) YAML() ([]byte, error) {
	return yaml.Marshal(t)
}
