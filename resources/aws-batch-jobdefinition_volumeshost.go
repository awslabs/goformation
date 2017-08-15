package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Batch::JobDefinition.VolumesHost AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-volumeshost.html
type AWSBatchJobDefinition_VolumesHost struct {

	// SourcePath AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-volumeshost.html#cfn-batch-jobdefinition-volumeshost-sourcepath
	SourcePath string `json:"SourcePath"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSBatchJobDefinition_VolumesHost) AWSCloudFormationType() string {
	return "AWS::Batch::JobDefinition.VolumesHost"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSBatchJobDefinition_VolumesHost) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSBatchJobDefinition_VolumesHostResources retrieves all AWSBatchJobDefinition_VolumesHost items from a CloudFormation template
func GetAllAWSBatchJobDefinition_VolumesHost(template *Template) map[string]*AWSBatchJobDefinition_VolumesHost {

	results := map[string]*AWSBatchJobDefinition_VolumesHost{}
	for name, resource := range template.Resources {
		result := &AWSBatchJobDefinition_VolumesHost{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSBatchJobDefinition_VolumesHostWithName retrieves all AWSBatchJobDefinition_VolumesHost items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSBatchJobDefinition_VolumesHost(name string, template *Template) (*AWSBatchJobDefinition_VolumesHost, error) {

	result := &AWSBatchJobDefinition_VolumesHost{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSBatchJobDefinition_VolumesHost{}, errors.New("resource not found")

}
