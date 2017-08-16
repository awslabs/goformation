package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSCodeDeployDeploymentConfig AWS CloudFormation Resource (AWS::CodeDeploy::DeploymentConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentconfig.html
type AWSCodeDeployDeploymentConfig struct {

	// DeploymentConfigName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentconfig.html#cfn-codedeploy-deploymentconfig-deploymentconfigname
	DeploymentConfigName string `json:"DeploymentConfigName"`

	// MinimumHealthyHosts AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentconfig.html#cfn-codedeploy-deploymentconfig-minimumhealthyhosts
	MinimumHealthyHosts AWSCodeDeployDeploymentConfig_MinimumHealthyHosts `json:"MinimumHealthyHosts"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodeDeployDeploymentConfig) AWSCloudFormationType() string {
	return "AWS::CodeDeploy::DeploymentConfig"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodeDeployDeploymentConfig) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCodeDeployDeploymentConfigResources retrieves all AWSCodeDeployDeploymentConfig items from a CloudFormation template
func (t *Template) GetAllAWSCodeDeployDeploymentConfigResources() map[string]*AWSCodeDeployDeploymentConfig {

	results := map[string]*AWSCodeDeployDeploymentConfig{}
	for name, resource := range t.Resources {
		result := &AWSCodeDeployDeploymentConfig{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCodeDeployDeploymentConfigWithName retrieves all AWSCodeDeployDeploymentConfig items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCodeDeployDeploymentConfigWithName(name string) (*AWSCodeDeployDeploymentConfig, error) {

	result := &AWSCodeDeployDeploymentConfig{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCodeDeployDeploymentConfig{}, errors.New("resource not found")

}
