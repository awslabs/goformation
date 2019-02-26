package cloudformation

// AWSAppSyncDataSource_AwsIamConfig AWS CloudFormation Resource (AWS::AppSync::DataSource.AwsIamConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-awsiamconfig.html
type AWSAppSyncDataSource_AwsIamConfig struct {

	// SigningRegion AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-awsiamconfig.html#cfn-appsync-datasource-awsiamconfig-signingregion
	SigningRegion string `json:"SigningRegion,omitempty"`

	// SigningServiceName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-awsiamconfig.html#cfn-appsync-datasource-awsiamconfig-signingservicename
	SigningServiceName string `json:"SigningServiceName,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSAppSyncDataSource_AwsIamConfig) AWSCloudFormationType() string {
	return "AWS::AppSync::DataSource.AwsIamConfig"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSAppSyncDataSource_AwsIamConfig) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
