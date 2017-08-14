package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EMR::Cluster.ScalingTrigger AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scalingtrigger.html
type AWSEMRCluster_ScalingTrigger struct {

	// CloudWatchAlarmDefinition AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scalingtrigger.html#cfn-elasticmapreduce-cluster-scalingtrigger-cloudwatchalarmdefinition

	CloudWatchAlarmDefinition AWSEMRCluster_CloudWatchAlarmDefinition `json:"CloudWatchAlarmDefinition"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRCluster_ScalingTrigger) AWSCloudFormationType() string {
	return "AWS::EMR::Cluster.ScalingTrigger"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRCluster_ScalingTrigger) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEMRCluster_ScalingTriggerResources retrieves all AWSEMRCluster_ScalingTrigger items from a CloudFormation template
func GetAllAWSEMRCluster_ScalingTrigger(template *Template) map[string]*AWSEMRCluster_ScalingTrigger {

	results := map[string]*AWSEMRCluster_ScalingTrigger{}
	for name, resource := range template.Resources {
		result := &AWSEMRCluster_ScalingTrigger{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEMRCluster_ScalingTriggerWithName retrieves all AWSEMRCluster_ScalingTrigger items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEMRCluster_ScalingTrigger(name string, template *Template) (*AWSEMRCluster_ScalingTrigger, error) {

	result := &AWSEMRCluster_ScalingTrigger{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEMRCluster_ScalingTrigger{}, errors.New("resource not found")

}
