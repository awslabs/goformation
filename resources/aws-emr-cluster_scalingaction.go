package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EMR::Cluster.ScalingAction AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scalingaction.html
type AWSEMRCluster_ScalingAction struct {

	// Market AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scalingaction.html#cfn-elasticmapreduce-cluster-scalingaction-market

	Market string `json:"Market"`

	// SimpleScalingPolicyConfiguration AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scalingaction.html#cfn-elasticmapreduce-cluster-scalingaction-simplescalingpolicyconfiguration

	SimpleScalingPolicyConfiguration AWSEMRCluster_SimpleScalingPolicyConfiguration `json:"SimpleScalingPolicyConfiguration"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRCluster_ScalingAction) AWSCloudFormationType() string {
	return "AWS::EMR::Cluster.ScalingAction"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRCluster_ScalingAction) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEMRCluster_ScalingActionResources retrieves all AWSEMRCluster_ScalingAction items from a CloudFormation template
func GetAllAWSEMRCluster_ScalingAction(template *Template) map[string]*AWSEMRCluster_ScalingAction {

	results := map[string]*AWSEMRCluster_ScalingAction{}
	for name, resource := range template.Resources {
		result := &AWSEMRCluster_ScalingAction{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEMRCluster_ScalingActionWithName retrieves all AWSEMRCluster_ScalingAction items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEMRCluster_ScalingAction(name string, template *Template) (*AWSEMRCluster_ScalingAction, error) {

	result := &AWSEMRCluster_ScalingAction{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEMRCluster_ScalingAction{}, errors.New("resource not found")

}
