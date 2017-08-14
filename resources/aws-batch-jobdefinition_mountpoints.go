package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Batch::JobDefinition.MountPoints AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-mountpoints.html
type AWSBatchJobDefinition_MountPoints struct {

	// ContainerPath AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-mountpoints.html#cfn-batch-jobdefinition-mountpoints-containerpath

	ContainerPath string `json:"ContainerPath"`

	// ReadOnly AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-mountpoints.html#cfn-batch-jobdefinition-mountpoints-readonly

	ReadOnly bool `json:"ReadOnly"`

	// SourceVolume AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-mountpoints.html#cfn-batch-jobdefinition-mountpoints-sourcevolume

	SourceVolume string `json:"SourceVolume"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSBatchJobDefinition_MountPoints) AWSCloudFormationType() string {
	return "AWS::Batch::JobDefinition.MountPoints"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSBatchJobDefinition_MountPoints) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSBatchJobDefinition_MountPointsResources retrieves all AWSBatchJobDefinition_MountPoints items from a CloudFormation template
func GetAllAWSBatchJobDefinition_MountPoints(template *Template) map[string]*AWSBatchJobDefinition_MountPoints {

	results := map[string]*AWSBatchJobDefinition_MountPoints{}
	for name, resource := range template.Resources {
		result := &AWSBatchJobDefinition_MountPoints{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSBatchJobDefinition_MountPointsWithName retrieves all AWSBatchJobDefinition_MountPoints items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSBatchJobDefinition_MountPoints(name string, template *Template) (*AWSBatchJobDefinition_MountPoints, error) {

	result := &AWSBatchJobDefinition_MountPoints{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSBatchJobDefinition_MountPoints{}, errors.New("resource not found")

}
