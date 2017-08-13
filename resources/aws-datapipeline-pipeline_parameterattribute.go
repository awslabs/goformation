package resources

// AWS::DataPipeline::Pipeline.ParameterAttribute AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-parameterobjects-attributes.html
type AWSDataPipelinePipelineParameterAttribute struct {
    
    // Key AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-parameterobjects-attributes.html#cfn-datapipeline-pipeline-parameterobjects-attribtues-key
    Key string `json:"Key"`
    
    // StringValue AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-datapipeline-pipeline-parameterobjects-attributes.html#cfn-datapipeline-pipeline-parameterobjects-attribtues-stringvalue
    StringValue string `json:"StringValue"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDataPipelinePipelineParameterAttribute) AWSCloudFormationType() string {
    return "AWS::DataPipeline::Pipeline.ParameterAttribute"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSDataPipelinePipelineParameterAttribute) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}