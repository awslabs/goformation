package resources

// AWS::ElasticLoadBalancingV2::Listener.Certificate AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-certificates.html
type AWSElasticLoadBalancingV2Listener_Certificate struct {

	// CertificateArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-certificates.html#cfn-elasticloadbalancingv2-listener-certificates-certificatearn

	CertificateArn string `json:"CertificateArn"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticLoadBalancingV2Listener_Certificate) AWSCloudFormationType() string {
	return "AWS::ElasticLoadBalancingV2::Listener.Certificate"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticLoadBalancingV2Listener_Certificate) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
