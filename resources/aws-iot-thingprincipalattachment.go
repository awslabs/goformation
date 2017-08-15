package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::IoT::ThingPrincipalAttachment AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thingprincipalattachment.html
type AWSIoTThingPrincipalAttachment struct {

	// Principal AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thingprincipalattachment.html#cfn-iot-thingprincipalattachment-principal
	Principal string `json:"Principal"`

	// ThingName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thingprincipalattachment.html#cfn-iot-thingprincipalattachment-thingname
	ThingName string `json:"ThingName"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTThingPrincipalAttachment) AWSCloudFormationType() string {
	return "AWS::IoT::ThingPrincipalAttachment"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIoTThingPrincipalAttachment) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSIoTThingPrincipalAttachmentResources retrieves all AWSIoTThingPrincipalAttachment items from a CloudFormation template
func GetAllAWSIoTThingPrincipalAttachment(template *Template) map[string]*AWSIoTThingPrincipalAttachment {

	results := map[string]*AWSIoTThingPrincipalAttachment{}
	for name, resource := range template.Resources {
		result := &AWSIoTThingPrincipalAttachment{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSIoTThingPrincipalAttachmentWithName retrieves all AWSIoTThingPrincipalAttachment items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSIoTThingPrincipalAttachment(name string, template *Template) (*AWSIoTThingPrincipalAttachment, error) {

	result := &AWSIoTThingPrincipalAttachment{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSIoTThingPrincipalAttachment{}, errors.New("resource not found")

}
