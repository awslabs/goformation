package resources

// AWS::Cognito::UserPoolUser.AttributeType AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpooluser-attributetype.html
type AWSCognitoUserPoolUser_AttributeType struct {

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpooluser-attributetype.html#cfn-cognito-userpooluser-attributetype-name
	Name string `json:"Name"`

	// Value AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpooluser-attributetype.html#cfn-cognito-userpooluser-attributetype-value
	Value string `json:"Value"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoUserPoolUser_AttributeType) AWSCloudFormationType() string {
	return "AWS::Cognito::UserPoolUser.AttributeType"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCognitoUserPoolUser_AttributeType) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
