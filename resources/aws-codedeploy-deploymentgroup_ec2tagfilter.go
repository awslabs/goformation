package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::CodeDeploy::DeploymentGroup.EC2TagFilter AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-ec2tagfilters.html
type AWSCodeDeployDeploymentGroup_EC2TagFilter struct {

	// Key AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-ec2tagfilters.html#cfn-properties-codedeploy-deploymentgroup-ec2tagfilters-key

	Key string `json:"Key"`

	// Type AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-ec2tagfilters.html#cfn-properties-codedeploy-deploymentgroup-ec2tagfilters-type

	Type string `json:"Type"`

	// Value AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-ec2tagfilters.html#cfn-properties-codedeploy-deploymentgroup-ec2tagfilters-value

	Value string `json:"Value"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodeDeployDeploymentGroup_EC2TagFilter) AWSCloudFormationType() string {
	return "AWS::CodeDeploy::DeploymentGroup.EC2TagFilter"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodeDeployDeploymentGroup_EC2TagFilter) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCodeDeployDeploymentGroup_EC2TagFilterResources retrieves all AWSCodeDeployDeploymentGroup_EC2TagFilter items from a CloudFormation template
func GetAllAWSCodeDeployDeploymentGroup_EC2TagFilter(template *Template) map[string]*AWSCodeDeployDeploymentGroup_EC2TagFilter {

	results := map[string]*AWSCodeDeployDeploymentGroup_EC2TagFilter{}
	for name, resource := range template.Resources {
		result := &AWSCodeDeployDeploymentGroup_EC2TagFilter{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCodeDeployDeploymentGroup_EC2TagFilterWithName retrieves all AWSCodeDeployDeploymentGroup_EC2TagFilter items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCodeDeployDeploymentGroup_EC2TagFilter(name string, template *Template) (*AWSCodeDeployDeploymentGroup_EC2TagFilter, error) {

	result := &AWSCodeDeployDeploymentGroup_EC2TagFilter{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCodeDeployDeploymentGroup_EC2TagFilter{}, errors.New("resource not found")

}
