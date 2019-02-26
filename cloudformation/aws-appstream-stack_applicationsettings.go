package cloudformation

// AWSAppStreamStack_ApplicationSettings AWS CloudFormation Resource (AWS::AppStream::Stack.ApplicationSettings)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-applicationsettings.html
type AWSAppStreamStack_ApplicationSettings struct {

	// Enabled AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-applicationsettings.html#cfn-appstream-stack-applicationsettings-enabled
	Enabled bool `json:"Enabled,omitempty"`

	// SettingsGroup AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appstream-stack-applicationsettings.html#cfn-appstream-stack-applicationsettings-settingsgroup
	SettingsGroup string `json:"SettingsGroup,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSAppStreamStack_ApplicationSettings) AWSCloudFormationType() string {
	return "AWS::AppStream::Stack.ApplicationSettings"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSAppStreamStack_ApplicationSettings) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
