package resources

// AWS::ElasticLoadBalancing::LoadBalancer.ConnectionDrainingPolicy AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-connectiondrainingpolicy.html
type AWSElasticLoadBalancingLoadBalancer_ConnectionDrainingPolicy struct {

	// Enabled AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-connectiondrainingpolicy.html#cfn-elb-connectiondrainingpolicy-enabled

	Enabled bool `json:"Enabled"`

	// Timeout AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-connectiondrainingpolicy.html#cfn-elb-connectiondrainingpolicy-timeout

	Timeout int64 `json:"Timeout"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticLoadBalancingLoadBalancer_ConnectionDrainingPolicy) AWSCloudFormationType() string {
	return "AWS::ElasticLoadBalancing::LoadBalancer.ConnectionDrainingPolicy"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticLoadBalancingLoadBalancer_ConnectionDrainingPolicy) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
