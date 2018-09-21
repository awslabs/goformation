package cloudformation

// AWSCodeBuildProject_S3LogsConfig AWS CloudFormation Resource (AWS::CodeBuild::Project.S3LogsConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-s3logsconfig.html
type AWSCodeBuildProject_S3LogsConfig struct {

	// Location AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-s3logsconfig.html#cfn-codebuild-project-s3logsconfig-location
	Location string `json:"Location,omitempty"`

	// Status AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-s3logsconfig.html#cfn-codebuild-project-s3logsconfig-status
	Status string `json:"Status,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodeBuildProject_S3LogsConfig) AWSCloudFormationType() string {
	return "AWS::CodeBuild::Project.S3LogsConfig"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSCodeBuildProject_S3LogsConfig) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
