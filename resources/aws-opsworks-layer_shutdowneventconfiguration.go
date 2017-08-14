package resources

// AWS::OpsWorks::Layer.ShutdownEventConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-lifecycleeventconfiguration-shutdowneventconfiguration.html
type AWSOpsWorksLayerShutdownEventConfiguration struct {

	// DelayUntilElbConnectionsDrained AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-lifecycleeventconfiguration-shutdowneventconfiguration.html#cfn-opsworks-layer-lifecycleconfiguration-shutdowneventconfiguration-delayuntilelbconnectionsdrained
	DelayUntilElbConnectionsDrained bool `json:"DelayUntilElbConnectionsDrained"`

	// ExecutionTimeout AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-lifecycleeventconfiguration-shutdowneventconfiguration.html#cfn-opsworks-layer-lifecycleconfiguration-shutdowneventconfiguration-executiontimeout
	ExecutionTimeout int64 `json:"ExecutionTimeout"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSOpsWorksLayerShutdownEventConfiguration) AWSCloudFormationType() string {
	return "AWS::OpsWorks::Layer.ShutdownEventConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSOpsWorksLayerShutdownEventConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
