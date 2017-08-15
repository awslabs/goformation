package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::CodePipeline::CustomActionType.ArtifactDetails AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-artifactdetails.html
type AWSCodePipelineCustomActionType_ArtifactDetails struct {

	// MaximumCount AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-artifactdetails.html#cfn-codepipeline-customactiontype-artifactdetails-maximumcount
	MaximumCount int64 `json:"MaximumCount"`

	// MinimumCount AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-artifactdetails.html#cfn-codepipeline-customactiontype-artifactdetails-minimumcount
	MinimumCount int64 `json:"MinimumCount"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodePipelineCustomActionType_ArtifactDetails) AWSCloudFormationType() string {
	return "AWS::CodePipeline::CustomActionType.ArtifactDetails"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodePipelineCustomActionType_ArtifactDetails) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCodePipelineCustomActionType_ArtifactDetailsResources retrieves all AWSCodePipelineCustomActionType_ArtifactDetails items from a CloudFormation template
func GetAllAWSCodePipelineCustomActionType_ArtifactDetails(template *Template) map[string]*AWSCodePipelineCustomActionType_ArtifactDetails {

	results := map[string]*AWSCodePipelineCustomActionType_ArtifactDetails{}
	for name, resource := range template.Resources {
		result := &AWSCodePipelineCustomActionType_ArtifactDetails{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCodePipelineCustomActionType_ArtifactDetailsWithName retrieves all AWSCodePipelineCustomActionType_ArtifactDetails items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCodePipelineCustomActionType_ArtifactDetails(name string, template *Template) (*AWSCodePipelineCustomActionType_ArtifactDetails, error) {

	result := &AWSCodePipelineCustomActionType_ArtifactDetails{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCodePipelineCustomActionType_ArtifactDetails{}, errors.New("resource not found")

}
