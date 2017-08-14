package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EMR::InstanceGroupConfig.ScalingAction AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-scalingaction.html
type AWSEMRInstanceGroupConfig_ScalingAction struct {

	// Market AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-scalingaction.html#cfn-elasticmapreduce-instancegroupconfig-scalingaction-market

	Market string `json:"Market"`

	// SimpleScalingPolicyConfiguration AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-scalingaction.html#cfn-elasticmapreduce-instancegroupconfig-scalingaction-simplescalingpolicyconfiguration

	SimpleScalingPolicyConfiguration AWSEMRInstanceGroupConfig_SimpleScalingPolicyConfiguration `json:"SimpleScalingPolicyConfiguration"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRInstanceGroupConfig_ScalingAction) AWSCloudFormationType() string {
	return "AWS::EMR::InstanceGroupConfig.ScalingAction"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRInstanceGroupConfig_ScalingAction) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEMRInstanceGroupConfig_ScalingActionResources retrieves all AWSEMRInstanceGroupConfig_ScalingAction items from a CloudFormation template
func GetAllAWSEMRInstanceGroupConfig_ScalingAction(template *Template) map[string]*AWSEMRInstanceGroupConfig_ScalingAction {

	results := map[string]*AWSEMRInstanceGroupConfig_ScalingAction{}
	for name, resource := range template.Resources {
		result := &AWSEMRInstanceGroupConfig_ScalingAction{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEMRInstanceGroupConfig_ScalingActionWithName retrieves all AWSEMRInstanceGroupConfig_ScalingAction items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEMRInstanceGroupConfig_ScalingAction(name string, template *Template) (*AWSEMRInstanceGroupConfig_ScalingAction, error) {

	result := &AWSEMRInstanceGroupConfig_ScalingAction{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEMRInstanceGroupConfig_ScalingAction{}, errors.New("resource not found")

}
