package resources

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
