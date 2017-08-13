package resources

// AWS::CodePipeline::Pipeline.EncryptionKey AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-artifactstore-encryptionkey.html
type AWSCodePipelinePipelineEncryptionKey struct {
    
    // Id AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-artifactstore-encryptionkey.html#cfn-codepipeline-pipeline-artifactstore-encryptionkey-id
    Id string `json:"Id"`
    
    // Type AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codepipeline-pipeline-artifactstore-encryptionkey.html#cfn-codepipeline-pipeline-artifactstore-encryptionkey-type
    Type string `json:"Type"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCodePipelinePipelineEncryptionKey) AWSCloudFormationType() string {
    return "AWS::CodePipeline::Pipeline.EncryptionKey"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCodePipelinePipelineEncryptionKey) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}