package cloudformation

// AWSRoute53ResolverResolverRule_TargetAddress AWS CloudFormation Resource (AWS::Route53Resolver::ResolverRule.TargetAddress)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53resolver-resolverrule-targetaddress.html
type AWSRoute53ResolverResolverRule_TargetAddress struct {

	// Ip AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53resolver-resolverrule-targetaddress.html#cfn-route53resolver-resolverrule-targetaddress-ip
	Ip string `json:"Ip,omitempty"`

	// Port AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53resolver-resolverrule-targetaddress.html#cfn-route53resolver-resolverrule-targetaddress-port
	Port string `json:"Port,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRoute53ResolverResolverRule_TargetAddress) AWSCloudFormationType() string {
	return "AWS::Route53Resolver::ResolverRule.TargetAddress"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSRoute53ResolverResolverRule_TargetAddress) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
