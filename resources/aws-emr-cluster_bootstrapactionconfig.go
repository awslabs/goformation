package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EMR::Cluster.BootstrapActionConfig AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-bootstrapactionconfig.html
type AWSEMRCluster_BootstrapActionConfig struct {

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-bootstrapactionconfig.html#cfn-emr-cluster-bootstrapactionconfig-name
	Name string `json:"Name"`

	// ScriptBootstrapAction AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-bootstrapactionconfig.html#cfn-emr-cluster-bootstrapactionconfig-scriptbootstrapaction
	ScriptBootstrapAction AWSEMRCluster_ScriptBootstrapActionConfig `json:"ScriptBootstrapAction"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRCluster_BootstrapActionConfig) AWSCloudFormationType() string {
	return "AWS::EMR::Cluster.BootstrapActionConfig"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRCluster_BootstrapActionConfig) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
