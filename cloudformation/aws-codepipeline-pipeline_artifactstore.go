package cloudformation

import (
	"encoding/json"
)

// AWSCodePipelinePipeline_ArtifactStore AWS CloudFormation Resource (AWS::CodePipeline::Pipeline.ArtifactStore)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-artifactstore.html
type AWSCodePipelinePipeline_ArtifactStore struct {

	// EncryptionKey AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-artifactstore.html#cfn-codepipeline-pipeline-artifactstore-encryptionkey
	EncryptionKey *AWSCodePipelinePipeline_EncryptionKey `json:"EncryptionKey,omitempty"`

	// Location AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-artifactstore.html#cfn-codepipeline-pipeline-artifactstore-location
	Location *Value `json:"Location,omitempty"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-artifactstore.html#cfn-codepipeline-pipeline-artifactstore-type
	Type *Value `json:"Type,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodePipelinePipeline_ArtifactStore) AWSCloudFormationType() string {
	return "AWS::CodePipeline::Pipeline.ArtifactStore"
}

func (r *AWSCodePipelinePipeline_ArtifactStore) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
