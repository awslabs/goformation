package resources

// AWS::OpsWorks::Layer.LoadBasedAutoScaling AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-loadbasedautoscaling.html
type AWSOpsWorksLayerLoadBasedAutoScaling struct {
    
    // DownScaling AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-loadbasedautoscaling.html#cfn-opsworks-layer-loadbasedautoscaling-downscaling
    DownScaling AWSOpsWorksLayerLoadBasedAutoScalingAutoScalingThresholds `json:"DownScaling"`
    
    // Enable AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-loadbasedautoscaling.html#cfn-opsworks-layer-loadbasedautoscaling-enable
    Enable bool `json:"Enable"`
    
    // UpScaling AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-loadbasedautoscaling.html#cfn-opsworks-layer-loadbasedautoscaling-upscaling
    UpScaling AWSOpsWorksLayerLoadBasedAutoScalingAutoScalingThresholds `json:"UpScaling"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSOpsWorksLayerLoadBasedAutoScaling) AWSCloudFormationType() string {
    return "AWS::OpsWorks::Layer.LoadBasedAutoScaling"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSOpsWorksLayerLoadBasedAutoScaling) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}