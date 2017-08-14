package resources

// AWS::EC2::EIP AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip.html
type AWSEC2EIP struct {

	// Domain AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip.html#cfn-ec2-eip-domain

	Domain string `json:"Domain"`

	// InstanceId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-eip.html#cfn-ec2-eip-instanceid

	InstanceId string `json:"InstanceId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2EIP) AWSCloudFormationType() string {
	return "AWS::EC2::EIP"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2EIP) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
