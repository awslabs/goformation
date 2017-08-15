package resources

// AWS::S3::Bucket.RoutingRule AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules.html
type AWSS3Bucket_RoutingRule struct {

	// RedirectRule AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules.html#cfn-s3-websiteconfiguration-routingrules-redirectrule
	RedirectRule AWSS3Bucket_RedirectRule `json:"RedirectRule"`

	// RoutingRuleCondition AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules.html#cfn-s3-websiteconfiguration-routingrules-routingrulecondition
	RoutingRuleCondition AWSS3Bucket_RoutingRuleCondition `json:"RoutingRuleCondition"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3Bucket_RoutingRule) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.RoutingRule"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSS3Bucket_RoutingRule) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
