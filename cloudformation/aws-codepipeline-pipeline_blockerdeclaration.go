package cloudformation

// AWSCodePipelinePipeline_BlockerDeclaration AWS CloudFormation Resource (AWS::CodePipeline::Pipeline.BlockerDeclaration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-blockerdeclaration.html
type AWSCodePipelinePipeline_BlockerDeclaration struct {

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-blockerdeclaration.html#cfn-codepipeline-pipeline-blockerdeclaration-name
	Name string `json:"Name,omitempty"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-blockerdeclaration.html#cfn-codepipeline-pipeline-blockerdeclaration-type
	Type string `json:"Type,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodePipelinePipeline_BlockerDeclaration) AWSCloudFormationType() string {
	return "AWS::CodePipeline::Pipeline.BlockerDeclaration"
}
