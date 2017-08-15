package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::CodeDeploy::DeploymentGroup.AlarmConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-alarmconfiguration.html
type AWSCodeDeployDeploymentGroup_AlarmConfiguration struct {

	// Alarms AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-alarmconfiguration.html#cfn-codedeploy-deploymentgroup-alarmconfiguration-alarms

	Alarms []AWSCodeDeployDeploymentGroup_Alarm `json:"Alarms"`

	// Enabled AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-alarmconfiguration.html#cfn-codedeploy-deploymentgroup-alarmconfiguration-enabled

	Enabled bool `json:"Enabled"`

	// IgnorePollAlarmFailure AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-alarmconfiguration.html#cfn-codedeploy-deploymentgroup-alarmconfiguration-ignorepollalarmfailure

	IgnorePollAlarmFailure bool `json:"IgnorePollAlarmFailure"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodeDeployDeploymentGroup_AlarmConfiguration) AWSCloudFormationType() string {
	return "AWS::CodeDeploy::DeploymentGroup.AlarmConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodeDeployDeploymentGroup_AlarmConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCodeDeployDeploymentGroup_AlarmConfigurationResources retrieves all AWSCodeDeployDeploymentGroup_AlarmConfiguration items from a CloudFormation template
func GetAllAWSCodeDeployDeploymentGroup_AlarmConfiguration(template *Template) map[string]*AWSCodeDeployDeploymentGroup_AlarmConfiguration {

	results := map[string]*AWSCodeDeployDeploymentGroup_AlarmConfiguration{}
	for name, resource := range template.Resources {
		result := &AWSCodeDeployDeploymentGroup_AlarmConfiguration{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCodeDeployDeploymentGroup_AlarmConfigurationWithName retrieves all AWSCodeDeployDeploymentGroup_AlarmConfiguration items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCodeDeployDeploymentGroup_AlarmConfiguration(name string, template *Template) (*AWSCodeDeployDeploymentGroup_AlarmConfiguration, error) {

	result := &AWSCodeDeployDeploymentGroup_AlarmConfiguration{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCodeDeployDeploymentGroup_AlarmConfiguration{}, errors.New("resource not found")

}
