package resources

import "github.com/awslabs/goformation/cloudformation/policies"

// AWSEC2TrafficMirrorFilterRule_TrafficMirrorPortRange AWS CloudFormation Resource (AWS::EC2::TrafficMirrorFilterRule.TrafficMirrorPortRange)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-trafficmirrorfilterrule-trafficmirrorportrange.html
type AWSEC2TrafficMirrorFilterRule_TrafficMirrorPortRange struct {

	// FromPort AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-trafficmirrorfilterrule-trafficmirrorportrange.html#cfn-ec2-trafficmirrorfilterrule-trafficmirrorportrange-fromport
	FromPort int `json:"FromPort"`

	// ToPort AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-trafficmirrorfilterrule-trafficmirrorportrange.html#cfn-ec2-trafficmirrorfilterrule-trafficmirrorportrange-toport
	ToPort int `json:"ToPort"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy policies.DeletionPolicy

	// _dependsOn stores the logical ID of the resources to be created before this resource
	_dependsOn []string

	// _metadata stores structured data associated with this resource
	_metadata map[string]interface{}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2TrafficMirrorFilterRule_TrafficMirrorPortRange) AWSCloudFormationType() string {
	return "AWS::EC2::TrafficMirrorFilterRule.TrafficMirrorPortRange"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSEC2TrafficMirrorFilterRule_TrafficMirrorPortRange) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSEC2TrafficMirrorFilterRule_TrafficMirrorPortRange) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// Metadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSEC2TrafficMirrorFilterRule_TrafficMirrorPortRange) Metadata() map[string]interface{} {
	return r._metadata
}

// SetMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSEC2TrafficMirrorFilterRule_TrafficMirrorPortRange) SetMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSEC2TrafficMirrorFilterRule_TrafficMirrorPortRange) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}
