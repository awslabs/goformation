package resources

// AWS::EMR::Cluster.ScriptBootstrapActionConfig AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-bootstrapactionconfig-scriptbootstrapactionconfig.html
type AWSEMRClusterScriptBootstrapActionConfig struct {
    
    // Args AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-bootstrapactionconfig-scriptbootstrapactionconfig.html#cfn-emr-cluster-bootstrapactionconfig-scriptbootstrapaction-args
    Args []string `json:"Args"`
    
    // Path AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-bootstrapactionconfig-scriptbootstrapactionconfig.html#cfn-emr-cluster-bootstrapactionconfig-scriptbootstrapaction-path
    Path string `json:"Path"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRClusterScriptBootstrapActionConfig) AWSCloudFormationType() string {
    return "AWS::EMR::Cluster.ScriptBootstrapActionConfig"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRClusterScriptBootstrapActionConfig) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}