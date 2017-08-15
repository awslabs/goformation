package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::CodeDeploy::DeploymentGroup AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html
type AWSCodeDeployDeploymentGroup struct {

	// AlarmConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-alarmconfiguration
	AlarmConfiguration AWSCodeDeployDeploymentGroup_AlarmConfiguration `json:"AlarmConfiguration"`

	// ApplicationName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-applicationname
	ApplicationName string `json:"ApplicationName"`

	// AutoScalingGroups AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-autoscalinggroups
	AutoScalingGroups []string `json:"AutoScalingGroups"`

	// Deployment AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-deployment
	Deployment AWSCodeDeployDeploymentGroup_Deployment `json:"Deployment"`

	// DeploymentConfigName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-deploymentconfigname
	DeploymentConfigName string `json:"DeploymentConfigName"`

	// DeploymentGroupName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-deploymentgroupname
	DeploymentGroupName string `json:"DeploymentGroupName"`

	// Ec2TagFilters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-ec2tagfilters
	Ec2TagFilters []AWSCodeDeployDeploymentGroup_EC2TagFilter `json:"Ec2TagFilters"`

	// OnPremisesInstanceTagFilters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-onpremisesinstancetagfilters
	OnPremisesInstanceTagFilters []AWSCodeDeployDeploymentGroup_TagFilter `json:"OnPremisesInstanceTagFilters"`

	// ServiceRoleArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-servicerolearn
	ServiceRoleArn string `json:"ServiceRoleArn"`

	// TriggerConfigurations AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codedeploy-deploymentgroup.html#cfn-codedeploy-deploymentgroup-triggerconfigurations
	TriggerConfigurations []AWSCodeDeployDeploymentGroup_TriggerConfig `json:"TriggerConfigurations"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodeDeployDeploymentGroup) AWSCloudFormationType() string {
	return "AWS::CodeDeploy::DeploymentGroup"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodeDeployDeploymentGroup) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCodeDeployDeploymentGroupResources retrieves all AWSCodeDeployDeploymentGroup items from a CloudFormation template
func GetAllAWSCodeDeployDeploymentGroupResources(template *Template) map[string]*AWSCodeDeployDeploymentGroup {

	results := map[string]*AWSCodeDeployDeploymentGroup{}
	for name, resource := range template.Resources {
		result := &AWSCodeDeployDeploymentGroup{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCodeDeployDeploymentGroupWithName retrieves all AWSCodeDeployDeploymentGroup items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSCodeDeployDeploymentGroupWithName(name string, template *Template) (*AWSCodeDeployDeploymentGroup, error) {

	result := &AWSCodeDeployDeploymentGroup{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCodeDeployDeploymentGroup{}, errors.New("resource not found")

}
