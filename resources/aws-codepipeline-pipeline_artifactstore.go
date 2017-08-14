package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::CodePipeline::Pipeline.ArtifactStore AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-artifactstore.html
type AWSCodePipelinePipeline_ArtifactStore struct {

	// EncryptionKey AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-artifactstore.html#cfn-codepipeline-pipeline-artifactstore-encryptionkey

	EncryptionKey AWSCodePipelinePipeline_EncryptionKey `json:"EncryptionKey"`

	// Location AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-artifactstore.html#cfn-codepipeline-pipeline-artifactstore-location

	Location string `json:"Location"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-artifactstore.html#cfn-codepipeline-pipeline-artifactstore-type

	Type string `json:"Type"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodePipelinePipeline_ArtifactStore) AWSCloudFormationType() string {
	return "AWS::CodePipeline::Pipeline.ArtifactStore"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodePipelinePipeline_ArtifactStore) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCodePipelinePipeline_ArtifactStoreResources retrieves all AWSCodePipelinePipeline_ArtifactStore items from a CloudFormation template
func GetAllAWSCodePipelinePipeline_ArtifactStore(template *Template) map[string]*AWSCodePipelinePipeline_ArtifactStore {

	results := map[string]*AWSCodePipelinePipeline_ArtifactStore{}
	for name, resource := range template.Resources {
		result := &AWSCodePipelinePipeline_ArtifactStore{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCodePipelinePipeline_ArtifactStoreWithName retrieves all AWSCodePipelinePipeline_ArtifactStore items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCodePipelinePipeline_ArtifactStore(name string, template *Template) (*AWSCodePipelinePipeline_ArtifactStore, error) {

	result := &AWSCodePipelinePipeline_ArtifactStore{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCodePipelinePipeline_ArtifactStore{}, errors.New("resource not found")

}
