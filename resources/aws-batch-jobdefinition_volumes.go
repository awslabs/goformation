package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Batch::JobDefinition.Volumes AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-volumes.html
type AWSBatchJobDefinition_Volumes struct {

	// Host AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-volumes.html#cfn-batch-jobdefinition-volumes-host
	Host AWSBatchJobDefinition_VolumesHost `json:"Host"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-volumes.html#cfn-batch-jobdefinition-volumes-name
	Name string `json:"Name"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSBatchJobDefinition_Volumes) AWSCloudFormationType() string {
	return "AWS::Batch::JobDefinition.Volumes"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSBatchJobDefinition_Volumes) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSBatchJobDefinition_VolumesResources retrieves all AWSBatchJobDefinition_Volumes items from a CloudFormation template
func GetAllAWSBatchJobDefinition_Volumes(template *Template) map[string]*AWSBatchJobDefinition_Volumes {

	results := map[string]*AWSBatchJobDefinition_Volumes{}
	for name, resource := range template.Resources {
		result := &AWSBatchJobDefinition_Volumes{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSBatchJobDefinition_VolumesWithName retrieves all AWSBatchJobDefinition_Volumes items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSBatchJobDefinition_Volumes(name string, template *Template) (*AWSBatchJobDefinition_Volumes, error) {

	result := &AWSBatchJobDefinition_Volumes{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSBatchJobDefinition_Volumes{}, errors.New("resource not found")

}
