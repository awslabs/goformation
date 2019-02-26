package cloudformation

// AWSAppStreamDirectoryConfig_ServiceAccountCredentials AWS CloudFormation Resource (AWS::AppStream::DirectoryConfig.ServiceAccountCredentials)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-directoryconfig-serviceaccountcredentials.html
type AWSAppStreamDirectoryConfig_ServiceAccountCredentials struct {

	// AccountName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-directoryconfig-serviceaccountcredentials.html#cfn-appstream-directoryconfig-serviceaccountcredentials-accountname
	AccountName string `json:"AccountName,omitempty"`

	// AccountPassword AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-directoryconfig-serviceaccountcredentials.html#cfn-appstream-directoryconfig-serviceaccountcredentials-accountpassword
	AccountPassword string `json:"AccountPassword,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSAppStreamDirectoryConfig_ServiceAccountCredentials) AWSCloudFormationType() string {
	return "AWS::AppStream::DirectoryConfig.ServiceAccountCredentials"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSAppStreamDirectoryConfig_ServiceAccountCredentials) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
