package resources

// AWS::CodeDeploy::DeploymentGroup.Alarm AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-alarm.html
type AWSCodeDeployDeploymentGroupAlarm struct {

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-alarm.html#cfn-codedeploy-deploymentgroup-alarm-name
	Name string `json:"Name"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodeDeployDeploymentGroupAlarm) AWSCloudFormationType() string {
	return "AWS::CodeDeploy::DeploymentGroup.Alarm"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodeDeployDeploymentGroupAlarm) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
