package resources

// AWS::ApplicationAutoScaling::ScalingPolicy.PredefinedMetricSpecification AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predefinedmetricspecification.html
type AWSApplicationAutoScalingScalingPolicyPredefinedMetricSpecification struct {
    
    // PredefinedMetricType AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predefinedmetricspecification.html#cfn-applicationautoscaling-scalingpolicy-predefinedmetricspecification-predefinedmetrictype
    PredefinedMetricType string `json:"PredefinedMetricType"`
    
    // ResourceLabel AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalingpolicy-predefinedmetricspecification.html#cfn-applicationautoscaling-scalingpolicy-predefinedmetricspecification-resourcelabel
    ResourceLabel string `json:"ResourceLabel"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApplicationAutoScalingScalingPolicyPredefinedMetricSpecification) AWSCloudFormationType() string {
    return "AWS::ApplicationAutoScaling::ScalingPolicy.PredefinedMetricSpecification"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSApplicationAutoScalingScalingPolicyPredefinedMetricSpecification) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}