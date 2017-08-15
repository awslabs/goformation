package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::CodePipeline::CustomActionType.Settings AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-settings.html
type AWSCodePipelineCustomActionType_Settings struct {

	// EntityUrlTemplate AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-settings.html#cfn-codepipeline-customactiontype-settings-entityurltemplate
	EntityUrlTemplate string `json:"EntityUrlTemplate"`

	// ExecutionUrlTemplate AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-settings.html#cfn-codepipeline-customactiontype-settings-executionurltemplate
	ExecutionUrlTemplate string `json:"ExecutionUrlTemplate"`

	// RevisionUrlTemplate AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-settings.html#cfn-codepipeline-customactiontype-settings-revisionurltemplate
	RevisionUrlTemplate string `json:"RevisionUrlTemplate"`

	// ThirdPartyConfigurationUrl AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-settings.html#cfn-codepipeline-customactiontype-settings-thirdpartyconfigurationurl
	ThirdPartyConfigurationUrl string `json:"ThirdPartyConfigurationUrl"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodePipelineCustomActionType_Settings) AWSCloudFormationType() string {
	return "AWS::CodePipeline::CustomActionType.Settings"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodePipelineCustomActionType_Settings) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCodePipelineCustomActionType_SettingsResources retrieves all AWSCodePipelineCustomActionType_Settings items from a CloudFormation template
func GetAllAWSCodePipelineCustomActionType_Settings(template *Template) map[string]*AWSCodePipelineCustomActionType_Settings {

	results := map[string]*AWSCodePipelineCustomActionType_Settings{}
	for name, resource := range template.Resources {
		result := &AWSCodePipelineCustomActionType_Settings{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCodePipelineCustomActionType_SettingsWithName retrieves all AWSCodePipelineCustomActionType_Settings items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCodePipelineCustomActionType_Settings(name string, template *Template) (*AWSCodePipelineCustomActionType_Settings, error) {

	result := &AWSCodePipelineCustomActionType_Settings{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCodePipelineCustomActionType_Settings{}, errors.New("resource not found")

}
