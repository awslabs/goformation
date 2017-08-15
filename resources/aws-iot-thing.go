package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::IoT::Thing AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thing.html
type AWSIoTThing struct {

	// AttributePayload AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thing.html#cfn-iot-thing-attributepayload

	AttributePayload AWSIoTThing_AttributePayload `json:"AttributePayload"`

	// ThingName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thing.html#cfn-iot-thing-thingname

	ThingName string `json:"ThingName"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTThing) AWSCloudFormationType() string {
	return "AWS::IoT::Thing"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIoTThing) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSIoTThingResources retrieves all AWSIoTThing items from a CloudFormation template
func GetAllAWSIoTThing(template *Template) map[string]*AWSIoTThing {

	results := map[string]*AWSIoTThing{}
	for name, resource := range template.Resources {
		result := &AWSIoTThing{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSIoTThingWithName retrieves all AWSIoTThing items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSIoTThing(name string, template *Template) (*AWSIoTThing, error) {

	result := &AWSIoTThing{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSIoTThing{}, errors.New("resource not found")

}
