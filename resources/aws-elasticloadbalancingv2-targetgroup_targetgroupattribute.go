package resources

// AWS::ElasticLoadBalancingV2::TargetGroup.TargetGroupAttribute AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-targetgroup-targetgroupattributes.html
type AWSElasticLoadBalancingV2TargetGroupTargetGroupAttribute struct {
    
    // Key AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-targetgroup-targetgroupattributes.html#cfn-elasticloadbalancingv2-targetgroup-targetgroupattributes-key
    Key string `json:"Key"`
    
    // Value AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-targetgroup-targetgroupattributes.html#cfn-elasticloadbalancingv2-targetgroup-targetgroupattributes-value
    Value string `json:"Value"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticLoadBalancingV2TargetGroupTargetGroupAttribute) AWSCloudFormationType() string {
    return "AWS::ElasticLoadBalancingV2::TargetGroup.TargetGroupAttribute"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticLoadBalancingV2TargetGroupTargetGroupAttribute) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}