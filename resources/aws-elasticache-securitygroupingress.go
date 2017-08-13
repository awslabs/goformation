package resources

// AWS::ElastiCache::SecurityGroupIngress AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-security-group-ingress.html
type AWSElastiCacheSecurityGroupIngress struct {
    
    // CacheSecurityGroupName AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-security-group-ingress.html#cfn-elasticache-securitygroupingress-cachesecuritygroupname
    CacheSecurityGroupName string `json:"CacheSecurityGroupName"`
    
    // EC2SecurityGroupName AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-security-group-ingress.html#cfn-elasticache-securitygroupingress-ec2securitygroupname
    EC2SecurityGroupName string `json:"EC2SecurityGroupName"`
    
    // EC2SecurityGroupOwnerId AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-security-group-ingress.html#cfn-elasticache-securitygroupingress-ec2securitygroupownerid
    EC2SecurityGroupOwnerId string `json:"EC2SecurityGroupOwnerId"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElastiCacheSecurityGroupIngress) AWSCloudFormationType() string {
    return "AWS::ElastiCache::SecurityGroupIngress"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElastiCacheSecurityGroupIngress) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}