package cloudformation

// AWSAppSyncResolver_PipelineConfig AWS CloudFormation Resource (AWS::AppSync::Resolver.PipelineConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-resolver-pipelineconfig.html
type AWSAppSyncResolver_PipelineConfig struct {

	// Functions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-appsync-resolver-pipelineconfig.html#cfn-appsync-resolver-pipelineconfig-functions
	Functions []string `json:"Functions,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSAppSyncResolver_PipelineConfig) AWSCloudFormationType() string {
	return "AWS::AppSync::Resolver.PipelineConfig"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSAppSyncResolver_PipelineConfig) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
