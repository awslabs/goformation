package resources

// AWS::EC2::Host AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-host.html
type AWSEC2Host struct {

	// AutoPlacement AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-host.html#cfn-ec2-host-autoplacement

	AutoPlacement string `json:"AutoPlacement"`

	// AvailabilityZone AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-host.html#cfn-ec2-host-availabilityzone

	AvailabilityZone string `json:"AvailabilityZone"`

	// InstanceType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-host.html#cfn-ec2-host-instancetype

	InstanceType string `json:"InstanceType"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2Host) AWSCloudFormationType() string {
	return "AWS::EC2::Host"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2Host) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
