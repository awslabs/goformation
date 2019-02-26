package cloudformation

// AWSKinesisAnalyticsV2Application_PropertyGroup AWS CloudFormation Resource (AWS::KinesisAnalyticsV2::Application.PropertyGroup)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-propertygroup.html
type AWSKinesisAnalyticsV2Application_PropertyGroup struct {

	// PropertyGroupId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-propertygroup.html#cfn-kinesisanalyticsv2-application-propertygroup-propertygroupid
	PropertyGroupId string `json:"PropertyGroupId,omitempty"`

	// PropertyMap AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalyticsv2-application-propertygroup.html#cfn-kinesisanalyticsv2-application-propertygroup-propertymap
	PropertyMap interface{} `json:"PropertyMap,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsV2Application_PropertyGroup) AWSCloudFormationType() string {
	return "AWS::KinesisAnalyticsV2::Application.PropertyGroup"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSKinesisAnalyticsV2Application_PropertyGroup) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
