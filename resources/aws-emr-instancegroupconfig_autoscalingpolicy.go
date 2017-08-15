package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EMR::InstanceGroupConfig.AutoScalingPolicy AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-autoscalingpolicy.html
type AWSEMRInstanceGroupConfig_AutoScalingPolicy struct {

	// Constraints AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-autoscalingpolicy.html#cfn-elasticmapreduce-instancegroupconfig-autoscalingpolicy-constraints

	Constraints AWSEMRInstanceGroupConfig_ScalingConstraints `json:"Constraints"`

	// Rules AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-autoscalingpolicy.html#cfn-elasticmapreduce-instancegroupconfig-autoscalingpolicy-rules

	Rules []AWSEMRInstanceGroupConfig_ScalingRule `json:"Rules"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRInstanceGroupConfig_AutoScalingPolicy) AWSCloudFormationType() string {
	return "AWS::EMR::InstanceGroupConfig.AutoScalingPolicy"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRInstanceGroupConfig_AutoScalingPolicy) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEMRInstanceGroupConfig_AutoScalingPolicyResources retrieves all AWSEMRInstanceGroupConfig_AutoScalingPolicy items from a CloudFormation template
func GetAllAWSEMRInstanceGroupConfig_AutoScalingPolicy(template *Template) map[string]*AWSEMRInstanceGroupConfig_AutoScalingPolicy {

	results := map[string]*AWSEMRInstanceGroupConfig_AutoScalingPolicy{}
	for name, resource := range template.Resources {
		result := &AWSEMRInstanceGroupConfig_AutoScalingPolicy{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEMRInstanceGroupConfig_AutoScalingPolicyWithName retrieves all AWSEMRInstanceGroupConfig_AutoScalingPolicy items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEMRInstanceGroupConfig_AutoScalingPolicy(name string, template *Template) (*AWSEMRInstanceGroupConfig_AutoScalingPolicy, error) {

	result := &AWSEMRInstanceGroupConfig_AutoScalingPolicy{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEMRInstanceGroupConfig_AutoScalingPolicy{}, errors.New("resource not found")

}
