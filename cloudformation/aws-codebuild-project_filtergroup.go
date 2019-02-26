package cloudformation

// AWSCodeBuildProject_FilterGroup AWS CloudFormation Resource (AWS::CodeBuild::Project.FilterGroup)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-filtergroup.html
type AWSCodeBuildProject_FilterGroup struct {

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodeBuildProject_FilterGroup) AWSCloudFormationType() string {
	return "AWS::CodeBuild::Project.FilterGroup"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSCodeBuildProject_FilterGroup) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
