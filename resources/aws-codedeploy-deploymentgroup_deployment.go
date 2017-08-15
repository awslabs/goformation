package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::CodeDeploy::DeploymentGroup.Deployment AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment.html
type AWSCodeDeployDeploymentGroup_Deployment struct {

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment.html#cfn-properties-codedeploy-deploymentgroup-deployment-description
	Description string `json:"Description"`

	// IgnoreApplicationStopFailures AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment.html#cfn-properties-codedeploy-deploymentgroup-deployment-ignoreapplicationstopfailures
	IgnoreApplicationStopFailures bool `json:"IgnoreApplicationStopFailures"`

	// Revision AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-deployment.html#cfn-properties-codedeploy-deploymentgroup-deployment-revision
	Revision AWSCodeDeployDeploymentGroup_RevisionLocation `json:"Revision"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodeDeployDeploymentGroup_Deployment) AWSCloudFormationType() string {
	return "AWS::CodeDeploy::DeploymentGroup.Deployment"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodeDeployDeploymentGroup_Deployment) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCodeDeployDeploymentGroup_DeploymentResources retrieves all AWSCodeDeployDeploymentGroup_Deployment items from a CloudFormation template
func GetAllAWSCodeDeployDeploymentGroup_Deployment(template *Template) map[string]*AWSCodeDeployDeploymentGroup_Deployment {

	results := map[string]*AWSCodeDeployDeploymentGroup_Deployment{}
	for name, resource := range template.Resources {
		result := &AWSCodeDeployDeploymentGroup_Deployment{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCodeDeployDeploymentGroup_DeploymentWithName retrieves all AWSCodeDeployDeploymentGroup_Deployment items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCodeDeployDeploymentGroup_Deployment(name string, template *Template) (*AWSCodeDeployDeploymentGroup_Deployment, error) {

	result := &AWSCodeDeployDeploymentGroup_Deployment{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCodeDeployDeploymentGroup_Deployment{}, errors.New("resource not found")

}
