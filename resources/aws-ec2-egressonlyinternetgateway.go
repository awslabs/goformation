package resources

// AWS::EC2::EgressOnlyInternetGateway AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-egressonlyinternetgateway.html
type AWSEC2EgressOnlyInternetGateway struct {

	// VpcId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-egressonlyinternetgateway.html#cfn-ec2-egressonlyinternetgateway-vpcid

	VpcId string `json:"VpcId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2EgressOnlyInternetGateway) AWSCloudFormationType() string {
	return "AWS::EC2::EgressOnlyInternetGateway"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2EgressOnlyInternetGateway) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
