package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Batch::JobDefinition.RetryStrategy AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-retrystrategy.html
type AWSBatchJobDefinition_RetryStrategy struct {

	// Attempts AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-retrystrategy.html#cfn-batch-jobdefinition-retrystrategy-attempts

	Attempts int64 `json:"Attempts"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSBatchJobDefinition_RetryStrategy) AWSCloudFormationType() string {
	return "AWS::Batch::JobDefinition.RetryStrategy"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSBatchJobDefinition_RetryStrategy) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSBatchJobDefinition_RetryStrategyResources retrieves all AWSBatchJobDefinition_RetryStrategy items from a CloudFormation template
func GetAllAWSBatchJobDefinition_RetryStrategy(template *Template) map[string]*AWSBatchJobDefinition_RetryStrategy {

	results := map[string]*AWSBatchJobDefinition_RetryStrategy{}
	for name, resource := range template.Resources {
		result := &AWSBatchJobDefinition_RetryStrategy{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSBatchJobDefinition_RetryStrategyWithName retrieves all AWSBatchJobDefinition_RetryStrategy items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSBatchJobDefinition_RetryStrategy(name string, template *Template) (*AWSBatchJobDefinition_RetryStrategy, error) {

	result := &AWSBatchJobDefinition_RetryStrategy{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSBatchJobDefinition_RetryStrategy{}, errors.New("resource not found")

}
