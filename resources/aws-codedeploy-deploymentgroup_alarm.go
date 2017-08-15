package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::CodeDeploy::DeploymentGroup.Alarm AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-alarm.html
type AWSCodeDeployDeploymentGroup_Alarm struct {

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-alarm.html#cfn-codedeploy-deploymentgroup-alarm-name

	Name string `json:"Name"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodeDeployDeploymentGroup_Alarm) AWSCloudFormationType() string {
	return "AWS::CodeDeploy::DeploymentGroup.Alarm"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodeDeployDeploymentGroup_Alarm) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCodeDeployDeploymentGroup_AlarmResources retrieves all AWSCodeDeployDeploymentGroup_Alarm items from a CloudFormation template
func GetAllAWSCodeDeployDeploymentGroup_Alarm(template *Template) map[string]*AWSCodeDeployDeploymentGroup_Alarm {

	results := map[string]*AWSCodeDeployDeploymentGroup_Alarm{}
	for name, resource := range template.Resources {
		result := &AWSCodeDeployDeploymentGroup_Alarm{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCodeDeployDeploymentGroup_AlarmWithName retrieves all AWSCodeDeployDeploymentGroup_Alarm items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCodeDeployDeploymentGroup_Alarm(name string, template *Template) (*AWSCodeDeployDeploymentGroup_Alarm, error) {

	result := &AWSCodeDeployDeploymentGroup_Alarm{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCodeDeployDeploymentGroup_Alarm{}, errors.New("resource not found")

}
