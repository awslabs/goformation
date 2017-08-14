package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EMR::InstanceGroupConfig.ScalingRule AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-scalingrule.html
type AWSEMRInstanceGroupConfig_ScalingRule struct {

	// Action AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-scalingrule.html#cfn-elasticmapreduce-instancegroupconfig-scalingrule-action

	Action AWSEMRInstanceGroupConfig_ScalingAction `json:"Action"`

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-scalingrule.html#cfn-elasticmapreduce-instancegroupconfig-scalingrule-description

	Description string `json:"Description"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-scalingrule.html#cfn-elasticmapreduce-instancegroupconfig-scalingrule-name

	Name string `json:"Name"`

	// Trigger AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-scalingrule.html#cfn-elasticmapreduce-instancegroupconfig-scalingrule-trigger

	Trigger AWSEMRInstanceGroupConfig_ScalingTrigger `json:"Trigger"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRInstanceGroupConfig_ScalingRule) AWSCloudFormationType() string {
	return "AWS::EMR::InstanceGroupConfig.ScalingRule"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRInstanceGroupConfig_ScalingRule) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEMRInstanceGroupConfig_ScalingRuleResources retrieves all AWSEMRInstanceGroupConfig_ScalingRule items from a CloudFormation template
func GetAllAWSEMRInstanceGroupConfig_ScalingRule(template *Template) map[string]*AWSEMRInstanceGroupConfig_ScalingRule {

	results := map[string]*AWSEMRInstanceGroupConfig_ScalingRule{}
	for name, resource := range template.Resources {
		result := &AWSEMRInstanceGroupConfig_ScalingRule{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEMRInstanceGroupConfig_ScalingRuleWithName retrieves all AWSEMRInstanceGroupConfig_ScalingRule items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEMRInstanceGroupConfig_ScalingRule(name string, template *Template) (*AWSEMRInstanceGroupConfig_ScalingRule, error) {

	result := &AWSEMRInstanceGroupConfig_ScalingRule{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEMRInstanceGroupConfig_ScalingRule{}, errors.New("resource not found")

}
