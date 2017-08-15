package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Batch::JobDefinition.Ulimit AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ulimit.html
type AWSBatchJobDefinition_Ulimit struct {

	// HardLimit AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ulimit.html#cfn-batch-jobdefinition-ulimit-hardlimit

	HardLimit int64 `json:"HardLimit"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ulimit.html#cfn-batch-jobdefinition-ulimit-name

	Name string `json:"Name"`

	// SoftLimit AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-ulimit.html#cfn-batch-jobdefinition-ulimit-softlimit

	SoftLimit int64 `json:"SoftLimit"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSBatchJobDefinition_Ulimit) AWSCloudFormationType() string {
	return "AWS::Batch::JobDefinition.Ulimit"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSBatchJobDefinition_Ulimit) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSBatchJobDefinition_UlimitResources retrieves all AWSBatchJobDefinition_Ulimit items from a CloudFormation template
func GetAllAWSBatchJobDefinition_Ulimit(template *Template) map[string]*AWSBatchJobDefinition_Ulimit {

	results := map[string]*AWSBatchJobDefinition_Ulimit{}
	for name, resource := range template.Resources {
		result := &AWSBatchJobDefinition_Ulimit{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSBatchJobDefinition_UlimitWithName retrieves all AWSBatchJobDefinition_Ulimit items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSBatchJobDefinition_Ulimit(name string, template *Template) (*AWSBatchJobDefinition_Ulimit, error) {

	result := &AWSBatchJobDefinition_Ulimit{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSBatchJobDefinition_Ulimit{}, errors.New("resource not found")

}
