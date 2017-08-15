package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::CodeDeploy::DeploymentGroup.TriggerConfig AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-triggerconfig.html
type AWSCodeDeployDeploymentGroup_TriggerConfig struct {

	// TriggerEvents AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-triggerconfig.html#cfn-codedeploy-deploymentgroup-triggerconfig-triggerevents
	TriggerEvents []string `json:"TriggerEvents"`

	// TriggerName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-triggerconfig.html#cfn-codedeploy-deploymentgroup-triggerconfig-triggername
	TriggerName string `json:"TriggerName"`

	// TriggerTargetArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-triggerconfig.html#cfn-codedeploy-deploymentgroup-triggerconfig-triggertargetarn
	TriggerTargetArn string `json:"TriggerTargetArn"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodeDeployDeploymentGroup_TriggerConfig) AWSCloudFormationType() string {
	return "AWS::CodeDeploy::DeploymentGroup.TriggerConfig"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodeDeployDeploymentGroup_TriggerConfig) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCodeDeployDeploymentGroup_TriggerConfigResources retrieves all AWSCodeDeployDeploymentGroup_TriggerConfig items from a CloudFormation template
func GetAllAWSCodeDeployDeploymentGroup_TriggerConfig(template *Template) map[string]*AWSCodeDeployDeploymentGroup_TriggerConfig {

	results := map[string]*AWSCodeDeployDeploymentGroup_TriggerConfig{}
	for name, resource := range template.Resources {
		result := &AWSCodeDeployDeploymentGroup_TriggerConfig{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCodeDeployDeploymentGroup_TriggerConfigWithName retrieves all AWSCodeDeployDeploymentGroup_TriggerConfig items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCodeDeployDeploymentGroup_TriggerConfig(name string, template *Template) (*AWSCodeDeployDeploymentGroup_TriggerConfig, error) {

	result := &AWSCodeDeployDeploymentGroup_TriggerConfig{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCodeDeployDeploymentGroup_TriggerConfig{}, errors.New("resource not found")

}
