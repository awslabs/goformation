package resources

// AWS::EC2::SecurityGroupIngress AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-ingress.html
type AWSEC2SecurityGroupIngress struct {
    
    // CidrIp AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-ingress.html#cfn-ec2-security-group-ingress-cidrip
    CidrIp string `json:"CidrIp"`
    
    // CidrIpv6 AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-ingress.html#cfn-ec2-security-group-ingress-cidripv6
    CidrIpv6 string `json:"CidrIpv6"`
    
    // FromPort AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-ingress.html#cfn-ec2-security-group-ingress-fromport
    FromPort int64 `json:"FromPort"`
    
    // GroupId AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-ingress.html#cfn-ec2-security-group-ingress-groupid
    GroupId string `json:"GroupId"`
    
    // GroupName AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-ingress.html#cfn-ec2-security-group-ingress-groupname
    GroupName string `json:"GroupName"`
    
    // IpProtocol AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-ingress.html#cfn-ec2-security-group-ingress-ipprotocol
    IpProtocol string `json:"IpProtocol"`
    
    // SourceSecurityGroupId AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-ingress.html#cfn-ec2-security-group-ingress-sourcesecuritygroupid
    SourceSecurityGroupId string `json:"SourceSecurityGroupId"`
    
    // SourceSecurityGroupName AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-ingress.html#cfn-ec2-security-group-ingress-sourcesecuritygroupname
    SourceSecurityGroupName string `json:"SourceSecurityGroupName"`
    
    // SourceSecurityGroupOwnerId AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-ingress.html#cfn-ec2-security-group-ingress-sourcesecuritygroupownerid
    SourceSecurityGroupOwnerId string `json:"SourceSecurityGroupOwnerId"`
    
    // ToPort AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-security-group-ingress.html#cfn-ec2-security-group-ingress-toport
    ToPort int64 `json:"ToPort"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2SecurityGroupIngress) AWSCloudFormationType() string {
    return "AWS::EC2::SecurityGroupIngress"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEC2SecurityGroupIngress) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}