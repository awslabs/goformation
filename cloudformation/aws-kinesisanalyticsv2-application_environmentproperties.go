package cloudformation

// AWSKinesisAnalyticsV2Application_EnvironmentProperties AWS CloudFormation Resource (AWS::KinesisAnalyticsV2::Application.EnvironmentProperties)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-environmentproperties.html
type AWSKinesisAnalyticsV2Application_EnvironmentProperties struct {

	// PropertyGroups AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-environmentproperties.html#cfn-kinesisanalyticsv2-application-environmentproperties-propertygroups
	PropertyGroups []AWSKinesisAnalyticsV2Application_PropertyGroup `json:"PropertyGroups,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsV2Application_EnvironmentProperties) AWSCloudFormationType() string {
	return "AWS::KinesisAnalyticsV2::Application.EnvironmentProperties"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSKinesisAnalyticsV2Application_EnvironmentProperties) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
