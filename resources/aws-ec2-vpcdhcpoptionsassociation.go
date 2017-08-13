package resources

// AWS::EC2::VPCDHCPOptionsAssociation AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc-dhcp-options-assoc.html
type AWSEC2VPCDHCPOptionsAssociation struct {
    
    // DhcpOptionsId AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc-dhcp-options-assoc.html#cfn-ec2-vpcdhcpoptionsassociation-dhcpoptionsid
    DhcpOptionsId string `json:"DhcpOptionsId"`
    
    // VpcId AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-vpc-dhcp-options-assoc.html#cfn-ec2-vpcdhcpoptionsassociation-vpcid
    VpcId string `json:"VpcId"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2VPCDHCPOptionsAssociation) AWSCloudFormationType() string {
    return "AWS::EC2::VPCDHCPOptionsAssociation"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2VPCDHCPOptionsAssociation) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}