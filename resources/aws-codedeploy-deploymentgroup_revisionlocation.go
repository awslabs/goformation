package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::CodeDeploy::DeploymentGroup.RevisionLocation AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment-revision.html
type AWSCodeDeployDeploymentGroup_RevisionLocation struct {

	// GitHubLocation AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment-revision.html#cfn-properties-codedeploy-deploymentgroup-deployment-revision-githublocation

	GitHubLocation AWSCodeDeployDeploymentGroup_GitHubLocation `json:"GitHubLocation"`

	// RevisionType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment-revision.html#cfn-properties-codedeploy-deploymentgroup-deployment-revision-revisiontype

	RevisionType string `json:"RevisionType"`

	// S3Location AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment-revision.html#cfn-properties-codedeploy-deploymentgroup-deployment-revision-s3location

	S3Location AWSCodeDeployDeploymentGroup_S3Location `json:"S3Location"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodeDeployDeploymentGroup_RevisionLocation) AWSCloudFormationType() string {
	return "AWS::CodeDeploy::DeploymentGroup.RevisionLocation"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodeDeployDeploymentGroup_RevisionLocation) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCodeDeployDeploymentGroup_RevisionLocationResources retrieves all AWSCodeDeployDeploymentGroup_RevisionLocation items from a CloudFormation template
func GetAllAWSCodeDeployDeploymentGroup_RevisionLocation(template *Template) map[string]*AWSCodeDeployDeploymentGroup_RevisionLocation {

	results := map[string]*AWSCodeDeployDeploymentGroup_RevisionLocation{}
	for name, resource := range template.Resources {
		result := &AWSCodeDeployDeploymentGroup_RevisionLocation{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCodeDeployDeploymentGroup_RevisionLocationWithName retrieves all AWSCodeDeployDeploymentGroup_RevisionLocation items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCodeDeployDeploymentGroup_RevisionLocation(name string, template *Template) (*AWSCodeDeployDeploymentGroup_RevisionLocation, error) {

	result := &AWSCodeDeployDeploymentGroup_RevisionLocation{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCodeDeployDeploymentGroup_RevisionLocation{}, errors.New("resource not found")

}
