package resources

// AWS::S3::Bucket.RedirectRule AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules-redirectrule.html
type AWSS3Bucket_RedirectRule struct {

	// HostName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules-redirectrule.html#cfn-s3-websiteconfiguration-redirectrule-hostname
	HostName string `json:"HostName"`

	// HttpRedirectCode AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules-redirectrule.html#cfn-s3-websiteconfiguration-redirectrule-httpredirectcode
	HttpRedirectCode string `json:"HttpRedirectCode"`

	// Protocol AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules-redirectrule.html#cfn-s3-websiteconfiguration-redirectrule-protocol
	Protocol string `json:"Protocol"`

	// ReplaceKeyPrefixWith AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules-redirectrule.html#cfn-s3-websiteconfiguration-redirectrule-replacekeyprefixwith
	ReplaceKeyPrefixWith string `json:"ReplaceKeyPrefixWith"`

	// ReplaceKeyWith AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration-routingrules-redirectrule.html#cfn-s3-websiteconfiguration-redirectrule-replacekeywith
	ReplaceKeyWith string `json:"ReplaceKeyWith"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3Bucket_RedirectRule) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.RedirectRule"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSS3Bucket_RedirectRule) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
