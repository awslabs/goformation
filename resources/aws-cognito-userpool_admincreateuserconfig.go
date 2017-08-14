package resources

// AWS::Cognito::UserPool.AdminCreateUserConfig AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-admincreateuserconfig.html
type AWSCognitoUserPoolAdminCreateUserConfig struct {

	// AllowAdminCreateUserOnly AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-admincreateuserconfig.html#cfn-cognito-userpool-admincreateuserconfig-allowadmincreateuseronly
	AllowAdminCreateUserOnly bool `json:"AllowAdminCreateUserOnly"`

	// InviteMessageTemplate AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-admincreateuserconfig.html#cfn-cognito-userpool-admincreateuserconfig-invitemessagetemplate
	InviteMessageTemplate AWSCognitoUserPoolAdminCreateUserConfigInviteMessageTemplate `json:"InviteMessageTemplate"`

	// UnusedAccountValidityDays AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-admincreateuserconfig.html#cfn-cognito-userpool-admincreateuserconfig-unusedaccountvaliditydays
	UnusedAccountValidityDays float64 `json:"UnusedAccountValidityDays"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoUserPoolAdminCreateUserConfig) AWSCloudFormationType() string {
	return "AWS::Cognito::UserPool.AdminCreateUserConfig"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCognitoUserPoolAdminCreateUserConfig) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
