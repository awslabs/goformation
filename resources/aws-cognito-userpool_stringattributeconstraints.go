package resources

// AWS::Cognito::UserPool.StringAttributeConstraints AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-stringattributeconstraints.html
type AWSCognitoUserPoolStringAttributeConstraints struct {

	// MaxLength AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-stringattributeconstraints.html#cfn-cognito-userpool-stringattributeconstraints-maxlength
	MaxLength string `json:"MaxLength"`

	// MinLength AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-stringattributeconstraints.html#cfn-cognito-userpool-stringattributeconstraints-minlength
	MinLength string `json:"MinLength"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoUserPoolStringAttributeConstraints) AWSCloudFormationType() string {
	return "AWS::Cognito::UserPool.StringAttributeConstraints"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCognitoUserPoolStringAttributeConstraints) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
