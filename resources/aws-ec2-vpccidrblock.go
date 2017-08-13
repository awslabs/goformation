package resources

// AWS::EC2::VPCCidrBlock AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpccidrblock.html
type AWSEC2VPCCidrBlock struct {
    
    // AmazonProvidedIpv6CidrBlock AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpccidrblock.html#cfn-ec2-vpccidrblock-amazonprovidedipv6cidrblock
    AmazonProvidedIpv6CidrBlock bool `json:"AmazonProvidedIpv6CidrBlock"`
    
    // VpcId AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpccidrblock.html#cfn-ec2-vpccidrblock-vpcid
    VpcId string `json:"VpcId"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2VPCCidrBlock) AWSCloudFormationType() string {
    return "AWS::EC2::VPCCidrBlock"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2VPCCidrBlock) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}