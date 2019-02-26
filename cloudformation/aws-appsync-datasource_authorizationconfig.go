package cloudformation

// AWSAppSyncDataSource_AuthorizationConfig AWS CloudFormation Resource (AWS::AppSync::DataSource.AuthorizationConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-authorizationconfig.html
type AWSAppSyncDataSource_AuthorizationConfig struct {

	// AuthorizationType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-authorizationconfig.html#cfn-appsync-datasource-authorizationconfig-authorizationtype
	AuthorizationType string `json:"AuthorizationType,omitempty"`

	// AwsIamConfig AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-authorizationconfig.html#cfn-appsync-datasource-authorizationconfig-awsiamconfig
	AwsIamConfig *AWSAppSyncDataSource_AwsIamConfig `json:"AwsIamConfig,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSAppSyncDataSource_AuthorizationConfig) AWSCloudFormationType() string {
	return "AWS::AppSync::DataSource.AuthorizationConfig"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSAppSyncDataSource_AuthorizationConfig) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
