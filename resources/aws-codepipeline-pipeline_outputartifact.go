package resources

// AWS::CodePipeline::Pipeline.OutputArtifact AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-actions-outputartifacts.html
type AWSCodePipelinePipelineOutputArtifact struct {
    
    // Name AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-actions-outputartifacts.html#cfn-codepipeline-pipeline-stages-actions-outputartifacts-name
    Name string `json:"Name"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodePipelinePipelineOutputArtifact) AWSCloudFormationType() string {
    return "AWS::CodePipeline::Pipeline.OutputArtifact"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodePipelinePipelineOutputArtifact) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}