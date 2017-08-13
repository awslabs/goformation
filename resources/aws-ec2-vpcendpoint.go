package resources

// AWS::EC2::VPCEndpoint AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpoint.html
type AWSEC2VPCEndpoint struct {
    
    // PolicyDocument AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpoint.html#cfn-ec2-vpcendpoint-policydocument
    PolicyDocument string `json:"PolicyDocument"`
    
    // RouteTableIds AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpoint.html#cfn-ec2-vpcendpoint-routetableids
    RouteTableIds []string `json:"RouteTableIds"`
    
    // ServiceName AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpoint.html#cfn-ec2-vpcendpoint-servicename
    ServiceName string `json:"ServiceName"`
    
    // VpcId AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpcendpoint.html#cfn-ec2-vpcendpoint-vpcid
    VpcId string `json:"VpcId"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2VPCEndpoint) AWSCloudFormationType() string {
    return "AWS::EC2::VPCEndpoint"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2VPCEndpoint) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}