package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::CodePipeline::CustomActionType AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype.html
type AWSCodePipelineCustomActionType struct {

	// Category AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype.html#cfn-codepipeline-customactiontype-category
	Category string `json:"Category"`

	// ConfigurationProperties AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype.html#cfn-codepipeline-customactiontype-configurationproperties
	ConfigurationProperties []AWSCodePipelineCustomActionType_ConfigurationProperties `json:"ConfigurationProperties"`

	// InputArtifactDetails AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype.html#cfn-codepipeline-customactiontype-inputartifactdetails
	InputArtifactDetails AWSCodePipelineCustomActionType_ArtifactDetails `json:"InputArtifactDetails"`

	// OutputArtifactDetails AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype.html#cfn-codepipeline-customactiontype-outputartifactdetails
	OutputArtifactDetails AWSCodePipelineCustomActionType_ArtifactDetails `json:"OutputArtifactDetails"`

	// Provider AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype.html#cfn-codepipeline-customactiontype-provider
	Provider string `json:"Provider"`

	// Settings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype.html#cfn-codepipeline-customactiontype-settings
	Settings AWSCodePipelineCustomActionType_Settings `json:"Settings"`

	// Version AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-customactiontype.html#cfn-codepipeline-customactiontype-version
	Version string `json:"Version"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodePipelineCustomActionType) AWSCloudFormationType() string {
	return "AWS::CodePipeline::CustomActionType"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodePipelineCustomActionType) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCodePipelineCustomActionTypeResources retrieves all AWSCodePipelineCustomActionType items from a CloudFormation template
func GetAllAWSCodePipelineCustomActionType(template *Template) map[string]*AWSCodePipelineCustomActionType {

	results := map[string]*AWSCodePipelineCustomActionType{}
	for name, resource := range template.Resources {
		result := &AWSCodePipelineCustomActionType{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCodePipelineCustomActionTypeWithName retrieves all AWSCodePipelineCustomActionType items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCodePipelineCustomActionType(name string, template *Template) (*AWSCodePipelineCustomActionType, error) {

	result := &AWSCodePipelineCustomActionType{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCodePipelineCustomActionType{}, errors.New("resource not found")

}
