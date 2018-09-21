package cloudformation

// AWSCodeDeployDeploymentGroup_EC2TagSetListObject AWS CloudFormation Resource (AWS::CodeDeploy::DeploymentGroup.EC2TagSetListObject)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-ec2tagsetlistobject.html
type AWSCodeDeployDeploymentGroup_EC2TagSetListObject struct {

	// Ec2TagGroup AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codedeploy-deploymentgroup-ec2tagsetlistobject.html#cfn-codedeploy-deploymentgroup-ec2tagsetlistobject-ec2taggroup
	Ec2TagGroup []AWSCodeDeployDeploymentGroup_EC2TagFilter `json:"Ec2TagGroup,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodeDeployDeploymentGroup_EC2TagSetListObject) AWSCloudFormationType() string {
	return "AWS::CodeDeploy::DeploymentGroup.EC2TagSetListObject"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSCodeDeployDeploymentGroup_EC2TagSetListObject) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
