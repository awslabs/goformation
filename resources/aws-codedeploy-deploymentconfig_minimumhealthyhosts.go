package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::CodeDeploy::DeploymentConfig.MinimumHealthyHosts AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentconfig-minimumhealthyhosts.html
type AWSCodeDeployDeploymentConfig_MinimumHealthyHosts struct {

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentconfig-minimumhealthyhosts.html#cfn-codedeploy-deploymentconfig-minimumhealthyhosts-type
	Type string `json:"Type"`

	// Value AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentconfig-minimumhealthyhosts.html#cfn-codedeploy-deploymentconfig-minimumhealthyhosts-value
	Value int64 `json:"Value"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodeDeployDeploymentConfig_MinimumHealthyHosts) AWSCloudFormationType() string {
	return "AWS::CodeDeploy::DeploymentConfig.MinimumHealthyHosts"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodeDeployDeploymentConfig_MinimumHealthyHosts) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCodeDeployDeploymentConfig_MinimumHealthyHostsResources retrieves all AWSCodeDeployDeploymentConfig_MinimumHealthyHosts items from a CloudFormation template
func GetAllAWSCodeDeployDeploymentConfig_MinimumHealthyHosts(template *Template) map[string]*AWSCodeDeployDeploymentConfig_MinimumHealthyHosts {

	results := map[string]*AWSCodeDeployDeploymentConfig_MinimumHealthyHosts{}
	for name, resource := range template.Resources {
		result := &AWSCodeDeployDeploymentConfig_MinimumHealthyHosts{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCodeDeployDeploymentConfig_MinimumHealthyHostsWithName retrieves all AWSCodeDeployDeploymentConfig_MinimumHealthyHosts items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCodeDeployDeploymentConfig_MinimumHealthyHosts(name string, template *Template) (*AWSCodeDeployDeploymentConfig_MinimumHealthyHosts, error) {

	result := &AWSCodeDeployDeploymentConfig_MinimumHealthyHosts{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCodeDeployDeploymentConfig_MinimumHealthyHosts{}, errors.New("resource not found")

}
