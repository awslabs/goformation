package resources

// AWS::EC2::SubnetCidrBlock AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnetcidrblock.html
type AWSEC2SubnetCidrBlock struct {
    
    // Ipv6CidrBlock AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnetcidrblock.html#cfn-ec2-subnetcidrblock-ipv6cidrblock
    Ipv6CidrBlock string `json:"Ipv6CidrBlock"`
    
    // SubnetId AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-subnetcidrblock.html#cfn-ec2-subnetcidrblock-subnetid
    SubnetId string `json:"SubnetId"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2SubnetCidrBlock) AWSCloudFormationType() string {
    return "AWS::EC2::SubnetCidrBlock"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2SubnetCidrBlock) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}