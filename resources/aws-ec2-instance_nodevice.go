package resources

// AWS::EC2::Instance.NoDevice AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-nodevice.html
type AWSEC2InstanceNoDevice struct {
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2InstanceNoDevice) AWSCloudFormationType() string {
	return "AWS::EC2::Instance.NoDevice"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2InstanceNoDevice) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
