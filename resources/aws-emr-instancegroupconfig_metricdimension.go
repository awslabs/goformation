package resources

// AWS::EMR::InstanceGroupConfig.MetricDimension AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-metricdimension.html
type AWSEMRInstanceGroupConfigMetricDimension struct {
    
    // Key AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-metricdimension.html#cfn-elasticmapreduce-instancegroupconfig-metricdimension-key
    Key string `json:"Key"`
    
    // Value AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-metricdimension.html#cfn-elasticmapreduce-instancegroupconfig-metricdimension-value
    Value string `json:"Value"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRInstanceGroupConfigMetricDimension) AWSCloudFormationType() string {
    return "AWS::EMR::InstanceGroupConfig.MetricDimension"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRInstanceGroupConfigMetricDimension) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}