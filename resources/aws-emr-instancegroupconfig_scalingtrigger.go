package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EMR::InstanceGroupConfig.ScalingTrigger AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-scalingtrigger.html
type AWSEMRInstanceGroupConfig_ScalingTrigger struct {

	// CloudWatchAlarmDefinition AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-scalingtrigger.html#cfn-elasticmapreduce-instancegroupconfig-scalingtrigger-cloudwatchalarmdefinition
	CloudWatchAlarmDefinition AWSEMRInstanceGroupConfig_CloudWatchAlarmDefinition `json:"CloudWatchAlarmDefinition"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRInstanceGroupConfig_ScalingTrigger) AWSCloudFormationType() string {
	return "AWS::EMR::InstanceGroupConfig.ScalingTrigger"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRInstanceGroupConfig_ScalingTrigger) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEMRInstanceGroupConfig_ScalingTriggerResources retrieves all AWSEMRInstanceGroupConfig_ScalingTrigger items from a CloudFormation template
func GetAllAWSEMRInstanceGroupConfig_ScalingTrigger(template *Template) map[string]*AWSEMRInstanceGroupConfig_ScalingTrigger {

	results := map[string]*AWSEMRInstanceGroupConfig_ScalingTrigger{}
	for name, resource := range template.Resources {
		result := &AWSEMRInstanceGroupConfig_ScalingTrigger{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEMRInstanceGroupConfig_ScalingTriggerWithName retrieves all AWSEMRInstanceGroupConfig_ScalingTrigger items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEMRInstanceGroupConfig_ScalingTrigger(name string, template *Template) (*AWSEMRInstanceGroupConfig_ScalingTrigger, error) {

	result := &AWSEMRInstanceGroupConfig_ScalingTrigger{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEMRInstanceGroupConfig_ScalingTrigger{}, errors.New("resource not found")

}
