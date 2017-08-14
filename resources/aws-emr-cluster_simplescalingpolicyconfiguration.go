package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EMR::Cluster.SimpleScalingPolicyConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-simplescalingpolicyconfiguration.html
type AWSEMRCluster_SimpleScalingPolicyConfiguration struct {

	// AdjustmentType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-simplescalingpolicyconfiguration.html#cfn-elasticmapreduce-cluster-simplescalingpolicyconfiguration-adjustmenttype

	AdjustmentType string `json:"AdjustmentType"`

	// CoolDown AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-simplescalingpolicyconfiguration.html#cfn-elasticmapreduce-cluster-simplescalingpolicyconfiguration-cooldown

	CoolDown int64 `json:"CoolDown"`

	// ScalingAdjustment AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-simplescalingpolicyconfiguration.html#cfn-elasticmapreduce-cluster-simplescalingpolicyconfiguration-scalingadjustment

	ScalingAdjustment int64 `json:"ScalingAdjustment"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRCluster_SimpleScalingPolicyConfiguration) AWSCloudFormationType() string {
	return "AWS::EMR::Cluster.SimpleScalingPolicyConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRCluster_SimpleScalingPolicyConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEMRCluster_SimpleScalingPolicyConfigurationResources retrieves all AWSEMRCluster_SimpleScalingPolicyConfiguration items from a CloudFormation template
func GetAllAWSEMRCluster_SimpleScalingPolicyConfiguration(template *Template) map[string]*AWSEMRCluster_SimpleScalingPolicyConfiguration {

	results := map[string]*AWSEMRCluster_SimpleScalingPolicyConfiguration{}
	for name, resource := range template.Resources {
		result := &AWSEMRCluster_SimpleScalingPolicyConfiguration{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEMRCluster_SimpleScalingPolicyConfigurationWithName retrieves all AWSEMRCluster_SimpleScalingPolicyConfiguration items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEMRCluster_SimpleScalingPolicyConfiguration(name string, template *Template) (*AWSEMRCluster_SimpleScalingPolicyConfiguration, error) {

	result := &AWSEMRCluster_SimpleScalingPolicyConfiguration{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEMRCluster_SimpleScalingPolicyConfiguration{}, errors.New("resource not found")

}
