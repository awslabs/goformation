package resources

import "github.com/awslabs/goformation/cloudformation/policies"

// AWSPinpointEmailConfigurationSet_ReputationOptions AWS CloudFormation Resource (AWS::PinpointEmail::ConfigurationSet.ReputationOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpointemail-configurationset-reputationoptions.html
type AWSPinpointEmailConfigurationSet_ReputationOptions struct {

	// ReputationMetricsEnabled AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-pinpointemail-configurationset-reputationoptions.html#cfn-pinpointemail-configurationset-reputationoptions-reputationmetricsenabled
	ReputationMetricsEnabled bool `json:"ReputationMetricsEnabled,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy policies.DeletionPolicy

	// _dependsOn stores the logical ID of the resources to be created before this resource
	_dependsOn []string

	// _metadata stores structured data associated with this resource
	_metadata map[string]interface{}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSPinpointEmailConfigurationSet_ReputationOptions) AWSCloudFormationType() string {
	return "AWS::PinpointEmail::ConfigurationSet.ReputationOptions"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSPinpointEmailConfigurationSet_ReputationOptions) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSPinpointEmailConfigurationSet_ReputationOptions) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// Metadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSPinpointEmailConfigurationSet_ReputationOptions) Metadata() map[string]interface{} {
	return r._metadata
}

// SetMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSPinpointEmailConfigurationSet_ReputationOptions) SetMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSPinpointEmailConfigurationSet_ReputationOptions) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}
