package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSStepFunctionsActivity AWS CloudFormation Resource (AWS::StepFunctions::Activity)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-activity.html
type AWSStepFunctionsActivity struct {

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-stepfunctions-activity.html#cfn-stepfunctions-activity-name
	Name string `json:"Name"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSStepFunctionsActivity) AWSCloudFormationType() string {
	return "AWS::StepFunctions::Activity"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSStepFunctionsActivity) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSStepFunctionsActivityResources retrieves all AWSStepFunctionsActivity items from a CloudFormation template
func (t *Template) GetAllAWSStepFunctionsActivityResources() map[string]*AWSStepFunctionsActivity {

	results := map[string]*AWSStepFunctionsActivity{}
	for name, resource := range t.Resources {
		result := &AWSStepFunctionsActivity{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSStepFunctionsActivityWithName retrieves all AWSStepFunctionsActivity items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSStepFunctionsActivityWithName(name string) (*AWSStepFunctionsActivity, error) {

	result := &AWSStepFunctionsActivity{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSStepFunctionsActivity{}, errors.New("resource not found")

}
