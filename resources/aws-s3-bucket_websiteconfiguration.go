package resources

// AWS::S3::Bucket.WebsiteConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration.html
type AWSS3BucketWebsiteConfiguration struct {

	// ErrorDocument AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration.html#cfn-s3-websiteconfiguration-errordocument
	ErrorDocument string `json:"ErrorDocument"`

	// IndexDocument AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration.html#cfn-s3-websiteconfiguration-indexdocument
	IndexDocument string `json:"IndexDocument"`

	// RedirectAllRequestsTo AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration.html#cfn-s3-websiteconfiguration-redirectallrequeststo
	RedirectAllRequestsTo AWSS3BucketWebsiteConfigurationRedirectAllRequestsTo `json:"RedirectAllRequestsTo"`

	// RoutingRules AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-websiteconfiguration.html#cfn-s3-websiteconfiguration-routingrules
	RoutingRules []AWSS3BucketWebsiteConfigurationRoutingRule `json:"RoutingRules"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSS3BucketWebsiteConfiguration) AWSCloudFormationType() string {
	return "AWS::S3::Bucket.WebsiteConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSS3BucketWebsiteConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
