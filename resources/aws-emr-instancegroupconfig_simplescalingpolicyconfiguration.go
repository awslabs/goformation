package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EMR::InstanceGroupConfig.SimpleScalingPolicyConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-simplescalingpolicyconfiguration.html
type AWSEMRInstanceGroupConfig_SimpleScalingPolicyConfiguration struct {

	// AdjustmentType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-simplescalingpolicyconfiguration.html#cfn-elasticmapreduce-instancegroupconfig-simplescalingpolicyconfiguration-adjustmenttype
	AdjustmentType string `json:"AdjustmentType"`

	// CoolDown AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-simplescalingpolicyconfiguration.html#cfn-elasticmapreduce-instancegroupconfig-simplescalingpolicyconfiguration-cooldown
	CoolDown int64 `json:"CoolDown"`

	// ScalingAdjustment AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-simplescalingpolicyconfiguration.html#cfn-elasticmapreduce-instancegroupconfig-simplescalingpolicyconfiguration-scalingadjustment
	ScalingAdjustment int64 `json:"ScalingAdjustment"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRInstanceGroupConfig_SimpleScalingPolicyConfiguration) AWSCloudFormationType() string {
	return "AWS::EMR::InstanceGroupConfig.SimpleScalingPolicyConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRInstanceGroupConfig_SimpleScalingPolicyConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEMRInstanceGroupConfig_SimpleScalingPolicyConfigurationResources retrieves all AWSEMRInstanceGroupConfig_SimpleScalingPolicyConfiguration items from a CloudFormation template
func GetAllAWSEMRInstanceGroupConfig_SimpleScalingPolicyConfiguration(template *Template) map[string]*AWSEMRInstanceGroupConfig_SimpleScalingPolicyConfiguration {

	results := map[string]*AWSEMRInstanceGroupConfig_SimpleScalingPolicyConfiguration{}
	for name, resource := range template.Resources {
		result := &AWSEMRInstanceGroupConfig_SimpleScalingPolicyConfiguration{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEMRInstanceGroupConfig_SimpleScalingPolicyConfigurationWithName retrieves all AWSEMRInstanceGroupConfig_SimpleScalingPolicyConfiguration items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEMRInstanceGroupConfig_SimpleScalingPolicyConfiguration(name string, template *Template) (*AWSEMRInstanceGroupConfig_SimpleScalingPolicyConfiguration, error) {

	result := &AWSEMRInstanceGroupConfig_SimpleScalingPolicyConfiguration{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEMRInstanceGroupConfig_SimpleScalingPolicyConfiguration{}, errors.New("resource not found")

}
