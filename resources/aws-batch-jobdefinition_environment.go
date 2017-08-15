package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Batch::JobDefinition.Environment AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-environment.html
type AWSBatchJobDefinition_Environment struct {

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-environment.html#cfn-batch-jobdefinition-environment-name
	Name string `json:"Name"`

	// Value AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-environment.html#cfn-batch-jobdefinition-environment-value
	Value string `json:"Value"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSBatchJobDefinition_Environment) AWSCloudFormationType() string {
	return "AWS::Batch::JobDefinition.Environment"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSBatchJobDefinition_Environment) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSBatchJobDefinition_EnvironmentResources retrieves all AWSBatchJobDefinition_Environment items from a CloudFormation template
func GetAllAWSBatchJobDefinition_Environment(template *Template) map[string]*AWSBatchJobDefinition_Environment {

	results := map[string]*AWSBatchJobDefinition_Environment{}
	for name, resource := range template.Resources {
		result := &AWSBatchJobDefinition_Environment{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSBatchJobDefinition_EnvironmentWithName retrieves all AWSBatchJobDefinition_Environment items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSBatchJobDefinition_Environment(name string, template *Template) (*AWSBatchJobDefinition_Environment, error) {

	result := &AWSBatchJobDefinition_Environment{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSBatchJobDefinition_Environment{}, errors.New("resource not found")

}
