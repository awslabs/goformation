package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EMR::Cluster.AutoScalingPolicy AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-autoscalingpolicy.html
type AWSEMRCluster_AutoScalingPolicy struct {

	// Constraints AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-autoscalingpolicy.html#cfn-elasticmapreduce-cluster-autoscalingpolicy-constraints
	Constraints AWSEMRCluster_ScalingConstraints `json:"Constraints"`

	// Rules AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-autoscalingpolicy.html#cfn-elasticmapreduce-cluster-autoscalingpolicy-rules
	Rules []AWSEMRCluster_ScalingRule `json:"Rules"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRCluster_AutoScalingPolicy) AWSCloudFormationType() string {
	return "AWS::EMR::Cluster.AutoScalingPolicy"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRCluster_AutoScalingPolicy) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEMRCluster_AutoScalingPolicyResources retrieves all AWSEMRCluster_AutoScalingPolicy items from a CloudFormation template
func GetAllAWSEMRCluster_AutoScalingPolicy(template *Template) map[string]*AWSEMRCluster_AutoScalingPolicy {

	results := map[string]*AWSEMRCluster_AutoScalingPolicy{}
	for name, resource := range template.Resources {
		result := &AWSEMRCluster_AutoScalingPolicy{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEMRCluster_AutoScalingPolicyWithName retrieves all AWSEMRCluster_AutoScalingPolicy items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEMRCluster_AutoScalingPolicy(name string, template *Template) (*AWSEMRCluster_AutoScalingPolicy, error) {

	result := &AWSEMRCluster_AutoScalingPolicy{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEMRCluster_AutoScalingPolicy{}, errors.New("resource not found")

}
