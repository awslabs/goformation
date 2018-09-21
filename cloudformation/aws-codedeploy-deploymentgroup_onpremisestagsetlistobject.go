package cloudformation

// AWSCodeDeployDeploymentGroup_OnPremisesTagSetListObject AWS CloudFormation Resource (AWS::CodeDeploy::DeploymentGroup.OnPremisesTagSetListObject)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-onpremisestagsetlistobject.html
type AWSCodeDeployDeploymentGroup_OnPremisesTagSetListObject struct {

	// OnPremisesTagGroup AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-onpremisestagsetlistobject.html#cfn-codedeploy-deploymentgroup-onpremisestagsetlistobject-onpremisestaggroup
	OnPremisesTagGroup []AWSCodeDeployDeploymentGroup_TagFilter `json:"OnPremisesTagGroup,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodeDeployDeploymentGroup_OnPremisesTagSetListObject) AWSCloudFormationType() string {
	return "AWS::CodeDeploy::DeploymentGroup.OnPremisesTagSetListObject"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSCodeDeployDeploymentGroup_OnPremisesTagSetListObject) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
