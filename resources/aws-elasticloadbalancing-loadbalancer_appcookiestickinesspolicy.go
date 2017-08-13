package resources

// AWS::ElasticLoadBalancing::LoadBalancer.AppCookieStickinessPolicy AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-AppCookieStickinessPolicy.html
type AWSElasticLoadBalancingLoadBalancerAppCookieStickinessPolicy struct {
    
    // CookieName AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-AppCookieStickinessPolicy.html#cfn-elb-appcookiestickinesspolicy-cookiename
    CookieName string `json:"CookieName"`
    
    // PolicyName AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-AppCookieStickinessPolicy.html#cfn-elb-appcookiestickinesspolicy-policyname
    PolicyName string `json:"PolicyName"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticLoadBalancingLoadBalancerAppCookieStickinessPolicy) AWSCloudFormationType() string {
    return "AWS::ElasticLoadBalancing::LoadBalancer.AppCookieStickinessPolicy"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticLoadBalancingLoadBalancerAppCookieStickinessPolicy) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}