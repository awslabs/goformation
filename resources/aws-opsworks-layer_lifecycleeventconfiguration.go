package resources

// AWS::OpsWorks::Layer.LifecycleEventConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-lifecycleeventconfiguration.html
type AWSOpsWorksLayerLifecycleEventConfiguration struct {
    
    // ShutdownEventConfiguration AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-lifecycleeventconfiguration.html#cfn-opsworks-layer-lifecycleconfiguration-shutdowneventconfiguration
    ShutdownEventConfiguration AWSOpsWorksLayerLifecycleEventConfigurationShutdownEventConfiguration `json:"ShutdownEventConfiguration"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSOpsWorksLayerLifecycleEventConfiguration) AWSCloudFormationType() string {
    return "AWS::OpsWorks::Layer.LifecycleEventConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSOpsWorksLayerLifecycleEventConfiguration) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}