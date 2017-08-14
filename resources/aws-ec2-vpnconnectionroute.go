package resources

// AWS::EC2::VPNConnectionRoute AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-connection-route.html
type AWSEC2VPNConnectionRoute struct {

	// DestinationCidrBlock AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-connection-route.html#cfn-ec2-vpnconnectionroute-cidrblock
	DestinationCidrBlock string `json:"DestinationCidrBlock"`

	// VpnConnectionId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-connection-route.html#cfn-ec2-vpnconnectionroute-connectionid
	VpnConnectionId string `json:"VpnConnectionId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2VPNConnectionRoute) AWSCloudFormationType() string {
	return "AWS::EC2::VPNConnectionRoute"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2VPNConnectionRoute) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
