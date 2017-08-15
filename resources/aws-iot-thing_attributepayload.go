package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::IoT::Thing.AttributePayload AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-thing-attributepayload.html
type AWSIoTThing_AttributePayload struct {

	// Attributes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-thing-attributepayload.html#cfn-iot-thing-attributepayload-attributes
	Attributes map[string]string `json:"Attributes"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTThing_AttributePayload) AWSCloudFormationType() string {
	return "AWS::IoT::Thing.AttributePayload"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIoTThing_AttributePayload) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSIoTThing_AttributePayloadResources retrieves all AWSIoTThing_AttributePayload items from a CloudFormation template
func GetAllAWSIoTThing_AttributePayload(template *Template) map[string]*AWSIoTThing_AttributePayload {

	results := map[string]*AWSIoTThing_AttributePayload{}
	for name, resource := range template.Resources {
		result := &AWSIoTThing_AttributePayload{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSIoTThing_AttributePayloadWithName retrieves all AWSIoTThing_AttributePayload items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSIoTThing_AttributePayload(name string, template *Template) (*AWSIoTThing_AttributePayload, error) {

	result := &AWSIoTThing_AttributePayload{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSIoTThing_AttributePayload{}, errors.New("resource not found")

}
