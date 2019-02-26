package cloudformation

// AWSFSxFileSystem_TagEntry AWS CloudFormation Resource (AWS::FSx::FileSystem.TagEntry)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-tagentry.html
type AWSFSxFileSystem_TagEntry struct {

	// Key AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-tagentry.html#cfn-fsx-filesystem-tagentry-key
	Key string `json:"Key,omitempty"`

	// Value AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-fsx-filesystem-tagentry.html#cfn-fsx-filesystem-tagentry-value
	Value string `json:"Value,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSFSxFileSystem_TagEntry) AWSCloudFormationType() string {
	return "AWS::FSx::FileSystem.TagEntry"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSFSxFileSystem_TagEntry) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
