package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::CodeDeploy::DeploymentGroup.GitHubLocation AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment-revision-githublocation.html
type AWSCodeDeployDeploymentGroup_GitHubLocation struct {

	// CommitId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment-revision-githublocation.html#cfn-properties-codedeploy-deploymentgroup-deployment-revision-githublocation-commitid
	CommitId string `json:"CommitId"`

	// Repository AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment-revision-githublocation.html#cfn-properties-codedeploy-deploymentgroup-deployment-revision-githublocation-repository
	Repository string `json:"Repository"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodeDeployDeploymentGroup_GitHubLocation) AWSCloudFormationType() string {
	return "AWS::CodeDeploy::DeploymentGroup.GitHubLocation"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodeDeployDeploymentGroup_GitHubLocation) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCodeDeployDeploymentGroup_GitHubLocationResources retrieves all AWSCodeDeployDeploymentGroup_GitHubLocation items from a CloudFormation template
func GetAllAWSCodeDeployDeploymentGroup_GitHubLocation(template *Template) map[string]*AWSCodeDeployDeploymentGroup_GitHubLocation {

	results := map[string]*AWSCodeDeployDeploymentGroup_GitHubLocation{}
	for name, resource := range template.Resources {
		result := &AWSCodeDeployDeploymentGroup_GitHubLocation{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCodeDeployDeploymentGroup_GitHubLocationWithName retrieves all AWSCodeDeployDeploymentGroup_GitHubLocation items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCodeDeployDeploymentGroup_GitHubLocation(name string, template *Template) (*AWSCodeDeployDeploymentGroup_GitHubLocation, error) {

	result := &AWSCodeDeployDeploymentGroup_GitHubLocation{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCodeDeployDeploymentGroup_GitHubLocation{}, errors.New("resource not found")

}
