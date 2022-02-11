package pinpointemail

import (
	"github.com/awslabs/goformation/v6/cloudformation/policies"
)

// ConfigurationSet_SendingOptions AWS CloudFormation Resource (AWS::PinpointEmail::ConfigurationSet.SendingOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpointemail-configurationset-sendingoptions.html
type ConfigurationSet_SendingOptions struct {

	// SendingEnabled AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpointemail-configurationset-sendingoptions.html#cfn-pinpointemail-configurationset-sendingoptions-sendingenabled
	SendingEnabled *bool `json:"SendingEnabled,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *ConfigurationSet_SendingOptions) AWSCloudFormationType() string {
	return "AWS::PinpointEmail::ConfigurationSet.SendingOptions"
}
