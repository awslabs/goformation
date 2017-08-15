package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::CloudFormation::WaitConditionHandle AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waitconditionhandle.html
type AWSCloudFormationWaitConditionHandle struct {
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCloudFormationWaitConditionHandle) AWSCloudFormationType() string {
	return "AWS::CloudFormation::WaitConditionHandle"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCloudFormationWaitConditionHandle) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCloudFormationWaitConditionHandleResources retrieves all AWSCloudFormationWaitConditionHandle items from a CloudFormation template
func GetAllAWSCloudFormationWaitConditionHandle(template *Template) map[string]*AWSCloudFormationWaitConditionHandle {

	results := map[string]*AWSCloudFormationWaitConditionHandle{}
	for name, resource := range template.Resources {
		result := &AWSCloudFormationWaitConditionHandle{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCloudFormationWaitConditionHandleWithName retrieves all AWSCloudFormationWaitConditionHandle items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCloudFormationWaitConditionHandle(name string, template *Template) (*AWSCloudFormationWaitConditionHandle, error) {

	result := &AWSCloudFormationWaitConditionHandle{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCloudFormationWaitConditionHandle{}, errors.New("resource not found")

}
