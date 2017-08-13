package resources

// AWS::ElasticLoadBalancing::LoadBalancer AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html
type AWSElasticLoadBalancingLoadBalancer struct {
    
    // AccessLoggingPolicy AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html#cfn-ec2-elb-accessloggingpolicy
    AccessLoggingPolicy AWSElasticLoadBalancingLoadBalancerAccessLoggingPolicy `json:"AccessLoggingPolicy"`
    
    // AppCookieStickinessPolicy AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html#cfn-ec2-elb-appcookiestickinesspolicy
    AppCookieStickinessPolicy []AWSElasticLoadBalancingLoadBalancerAppCookieStickinessPolicy `json:"AppCookieStickinessPolicy"`
    
    // AvailabilityZones AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html#cfn-ec2-elb-availabilityzones
    AvailabilityZones []string `json:"AvailabilityZones"`
    
    // ConnectionDrainingPolicy AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html#cfn-ec2-elb-connectiondrainingpolicy
    ConnectionDrainingPolicy AWSElasticLoadBalancingLoadBalancerConnectionDrainingPolicy `json:"ConnectionDrainingPolicy"`
    
    // ConnectionSettings AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html#cfn-ec2-elb-connectionsettings
    ConnectionSettings AWSElasticLoadBalancingLoadBalancerConnectionSettings `json:"ConnectionSettings"`
    
    // CrossZone AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html#cfn-ec2-elb-crosszone
    CrossZone bool `json:"CrossZone"`
    
    // HealthCheck AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html#cfn-ec2-elb-healthcheck
    HealthCheck AWSElasticLoadBalancingLoadBalancerHealthCheck `json:"HealthCheck"`
    
    // Instances AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html#cfn-ec2-elb-instances
    Instances []string `json:"Instances"`
    
    // LBCookieStickinessPolicy AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html#cfn-ec2-elb-lbcookiestickinesspolicy
    LBCookieStickinessPolicy []AWSElasticLoadBalancingLoadBalancerLBCookieStickinessPolicy `json:"LBCookieStickinessPolicy"`
    
    // Listeners AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html#cfn-ec2-elb-listeners
    Listeners []AWSElasticLoadBalancingLoadBalancerListeners `json:"Listeners"`
    
    // LoadBalancerName AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html#cfn-ec2-elb-elbname
    LoadBalancerName string `json:"LoadBalancerName"`
    
    // Policies AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html#cfn-ec2-elb-policies
    Policies []AWSElasticLoadBalancingLoadBalancerPolicies `json:"Policies"`
    
    // Scheme AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html#cfn-ec2-elb-scheme
    Scheme string `json:"Scheme"`
    
    // SecurityGroups AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html#cfn-ec2-elb-securitygroups
    SecurityGroups []string `json:"SecurityGroups"`
    
    // Subnets AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html#cfn-ec2-elb-subnets
    Subnets []string `json:"Subnets"`
    
    // Tags AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb.html#cfn-elasticloadbalancing-loadbalancer-tags
    Tags []AWSElasticLoadBalancingLoadBalancerTag `json:"Tags"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticLoadBalancingLoadBalancer) AWSCloudFormationType() string {
    return "AWS::ElasticLoadBalancing::LoadBalancer"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticLoadBalancingLoadBalancer) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}