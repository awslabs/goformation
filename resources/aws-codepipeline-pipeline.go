package resources

// AWS::CodePipeline::Pipeline AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html
type AWSCodePipelinePipeline struct {

	// ArtifactStore AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-artifactstore

	ArtifactStore AWSCodePipelinePipeline_ArtifactStore `json:"ArtifactStore"`

	// DisableInboundStageTransitions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-disableinboundstagetransitions

	DisableInboundStageTransitions []AWSCodePipelinePipeline_StageTransition `json:"DisableInboundStageTransitions"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-name

	Name string `json:"Name"`

	// RestartExecutionOnUpdate AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-restartexecutiononupdate

	RestartExecutionOnUpdate bool `json:"RestartExecutionOnUpdate"`

	// RoleArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-rolearn

	RoleArn string `json:"RoleArn"`

	// Stages AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-codepipeline-pipeline.html#cfn-codepipeline-pipeline-stages

	Stages []AWSCodePipelinePipeline_StageDeclaration `json:"Stages"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodePipelinePipeline) AWSCloudFormationType() string {
	return "AWS::CodePipeline::Pipeline"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodePipelinePipeline) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
