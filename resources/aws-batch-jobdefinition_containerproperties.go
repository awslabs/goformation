package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Batch::JobDefinition.ContainerProperties AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html
type AWSBatchJobDefinition_ContainerProperties struct {

	// Command AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-command
	Command []string `json:"Command"`

	// Environment AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-environment
	Environment []AWSBatchJobDefinition_Environment `json:"Environment"`

	// Image AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-image
	Image string `json:"Image"`

	// JobRoleArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-jobrolearn
	JobRoleArn string `json:"JobRoleArn"`

	// Memory AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-memory
	Memory int64 `json:"Memory"`

	// MountPoints AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-mountpoints
	MountPoints []AWSBatchJobDefinition_MountPoints `json:"MountPoints"`

	// Privileged AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-privileged
	Privileged bool `json:"Privileged"`

	// ReadonlyRootFilesystem AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-readonlyrootfilesystem
	ReadonlyRootFilesystem bool `json:"ReadonlyRootFilesystem"`

	// Ulimits AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-ulimits
	Ulimits []AWSBatchJobDefinition_Ulimit `json:"Ulimits"`

	// User AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-user
	User string `json:"User"`

	// Vcpus AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-vcpus
	Vcpus int64 `json:"Vcpus"`

	// Volumes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-containerproperties.html#cfn-batch-jobdefinition-containerproperties-volumes
	Volumes []AWSBatchJobDefinition_Volumes `json:"Volumes"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSBatchJobDefinition_ContainerProperties) AWSCloudFormationType() string {
	return "AWS::Batch::JobDefinition.ContainerProperties"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSBatchJobDefinition_ContainerProperties) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSBatchJobDefinition_ContainerPropertiesResources retrieves all AWSBatchJobDefinition_ContainerProperties items from a CloudFormation template
func GetAllAWSBatchJobDefinition_ContainerProperties(template *Template) map[string]*AWSBatchJobDefinition_ContainerProperties {

	results := map[string]*AWSBatchJobDefinition_ContainerProperties{}
	for name, resource := range template.Resources {
		result := &AWSBatchJobDefinition_ContainerProperties{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSBatchJobDefinition_ContainerPropertiesWithName retrieves all AWSBatchJobDefinition_ContainerProperties items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSBatchJobDefinition_ContainerProperties(name string, template *Template) (*AWSBatchJobDefinition_ContainerProperties, error) {

	result := &AWSBatchJobDefinition_ContainerProperties{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSBatchJobDefinition_ContainerProperties{}, errors.New("resource not found")

}
