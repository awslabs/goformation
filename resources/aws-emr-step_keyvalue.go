package resources

// AWS::EMR::Step.KeyValue AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-step-keyvalue.html
type AWSEMRStepKeyValue struct {
    
    // Key AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-step-keyvalue.html#cfn-elasticmapreduce-step-keyvalue-key
    Key string `json:"Key"`
    
    // Value AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-step-keyvalue.html#cfn-elasticmapreduce-step-keyvalue-value
    Value string `json:"Value"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRStepKeyValue) AWSCloudFormationType() string {
    return "AWS::EMR::Step.KeyValue"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRStepKeyValue) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}