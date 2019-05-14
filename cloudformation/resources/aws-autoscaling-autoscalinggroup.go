package resources

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/vrealzhou/goformation/cloudformation/policies"
)

// AWSAutoScalingAutoScalingGroup AWS CloudFormation Resource (AWS::AutoScaling::AutoScalingGroup)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html
type AWSAutoScalingAutoScalingGroup struct {

	// AutoScalingGroupName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-autoscaling-autoscalinggroup-autoscalinggroupname
	AutoScalingGroupName string `json:"AutoScalingGroupName,omitempty"`

	// AvailabilityZones AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-availabilityzones
	AvailabilityZones []string `json:"AvailabilityZones,omitempty"`

	// Cooldown AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-cooldown
	Cooldown string `json:"Cooldown,omitempty"`

	// DesiredCapacity AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-desiredcapacity
	DesiredCapacity string `json:"DesiredCapacity,omitempty"`

	// HealthCheckGracePeriod AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-healthcheckgraceperiod
	HealthCheckGracePeriod int `json:"HealthCheckGracePeriod,omitempty"`

	// HealthCheckType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-healthchecktype
	HealthCheckType string `json:"HealthCheckType,omitempty"`

	// InstanceId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-instanceid
	InstanceId string `json:"InstanceId,omitempty"`

	// LaunchConfigurationName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-launchconfigurationname
	LaunchConfigurationName string `json:"LaunchConfigurationName,omitempty"`

	// LaunchTemplate AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-launchtemplate
	LaunchTemplate *AWSAutoScalingAutoScalingGroup_LaunchTemplateSpecification `json:"LaunchTemplate,omitempty"`

	// LifecycleHookSpecificationList AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-autoscaling-autoscalinggroup-lifecyclehookspecificationlist
	LifecycleHookSpecificationList []AWSAutoScalingAutoScalingGroup_LifecycleHookSpecification `json:"LifecycleHookSpecificationList,omitempty"`

	// LoadBalancerNames AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-loadbalancernames
	LoadBalancerNames []string `json:"LoadBalancerNames,omitempty"`

	// MaxSize AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-maxsize
	MaxSize string `json:"MaxSize,omitempty"`

	// MetricsCollection AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-metricscollection
	MetricsCollection []AWSAutoScalingAutoScalingGroup_MetricsCollection `json:"MetricsCollection,omitempty"`

	// MinSize AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-minsize
	MinSize string `json:"MinSize,omitempty"`

	// MixedInstancesPolicy AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-mixedinstancespolicy
	MixedInstancesPolicy *AWSAutoScalingAutoScalingGroup_MixedInstancesPolicy `json:"MixedInstancesPolicy,omitempty"`

	// NotificationConfigurations AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-notificationconfigurations
	NotificationConfigurations []AWSAutoScalingAutoScalingGroup_NotificationConfiguration `json:"NotificationConfigurations,omitempty"`

	// PlacementGroup AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-placementgroup
	PlacementGroup string `json:"PlacementGroup,omitempty"`

	// ServiceLinkedRoleARN AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-autoscaling-autoscalinggroup-servicelinkedrolearn
	ServiceLinkedRoleARN string `json:"ServiceLinkedRoleARN,omitempty"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-tags
	Tags []AWSAutoScalingAutoScalingGroup_TagProperty `json:"Tags,omitempty"`

	// TargetGroupARNs AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-targetgrouparns
	TargetGroupARNs []string `json:"TargetGroupARNs,omitempty"`

	// TerminationPolicies AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-termpolicy
	TerminationPolicies []string `json:"TerminationPolicies,omitempty"`

	// VPCZoneIdentifier AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-as-group.html#cfn-as-group-vpczoneidentifier
	VPCZoneIdentifier []string `json:"VPCZoneIdentifier,omitempty"`

	// _updatePolicy represents a CloudFormation UpdatePolicy
	_updatePolicy *policies.UpdatePolicy

	// _creationPolicy represents a CloudFormation CreationPolicy
	_creationPolicy *policies.CreationPolicy

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy policies.DeletionPolicy

	// _dependsOn stores the logical ID of the resources to be created before this resource
	_dependsOn []string

	// _metadata stores structured data associated with this resource
	_metadata map[string]interface{}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSAutoScalingAutoScalingGroup) AWSCloudFormationType() string {
	return "AWS::AutoScaling::AutoScalingGroup"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSAutoScalingAutoScalingGroup) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSAutoScalingAutoScalingGroup) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// Metadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSAutoScalingAutoScalingGroup) Metadata() map[string]interface{} {
	return r._metadata
}

// SetMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSAutoScalingAutoScalingGroup) SetMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSAutoScalingAutoScalingGroup) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}

// SetUpdatePolicy applies an AWS CloudFormation UpdatePolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-updatepolicy.html
func (r *AWSAutoScalingAutoScalingGroup) SetUpdatePolicy(policy *policies.UpdatePolicy) {
	r._updatePolicy = policy
}

// SetCreationPolicy applies an AWS CloudFormation CreationPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-creationpolicy.html
func (r *AWSAutoScalingAutoScalingGroup) SetCreationPolicy(policy *policies.CreationPolicy) {
	r._creationPolicy = policy
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r AWSAutoScalingAutoScalingGroup) MarshalJSON() ([]byte, error) {
	type Properties AWSAutoScalingAutoScalingGroup
	return json.Marshal(&struct {
		Type           string
		Properties     Properties
		DependsOn      []string                 `json:"DependsOn,omitempty"`
		Metadata       map[string]interface{}   `json:"Metadata,omitempty"`
		DeletionPolicy policies.DeletionPolicy  `json:"DeletionPolicy,omitempty"`
		UpdatePolicy   *policies.UpdatePolicy   `json:"UpdatePolicy,omitempty"`
		CreationPolicy *policies.CreationPolicy `json:"CreationPolicy,omitempty"`
	}{
		Type:           r.AWSCloudFormationType(),
		Properties:     (Properties)(r),
		DependsOn:      r._dependsOn,
		Metadata:       r._metadata,
		DeletionPolicy: r._deletionPolicy,
		UpdatePolicy:   r._updatePolicy,
		CreationPolicy: r._creationPolicy,
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *AWSAutoScalingAutoScalingGroup) UnmarshalJSON(b []byte) error {
	type Properties AWSAutoScalingAutoScalingGroup
	res := &struct {
		Type       string
		Properties *Properties
		DependsOn  []string
		Metadata   map[string]interface{}
	}{}

	dec := json.NewDecoder(bytes.NewReader(b))
	dec.DisallowUnknownFields() // Force error if unknown field is found

	if err := dec.Decode(&res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}

	// If the resource has no Properties set, it could be nil
	if res.Properties != nil {
		*r = AWSAutoScalingAutoScalingGroup(*res.Properties)
	}
	if res.DependsOn != nil {
		r._dependsOn = res.DependsOn
	}
	if res.Metadata != nil {
		r._metadata = res.Metadata
	}

	return nil
}
