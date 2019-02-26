package cloudformation

// AWSEC2Instance_ElasticInferenceAccelerator AWS CloudFormation Resource (AWS::EC2::Instance.ElasticInferenceAccelerator)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-elasticinferenceaccelerator.html
type AWSEC2Instance_ElasticInferenceAccelerator struct {

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-elasticinferenceaccelerator.html#cfn-ec2-instance-elasticinferenceaccelerator-type
	Type string `json:"Type,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2Instance_ElasticInferenceAccelerator) AWSCloudFormationType() string {
	return "AWS::EC2::Instance.ElasticInferenceAccelerator"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSEC2Instance_ElasticInferenceAccelerator) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
