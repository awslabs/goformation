package resources

// AWS::Route53::HostedZone AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html
type AWSRoute53HostedZone struct {

	// HostedZoneConfig AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html#cfn-route53-hostedzone-hostedzoneconfig

	HostedZoneConfig AWSRoute53HostedZone_HostedZoneConfig `json:"HostedZoneConfig"`

	// HostedZoneTags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html#cfn-route53-hostedzone-hostedzonetags

	HostedZoneTags []AWSRoute53HostedZone_HostedZoneTag `json:"HostedZoneTags"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html#cfn-route53-hostedzone-name

	Name string `json:"Name"`

	// VPCs AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone.html#cfn-route53-hostedzone-vpcs

	VPCs []AWSRoute53HostedZone_VPC `json:"VPCs"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRoute53HostedZone) AWSCloudFormationType() string {
	return "AWS::Route53::HostedZone"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSRoute53HostedZone) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
