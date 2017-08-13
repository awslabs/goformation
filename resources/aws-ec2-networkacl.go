package resources

// AWS::EC2::NetworkAcl AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl.html
type AWSEC2NetworkAcl struct {
    
    // Tags AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl.html#cfn-ec2-networkacl-tags
    Tags []AWSEC2NetworkAclTag `json:"Tags"`
    
    // VpcId AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-network-acl.html#cfn-ec2-networkacl-vpcid
    VpcId string `json:"VpcId"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2NetworkAcl) AWSCloudFormationType() string {
    return "AWS::EC2::NetworkAcl"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2NetworkAcl) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}