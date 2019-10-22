package resources

import "github.com/awslabs/goformation/cloudformation/policies"

// AWSApplicationAutoScalingScalableTarget_SuspendedState AWS CloudFormation Resource (AWS::ApplicationAutoScaling::ScalableTarget.SuspendedState)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalabletarget-suspendedstate.html
type AWSApplicationAutoScalingScalableTarget_SuspendedState struct {

	// DynamicScalingInSuspended AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalabletarget-suspendedstate.html#cfn-applicationautoscaling-scalabletarget-suspendedstate-dynamicscalinginsuspended
	DynamicScalingInSuspended bool `json:"DynamicScalingInSuspended,omitempty"`

	// DynamicScalingOutSuspended AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalabletarget-suspendedstate.html#cfn-applicationautoscaling-scalabletarget-suspendedstate-dynamicscalingoutsuspended
	DynamicScalingOutSuspended bool `json:"DynamicScalingOutSuspended,omitempty"`

	// ScheduledScalingSuspended AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-applicationautoscaling-scalabletarget-suspendedstate.html#cfn-applicationautoscaling-scalabletarget-suspendedstate-scheduledscalingsuspended
	ScheduledScalingSuspended bool `json:"ScheduledScalingSuspended,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy policies.DeletionPolicy

	// _dependsOn stores the logical ID of the resources to be created before this resource
	_dependsOn []string

	// _metadata stores structured data associated with this resource
	_metadata map[string]interface{}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApplicationAutoScalingScalableTarget_SuspendedState) AWSCloudFormationType() string {
	return "AWS::ApplicationAutoScaling::ScalableTarget.SuspendedState"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSApplicationAutoScalingScalableTarget_SuspendedState) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSApplicationAutoScalingScalableTarget_SuspendedState) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// Metadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSApplicationAutoScalingScalableTarget_SuspendedState) Metadata() map[string]interface{} {
	return r._metadata
}

// SetMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSApplicationAutoScalingScalableTarget_SuspendedState) SetMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSApplicationAutoScalingScalableTarget_SuspendedState) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}
