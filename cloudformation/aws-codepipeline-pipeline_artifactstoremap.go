package cloudformation

// AWSCodePipelinePipeline_ArtifactStoreMap AWS CloudFormation Resource (AWS::CodePipeline::Pipeline.ArtifactStoreMap)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-artifactstoremap.html
type AWSCodePipelinePipeline_ArtifactStoreMap struct {

	// ArtifactStore AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-artifactstoremap.html#cfn-codepipeline-pipeline-artifactstoremap-artifactstore
	ArtifactStore *AWSCodePipelinePipeline_ArtifactStore `json:"ArtifactStore,omitempty"`

	// Region AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-artifactstoremap.html#cfn-codepipeline-pipeline-artifactstoremap-region
	Region string `json:"Region,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodePipelinePipeline_ArtifactStoreMap) AWSCloudFormationType() string {
	return "AWS::CodePipeline::Pipeline.ArtifactStoreMap"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSCodePipelinePipeline_ArtifactStoreMap) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
