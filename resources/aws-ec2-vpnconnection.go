package resources

// AWS::EC2::VPNConnection AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-connection.html
type AWSEC2VPNConnection struct {
    
    // CustomerGatewayId AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-connection.html#cfn-ec2-vpnconnection-customergatewayid
    CustomerGatewayId string `json:"CustomerGatewayId"`
    
    // StaticRoutesOnly AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-connection.html#cfn-ec2-vpnconnection-StaticRoutesOnly
    StaticRoutesOnly bool `json:"StaticRoutesOnly"`
    
    // Tags AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-connection.html#cfn-ec2-vpnconnection-tags
    Tags []AWSEC2VPNConnectionTag `json:"Tags"`
    
    // Type AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-connection.html#cfn-ec2-vpnconnection-type
    Type string `json:"Type"`
    
    // VpnGatewayId AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpn-connection.html#cfn-ec2-vpnconnection-vpngatewayid
    VpnGatewayId string `json:"VpnGatewayId"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2VPNConnection) AWSCloudFormationType() string {
    return "AWS::EC2::VPNConnection"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2VPNConnection) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}