package resources

// AWS::Route53::HostedZone.HostedZoneConfig AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-hostedzoneconfig.html
type AWSRoute53HostedZoneHostedZoneConfig struct {

	// Comment AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-hostedzoneconfig.html#cfn-route53-hostedzone-hostedzoneconfig-comment
	Comment string `json:"Comment"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRoute53HostedZoneHostedZoneConfig) AWSCloudFormationType() string {
	return "AWS::Route53::HostedZone.HostedZoneConfig"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSRoute53HostedZoneHostedZoneConfig) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
