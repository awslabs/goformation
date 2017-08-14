package resources

// AWS::Route53::HealthCheck AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-healthcheck.html
type AWSRoute53HealthCheck struct {

	// HealthCheckConfig AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-healthcheck.html#cfn-route53-healthcheck-healthcheckconfig

	HealthCheckConfig AWSRoute53HealthCheck_HealthCheckConfig `json:"HealthCheckConfig"`

	// HealthCheckTags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-healthcheck.html#cfn-route53-healthcheck-healthchecktags

	HealthCheckTags []AWSRoute53HealthCheck_HealthCheckTag `json:"HealthCheckTags"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRoute53HealthCheck) AWSCloudFormationType() string {
	return "AWS::Route53::HealthCheck"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSRoute53HealthCheck) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
