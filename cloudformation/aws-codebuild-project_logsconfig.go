package cloudformation

// AWSCodeBuildProject_LogsConfig AWS CloudFormation Resource (AWS::CodeBuild::Project.LogsConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-logsconfig.html
type AWSCodeBuildProject_LogsConfig struct {

	// CloudWatchLogs AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-logsconfig.html#cfn-codebuild-project-logsconfig-cloudwatchlogs
	CloudWatchLogs *AWSCodeBuildProject_CloudWatchLogsConfig `json:"CloudWatchLogs,omitempty"`

	// S3Logs AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-logsconfig.html#cfn-codebuild-project-logsconfig-s3logs
	S3Logs *AWSCodeBuildProject_S3LogsConfig `json:"S3Logs,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodeBuildProject_LogsConfig) AWSCloudFormationType() string {
	return "AWS::CodeBuild::Project.LogsConfig"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSCodeBuildProject_LogsConfig) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
