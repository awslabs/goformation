package resources

// AWS::EC2::SecurityGroupEgress AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-security-group-egress.html
type AWSEC2SecurityGroupEgress struct {
    
    // CidrIp AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-security-group-egress.html#cfn-ec2-securitygroupegress-cidrip
    CidrIp string `json:"CidrIp"`
    
    // CidrIpv6 AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-security-group-egress.html#cfn-ec2-securitygroupegress-cidripv6
    CidrIpv6 string `json:"CidrIpv6"`
    
    // DestinationPrefixListId AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-security-group-egress.html#cfn-ec2-securitygroupegress-destinationprefixlistid
    DestinationPrefixListId string `json:"DestinationPrefixListId"`
    
    // DestinationSecurityGroupId AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-security-group-egress.html#cfn-ec2-securitygroupegress-destinationsecuritygroupid
    DestinationSecurityGroupId string `json:"DestinationSecurityGroupId"`
    
    // FromPort AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-security-group-egress.html#cfn-ec2-securitygroupegress-fromport
    FromPort int64 `json:"FromPort"`
    
    // GroupId AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-security-group-egress.html#cfn-ec2-securitygroupegress-groupid
    GroupId string `json:"GroupId"`
    
    // IpProtocol AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-security-group-egress.html#cfn-ec2-securitygroupegress-ipprotocol
    IpProtocol string `json:"IpProtocol"`
    
    // ToPort AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-security-group-egress.html#cfn-ec2-securitygroupegress-toport
    ToPort int64 `json:"ToPort"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2SecurityGroupEgress) AWSCloudFormationType() string {
    return "AWS::EC2::SecurityGroupEgress"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2SecurityGroupEgress) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}