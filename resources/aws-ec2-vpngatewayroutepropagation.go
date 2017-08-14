package resources

// AWS::EC2::VPNGatewayRoutePropagation AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-gatewayrouteprop.html
type AWSEC2VPNGatewayRoutePropagation struct {

	// RouteTableIds AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-gatewayrouteprop.html#cfn-ec2-vpngatewayrouteprop-routetableids
	RouteTableIds []AWSEC2VPNGatewayRoutePropagationstring `json:"RouteTableIds"`

	// VpnGatewayId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-gatewayrouteprop.html#cfn-ec2-vpngatewayrouteprop-vpngatewayid
	VpnGatewayId string `json:"VpnGatewayId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2VPNGatewayRoutePropagation) AWSCloudFormationType() string {
	return "AWS::EC2::VPNGatewayRoutePropagation"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2VPNGatewayRoutePropagation) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
