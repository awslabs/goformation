package cloudformation

// AWSCodeDeployDeploymentGroup_OnPremisesTagSet AWS CloudFormation Resource (AWS::CodeDeploy::DeploymentGroup.OnPremisesTagSet)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-onpremisestagset.html
type AWSCodeDeployDeploymentGroup_OnPremisesTagSet struct {

	// OnPremisesTagSetList AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-onpremisestagset.html#cfn-codedeploy-deploymentgroup-onpremisestagset-onpremisestagsetlist
	OnPremisesTagSetList []AWSCodeDeployDeploymentGroup_OnPremisesTagSetListObject `json:"OnPremisesTagSetList,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodeDeployDeploymentGroup_OnPremisesTagSet) AWSCloudFormationType() string {
	return "AWS::CodeDeploy::DeploymentGroup.OnPremisesTagSet"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSCodeDeployDeploymentGroup_OnPremisesTagSet) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
