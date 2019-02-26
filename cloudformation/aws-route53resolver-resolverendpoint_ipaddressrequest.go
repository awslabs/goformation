package cloudformation

// AWSRoute53ResolverResolverEndpoint_IpAddressRequest AWS CloudFormation Resource (AWS::Route53Resolver::ResolverEndpoint.IpAddressRequest)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53resolver-resolverendpoint-ipaddressrequest.html
type AWSRoute53ResolverResolverEndpoint_IpAddressRequest struct {

	// Ip AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53resolver-resolverendpoint-ipaddressrequest.html#cfn-route53resolver-resolverendpoint-ipaddressrequest-ip
	Ip string `json:"Ip,omitempty"`

	// SubnetId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53resolver-resolverendpoint-ipaddressrequest.html#cfn-route53resolver-resolverendpoint-ipaddressrequest-subnetid
	SubnetId string `json:"SubnetId,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRoute53ResolverResolverEndpoint_IpAddressRequest) AWSCloudFormationType() string {
	return "AWS::Route53Resolver::ResolverEndpoint.IpAddressRequest"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSRoute53ResolverResolverEndpoint_IpAddressRequest) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
