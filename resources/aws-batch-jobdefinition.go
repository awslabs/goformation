package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSBatchJobDefinition AWS CloudFormation Resource (AWS::Batch::JobDefinition)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobdefinition.html
type AWSBatchJobDefinition struct {

	// ContainerProperties AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobdefinition.html#cfn-batch-jobdefinition-containerproperties
	ContainerProperties AWSBatchJobDefinition_ContainerProperties `json:"ContainerProperties"`

	// JobDefinitionName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobdefinition.html#cfn-batch-jobdefinition-jobdefinitionname
	JobDefinitionName string `json:"JobDefinitionName"`

	// Parameters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobdefinition.html#cfn-batch-jobdefinition-parameters
	Parameters interface{} `json:"Parameters"`

	// RetryStrategy AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobdefinition.html#cfn-batch-jobdefinition-retrystrategy
	RetryStrategy AWSBatchJobDefinition_RetryStrategy `json:"RetryStrategy"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-batch-jobdefinition.html#cfn-batch-jobdefinition-type
	Type string `json:"Type"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSBatchJobDefinition) AWSCloudFormationType() string {
	return "AWS::Batch::JobDefinition"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSBatchJobDefinition) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSBatchJobDefinitionResources retrieves all AWSBatchJobDefinition items from a CloudFormation template
func GetAllAWSBatchJobDefinitionResources(template *Template) map[string]*AWSBatchJobDefinition {

	results := map[string]*AWSBatchJobDefinition{}
	for name, resource := range template.Resources {
		result := &AWSBatchJobDefinition{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSBatchJobDefinitionWithName retrieves all AWSBatchJobDefinition items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSBatchJobDefinitionWithName(name string, template *Template) (*AWSBatchJobDefinition, error) {

	result := &AWSBatchJobDefinition{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSBatchJobDefinition{}, errors.New("resource not found")

}
