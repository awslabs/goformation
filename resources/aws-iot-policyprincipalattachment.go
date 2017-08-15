package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::IoT::PolicyPrincipalAttachment AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-policyprincipalattachment.html
type AWSIoTPolicyPrincipalAttachment struct {

	// PolicyName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-policyprincipalattachment.html#cfn-iot-policyprincipalattachment-policyname
	PolicyName string `json:"PolicyName"`

	// Principal AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-policyprincipalattachment.html#cfn-iot-policyprincipalattachment-principal
	Principal string `json:"Principal"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTPolicyPrincipalAttachment) AWSCloudFormationType() string {
	return "AWS::IoT::PolicyPrincipalAttachment"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIoTPolicyPrincipalAttachment) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSIoTPolicyPrincipalAttachmentResources retrieves all AWSIoTPolicyPrincipalAttachment items from a CloudFormation template
func GetAllAWSIoTPolicyPrincipalAttachmentResources(template *Template) map[string]*AWSIoTPolicyPrincipalAttachment {

	results := map[string]*AWSIoTPolicyPrincipalAttachment{}
	for name, resource := range template.Resources {
		result := &AWSIoTPolicyPrincipalAttachment{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSIoTPolicyPrincipalAttachmentWithName retrieves all AWSIoTPolicyPrincipalAttachment items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSIoTPolicyPrincipalAttachmentWithName(name string, template *Template) (*AWSIoTPolicyPrincipalAttachment, error) {

	result := &AWSIoTPolicyPrincipalAttachment{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSIoTPolicyPrincipalAttachment{}, errors.New("resource not found")

}
