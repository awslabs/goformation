package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSCloudFormationStack AWS CloudFormation Resource (AWS::CloudFormation::Stack)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stack.html
type AWSCloudFormationStack struct {

	// NotificationARNs AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stack.html#cfn-cloudformation-stack-notificationarns
	NotificationARNs []string `json:"NotificationARNs"`

	// Parameters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stack.html#cfn-cloudformation-stack-parameters
	Parameters map[string]string `json:"Parameters"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stack.html#cfn-cloudformation-stack-tags
	Tags []Tag `json:"Tags"`

	// TemplateURL AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stack.html#cfn-cloudformation-stack-templateurl
	TemplateURL string `json:"TemplateURL"`

	// TimeoutInMinutes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-stack.html#cfn-cloudformation-stack-timeoutinminutes
	TimeoutInMinutes int64 `json:"TimeoutInMinutes"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCloudFormationStack) AWSCloudFormationType() string {
	return "AWS::CloudFormation::Stack"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCloudFormationStack) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCloudFormationStackResources retrieves all AWSCloudFormationStack items from a CloudFormation template
func GetAllAWSCloudFormationStackResources(template *Template) map[string]*AWSCloudFormationStack {

	results := map[string]*AWSCloudFormationStack{}
	for name, resource := range template.Resources {
		result := &AWSCloudFormationStack{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCloudFormationStackWithName retrieves all AWSCloudFormationStack items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSCloudFormationStackWithName(name string, template *Template) (*AWSCloudFormationStack, error) {

	result := &AWSCloudFormationStack{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCloudFormationStack{}, errors.New("resource not found")

}
