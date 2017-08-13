package resources

// AWS::CodePipeline::Pipeline.BlockerDeclaration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-blockers.html
type AWSCodePipelinePipelineBlockerDeclaration struct {
    
    // Name AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-blockers.html#cfn-codepipeline-pipeline-stages-blockers-name
    Name string `json:"Name"`
    
    // Type AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-stages-blockers.html#cfn-codepipeline-pipeline-stages-blockers-type
    Type string `json:"Type"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodePipelinePipelineBlockerDeclaration) AWSCloudFormationType() string {
    return "AWS::CodePipeline::Pipeline.BlockerDeclaration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodePipelinePipelineBlockerDeclaration) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}