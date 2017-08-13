package resources

// AWS::ElasticLoadBalancingV2::LoadBalancer AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-loadbalancer.html
type AWSElasticLoadBalancingV2LoadBalancer struct {
    
    // IpAddressType AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-loadbalancer.html#cfn-elasticloadbalancingv2-loadbalancer-ipaddresstype
    IpAddressType string `json:"IpAddressType"`
    
    // LoadBalancerAttributes AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-loadbalancer.html#cfn-elasticloadbalancingv2-loadbalancer-loadbalancerattributes
    LoadBalancerAttributes []AWSElasticLoadBalancingV2LoadBalancerLoadBalancerAttribute `json:"LoadBalancerAttributes"`
    
    // Name AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-loadbalancer.html#cfn-elasticloadbalancingv2-loadbalancer-name
    Name string `json:"Name"`
    
    // Scheme AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-loadbalancer.html#cfn-elasticloadbalancingv2-loadbalancer-scheme
    Scheme string `json:"Scheme"`
    
    // SecurityGroups AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-loadbalancer.html#cfn-elasticloadbalancingv2-loadbalancer-securitygroups
    SecurityGroups []string `json:"SecurityGroups"`
    
    // Subnets AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-loadbalancer.html#cfn-elasticloadbalancingv2-loadbalancer-subnets
    Subnets []string `json:"Subnets"`
    
    // Tags AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticloadbalancingv2-loadbalancer.html#cfn-elasticloadbalancingv2-loadbalancer-tags
    Tags []AWSElasticLoadBalancingV2LoadBalancerTag `json:"Tags"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticLoadBalancingV2LoadBalancer) AWSCloudFormationType() string {
    return "AWS::ElasticLoadBalancingV2::LoadBalancer"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticLoadBalancingV2LoadBalancer) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}