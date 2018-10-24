package cloudformation

// AWSSageMakerNotebookInstanceLifecycleConfig_NotebookInstanceLifecycleHook AWS CloudFormation Resource (AWS::SageMaker::NotebookInstanceLifecycleConfig.NotebookInstanceLifecycleHook)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-notebookinstancelifecycleconfig-notebookinstancelifecyclehook.html
type AWSSageMakerNotebookInstanceLifecycleConfig_NotebookInstanceLifecycleHook struct {

	// Content AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-notebookinstancelifecycleconfig-notebookinstancelifecyclehook.html#cfn-sagemaker-notebookinstancelifecycleconfig-notebookinstancelifecyclehook-content
	Content string `json:"Content,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy

	// _dependsOn stores the logical ID of the resources to be created before this resource
	_dependsOn []string

	// _metadata stores structured data associated with this resource
	_metadata map[string]interface{}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSageMakerNotebookInstanceLifecycleConfig_NotebookInstanceLifecycleHook) AWSCloudFormationType() string {
	return "AWS::SageMaker::NotebookInstanceLifecycleConfig.NotebookInstanceLifecycleHook"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSSageMakerNotebookInstanceLifecycleConfig_NotebookInstanceLifecycleHook) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSSageMakerNotebookInstanceLifecycleConfig_NotebookInstanceLifecycleHook) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// Metadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSSageMakerNotebookInstanceLifecycleConfig_NotebookInstanceLifecycleHook) Metadata() map[string]interface{} {
	return r._metadata
}

// SetMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSSageMakerNotebookInstanceLifecycleConfig_NotebookInstanceLifecycleHook) SetMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSSageMakerNotebookInstanceLifecycleConfig_NotebookInstanceLifecycleHook) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
