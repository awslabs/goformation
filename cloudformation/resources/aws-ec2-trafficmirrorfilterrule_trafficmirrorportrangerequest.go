package resources

import "github.com/awslabs/goformation/cloudformation/policies"

// AWSEC2TrafficMirrorFilterRule_TrafficMirrorPortRangeRequest AWS CloudFormation Resource (AWS::EC2::TrafficMirrorFilterRule.TrafficMirrorPortRangeRequest)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-trafficmirrorfilterrule-trafficmirrorportrangerequest.html
type AWSEC2TrafficMirrorFilterRule_TrafficMirrorPortRangeRequest struct {

	// FromPort AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-trafficmirrorfilterrule-trafficmirrorportrangerequest.html#cfn-ec2-trafficmirrorfilterrule-trafficmirrorportrangerequest-fromport
	FromPort int `json:"FromPort"`

	// ToPort AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-trafficmirrorfilterrule-trafficmirrorportrangerequest.html#cfn-ec2-trafficmirrorfilterrule-trafficmirrorportrangerequest-toport
	ToPort int `json:"ToPort"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy policies.DeletionPolicy

	// _dependsOn stores the logical ID of the resources to be created before this resource
	_dependsOn []string

	// _metadata stores structured data associated with this resource
	_metadata map[string]interface{}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2TrafficMirrorFilterRule_TrafficMirrorPortRangeRequest) AWSCloudFormationType() string {
	return "AWS::EC2::TrafficMirrorFilterRule.TrafficMirrorPortRangeRequest"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSEC2TrafficMirrorFilterRule_TrafficMirrorPortRangeRequest) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSEC2TrafficMirrorFilterRule_TrafficMirrorPortRangeRequest) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// Metadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSEC2TrafficMirrorFilterRule_TrafficMirrorPortRangeRequest) Metadata() map[string]interface{} {
	return r._metadata
}

// SetMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSEC2TrafficMirrorFilterRule_TrafficMirrorPortRangeRequest) SetMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// DeletionPolicy returns the AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSEC2TrafficMirrorFilterRule_TrafficMirrorPortRangeRequest) DeletionPolicy() policies.DeletionPolicy {
	return r._deletionPolicy
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSEC2TrafficMirrorFilterRule_TrafficMirrorPortRangeRequest) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}
