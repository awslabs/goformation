package resources

// AWS::IAM::AccessKey AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-accesskey.html
type AWSIAMAccessKey struct {

	// Serial AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-accesskey.html#cfn-iam-accesskey-serial

	Serial int64 `json:"Serial"`

	// Status AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-accesskey.html#cfn-iam-accesskey-status

	Status string `json:"Status"`

	// UserName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-accesskey.html#cfn-iam-accesskey-username

	UserName string `json:"UserName"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIAMAccessKey) AWSCloudFormationType() string {
	return "AWS::IAM::AccessKey"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIAMAccessKey) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
