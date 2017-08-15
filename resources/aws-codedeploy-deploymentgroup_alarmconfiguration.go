package resources

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
