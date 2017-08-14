package resources

// AWS::Cognito::UserPool.SchemaAttribute AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-schemaattribute.html
type AWSCognitoUserPool_SchemaAttribute struct {

	// AttributeDataType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-schemaattribute.html#cfn-cognito-userpool-schemaattribute-attributedatatype

	AttributeDataType string `json:"AttributeDataType"`

	// DeveloperOnlyAttribute AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-schemaattribute.html#cfn-cognito-userpool-schemaattribute-developeronlyattribute

	DeveloperOnlyAttribute bool `json:"DeveloperOnlyAttribute"`

	// Mutable AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-schemaattribute.html#cfn-cognito-userpool-schemaattribute-mutable

	Mutable bool `json:"Mutable"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-schemaattribute.html#cfn-cognito-userpool-schemaattribute-name

	Name string `json:"Name"`

	// NumberAttributeConstraints AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-schemaattribute.html#cfn-cognito-userpool-schemaattribute-numberattributeconstraints

	NumberAttributeConstraints AWSCognitoUserPool_NumberAttributeConstraints `json:"NumberAttributeConstraints"`

	// Required AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-schemaattribute.html#cfn-cognito-userpool-schemaattribute-required

	Required bool `json:"Required"`

	// StringAttributeConstraints AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-schemaattribute.html#cfn-cognito-userpool-schemaattribute-stringattributeconstraints

	StringAttributeConstraints AWSCognitoUserPool_StringAttributeConstraints `json:"StringAttributeConstraints"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoUserPool_SchemaAttribute) AWSCloudFormationType() string {
	return "AWS::Cognito::UserPool.SchemaAttribute"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCognitoUserPool_SchemaAttribute) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
