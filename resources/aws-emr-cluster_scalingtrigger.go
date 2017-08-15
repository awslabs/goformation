package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EMR::Cluster.ScalingTrigger AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scalingtrigger.html
type AWSEMRCluster_ScalingTrigger struct {

	// CloudWatchAlarmDefinition AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scalingtrigger.html#cfn-elasticmapreduce-cluster-scalingtrigger-cloudwatchalarmdefinition
	CloudWatchAlarmDefinition AWSEMRCluster_CloudWatchAlarmDefinition `json:"CloudWatchAlarmDefinition"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRCluster_ScalingTrigger) AWSCloudFormationType() string {
	return "AWS::EMR::Cluster.ScalingTrigger"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRCluster_ScalingTrigger) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
