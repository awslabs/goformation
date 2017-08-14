package resources

// AWS::Cognito::UserPool.EmailConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-emailconfiguration.html
type AWSCognitoUserPool_EmailConfiguration struct {

	// ReplyToEmailAddress AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-emailconfiguration.html#cfn-cognito-userpool-emailconfiguration-replytoemailaddress

	ReplyToEmailAddress string `json:"ReplyToEmailAddress"`

	// SourceArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-emailconfiguration.html#cfn-cognito-userpool-emailconfiguration-sourcearn

	SourceArn string `json:"SourceArn"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoUserPool_EmailConfiguration) AWSCloudFormationType() string {
	return "AWS::Cognito::UserPool.EmailConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCognitoUserPool_EmailConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
