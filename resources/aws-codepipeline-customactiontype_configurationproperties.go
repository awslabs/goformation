package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::CodePipeline::CustomActionType.ConfigurationProperties AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-configurationproperties.html
type AWSCodePipelineCustomActionType_ConfigurationProperties struct {

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-configurationproperties.html#cfn-codepipeline-customactiontype-configurationproperties-description
	Description string `json:"Description"`

	// Key AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-configurationproperties.html#cfn-codepipeline-customactiontype-configurationproperties-key
	Key bool `json:"Key"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-configurationproperties.html#cfn-codepipeline-customactiontype-configurationproperties-name
	Name string `json:"Name"`

	// Queryable AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-configurationproperties.html#cfn-codepipeline-customactiontype-configurationproperties-queryable
	Queryable bool `json:"Queryable"`

	// Required AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-configurationproperties.html#cfn-codepipeline-customactiontype-configurationproperties-required
	Required bool `json:"Required"`

	// Secret AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-configurationproperties.html#cfn-codepipeline-customactiontype-configurationproperties-secret
	Secret bool `json:"Secret"`

	// Type AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-configurationproperties.html#cfn-codepipeline-customactiontype-configurationproperties-type
	Type string `json:"Type"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodePipelineCustomActionType_ConfigurationProperties) AWSCloudFormationType() string {
	return "AWS::CodePipeline::CustomActionType.ConfigurationProperties"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodePipelineCustomActionType_ConfigurationProperties) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCodePipelineCustomActionType_ConfigurationPropertiesResources retrieves all AWSCodePipelineCustomActionType_ConfigurationProperties items from a CloudFormation template
func GetAllAWSCodePipelineCustomActionType_ConfigurationProperties(template *Template) map[string]*AWSCodePipelineCustomActionType_ConfigurationProperties {

	results := map[string]*AWSCodePipelineCustomActionType_ConfigurationProperties{}
	for name, resource := range template.Resources {
		result := &AWSCodePipelineCustomActionType_ConfigurationProperties{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCodePipelineCustomActionType_ConfigurationPropertiesWithName retrieves all AWSCodePipelineCustomActionType_ConfigurationProperties items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCodePipelineCustomActionType_ConfigurationProperties(name string, template *Template) (*AWSCodePipelineCustomActionType_ConfigurationProperties, error) {

	result := &AWSCodePipelineCustomActionType_ConfigurationProperties{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCodePipelineCustomActionType_ConfigurationProperties{}, errors.New("resource not found")

}
