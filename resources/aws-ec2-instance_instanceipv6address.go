package resources

// AWS::EC2::Instance.InstanceIpv6Address AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-instanceipv6address.html
type AWSEC2InstanceInstanceIpv6Address struct {

	// Ipv6Address AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-instance-instanceipv6address.html#cfn-ec2-instance-instanceipv6address-ipv6address
	Ipv6Address string `json:"Ipv6Address"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2InstanceInstanceIpv6Address) AWSCloudFormationType() string {
	return "AWS::EC2::Instance.InstanceIpv6Address"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2InstanceInstanceIpv6Address) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
