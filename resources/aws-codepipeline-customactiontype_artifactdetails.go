package resources

// AWS::CodePipeline::CustomActionType.ArtifactDetails AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-artifactdetails.html
type AWSCodePipelineCustomActionTypeArtifactDetails struct {
    
    // MaximumCount AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-artifactdetails.html#cfn-codepipeline-customactiontype-artifactdetails-maximumcount
    MaximumCount int64 `json:"MaximumCount"`
    
    // MinimumCount AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-customactiontype-artifactdetails.html#cfn-codepipeline-customactiontype-artifactdetails-minimumcount
    MinimumCount int64 `json:"MinimumCount"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodePipelineCustomActionTypeArtifactDetails) AWSCloudFormationType() string {
    return "AWS::CodePipeline::CustomActionType.ArtifactDetails"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodePipelineCustomActionTypeArtifactDetails) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}