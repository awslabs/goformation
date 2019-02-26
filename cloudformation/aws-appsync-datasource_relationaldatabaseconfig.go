package cloudformation

// AWSAppSyncDataSource_RelationalDatabaseConfig AWS CloudFormation Resource (AWS::AppSync::DataSource.RelationalDatabaseConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-relationaldatabaseconfig.html
type AWSAppSyncDataSource_RelationalDatabaseConfig struct {

	// RdsHttpEndpointConfig AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-relationaldatabaseconfig.html#cfn-appsync-datasource-relationaldatabaseconfig-rdshttpendpointconfig
	RdsHttpEndpointConfig *AWSAppSyncDataSource_RdsHttpEndpointConfig `json:"RdsHttpEndpointConfig,omitempty"`

	// RelationalDatabaseSourceType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-datasource-relationaldatabaseconfig.html#cfn-appsync-datasource-relationaldatabaseconfig-relationaldatabasesourcetype
	RelationalDatabaseSourceType string `json:"RelationalDatabaseSourceType,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSAppSyncDataSource_RelationalDatabaseConfig) AWSCloudFormationType() string {
	return "AWS::AppSync::DataSource.RelationalDatabaseConfig"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSAppSyncDataSource_RelationalDatabaseConfig) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
