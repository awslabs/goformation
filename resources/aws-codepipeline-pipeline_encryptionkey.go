package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::CodePipeline::Pipeline.EncryptionKey AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-artifactstore-encryptionkey.html
type AWSCodePipelinePipeline_EncryptionKey struct {

	// Id AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-artifactstore-encryptionkey.html#cfn-codepipeline-pipeline-artifactstore-encryptionkey-id
	Id string `json:"Id"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-artifactstore-encryptionkey.html#cfn-codepipeline-pipeline-artifactstore-encryptionkey-type
	Type string `json:"Type"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodePipelinePipeline_EncryptionKey) AWSCloudFormationType() string {
	return "AWS::CodePipeline::Pipeline.EncryptionKey"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodePipelinePipeline_EncryptionKey) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCodePipelinePipeline_EncryptionKeyResources retrieves all AWSCodePipelinePipeline_EncryptionKey items from a CloudFormation template
func GetAllAWSCodePipelinePipeline_EncryptionKey(template *Template) map[string]*AWSCodePipelinePipeline_EncryptionKey {

	results := map[string]*AWSCodePipelinePipeline_EncryptionKey{}
	for name, resource := range template.Resources {
		result := &AWSCodePipelinePipeline_EncryptionKey{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCodePipelinePipeline_EncryptionKeyWithName retrieves all AWSCodePipelinePipeline_EncryptionKey items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCodePipelinePipeline_EncryptionKey(name string, template *Template) (*AWSCodePipelinePipeline_EncryptionKey, error) {

	result := &AWSCodePipelinePipeline_EncryptionKey{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCodePipelinePipeline_EncryptionKey{}, errors.New("resource not found")

}
