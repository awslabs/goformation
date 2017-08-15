package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EMR::InstanceGroupConfig.CloudWatchAlarmDefinition AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition.html
type AWSEMRInstanceGroupConfig_CloudWatchAlarmDefinition struct {

	// ComparisonOperator AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition.html#cfn-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition-comparisonoperator
	ComparisonOperator string `json:"ComparisonOperator"`

	// Dimensions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition.html#cfn-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition-dimensions
	Dimensions []AWSEMRInstanceGroupConfig_MetricDimension `json:"Dimensions"`

	// EvaluationPeriods AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition.html#cfn-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition-evaluationperiods
	EvaluationPeriods int64 `json:"EvaluationPeriods"`

	// MetricName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition.html#cfn-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition-metricname
	MetricName string `json:"MetricName"`

	// Namespace AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition.html#cfn-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition-namespace
	Namespace string `json:"Namespace"`

	// Period AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition.html#cfn-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition-period
	Period int64 `json:"Period"`

	// Statistic AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition.html#cfn-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition-statistic
	Statistic string `json:"Statistic"`

	// Threshold AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition.html#cfn-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition-threshold
	Threshold float64 `json:"Threshold"`

	// Unit AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition.html#cfn-elasticmapreduce-instancegroupconfig-cloudwatchalarmdefinition-unit
	Unit string `json:"Unit"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRInstanceGroupConfig_CloudWatchAlarmDefinition) AWSCloudFormationType() string {
	return "AWS::EMR::InstanceGroupConfig.CloudWatchAlarmDefinition"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRInstanceGroupConfig_CloudWatchAlarmDefinition) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEMRInstanceGroupConfig_CloudWatchAlarmDefinitionResources retrieves all AWSEMRInstanceGroupConfig_CloudWatchAlarmDefinition items from a CloudFormation template
func GetAllAWSEMRInstanceGroupConfig_CloudWatchAlarmDefinition(template *Template) map[string]*AWSEMRInstanceGroupConfig_CloudWatchAlarmDefinition {

	results := map[string]*AWSEMRInstanceGroupConfig_CloudWatchAlarmDefinition{}
	for name, resource := range template.Resources {
		result := &AWSEMRInstanceGroupConfig_CloudWatchAlarmDefinition{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEMRInstanceGroupConfig_CloudWatchAlarmDefinitionWithName retrieves all AWSEMRInstanceGroupConfig_CloudWatchAlarmDefinition items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEMRInstanceGroupConfig_CloudWatchAlarmDefinition(name string, template *Template) (*AWSEMRInstanceGroupConfig_CloudWatchAlarmDefinition, error) {

	result := &AWSEMRInstanceGroupConfig_CloudWatchAlarmDefinition{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEMRInstanceGroupConfig_CloudWatchAlarmDefinition{}, errors.New("resource not found")

}
