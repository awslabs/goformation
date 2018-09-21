package cloudformation

// AWSCodeDeployDeploymentGroup_EC2TagSet AWS CloudFormation Resource (AWS::CodeDeploy::DeploymentGroup.EC2TagSet)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-ec2tagset.html
type AWSCodeDeployDeploymentGroup_EC2TagSet struct {

	// Ec2TagSetList AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-ec2tagset.html#cfn-codedeploy-deploymentgroup-ec2tagset-ec2tagsetlist
	Ec2TagSetList []AWSCodeDeployDeploymentGroup_EC2TagSetListObject `json:"Ec2TagSetList,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodeDeployDeploymentGroup_EC2TagSet) AWSCloudFormationType() string {
	return "AWS::CodeDeploy::DeploymentGroup.EC2TagSet"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSCodeDeployDeploymentGroup_EC2TagSet) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
