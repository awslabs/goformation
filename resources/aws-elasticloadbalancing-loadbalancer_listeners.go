package resources

// AWS::ElasticLoadBalancing::LoadBalancer.Listeners AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-listener.html
type AWSElasticLoadBalancingLoadBalancerListeners struct {
    
    // InstancePort AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-listener.html#cfn-ec2-elb-listener-instanceport
    InstancePort string `json:"InstancePort"`
    
    // InstanceProtocol AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-listener.html#cfn-ec2-elb-listener-instanceprotocol
    InstanceProtocol string `json:"InstanceProtocol"`
    
    // LoadBalancerPort AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-listener.html#cfn-ec2-elb-listener-loadbalancerport
    LoadBalancerPort string `json:"LoadBalancerPort"`
    
    // PolicyNames AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-listener.html#cfn-ec2-elb-listener-policynames
    PolicyNames []string `json:"PolicyNames"`
    
    // Protocol AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-listener.html#cfn-ec2-elb-listener-protocol
    Protocol string `json:"Protocol"`
    
    // SSLCertificateId AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-listener.html#cfn-ec2-elb-listener-sslcertificateid
    SSLCertificateId string `json:"SSLCertificateId"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticLoadBalancingLoadBalancerListeners) AWSCloudFormationType() string {
    return "AWS::ElasticLoadBalancing::LoadBalancer.Listeners"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticLoadBalancingLoadBalancerListeners) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}