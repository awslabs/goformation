package resources

// AWS::KMS::Alias AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-alias.html
type AWSKMSAlias struct {

	// AliasName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-alias.html#cfn-kms-alias-aliasname
	AliasName string `json:"AliasName"`

	// TargetKeyId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-alias.html#cfn-kms-alias-targetkeyid
	TargetKeyId string `json:"TargetKeyId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKMSAlias) AWSCloudFormationType() string {
	return "AWS::KMS::Alias"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKMSAlias) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
