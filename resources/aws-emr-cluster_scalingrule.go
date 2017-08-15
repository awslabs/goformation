package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EMR::Cluster.ScalingRule AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scalingrule.html
type AWSEMRCluster_ScalingRule struct {

	// Action AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scalingrule.html#cfn-elasticmapreduce-cluster-scalingrule-action
	Action AWSEMRCluster_ScalingAction `json:"Action"`

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scalingrule.html#cfn-elasticmapreduce-cluster-scalingrule-description
	Description string `json:"Description"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scalingrule.html#cfn-elasticmapreduce-cluster-scalingrule-name
	Name string `json:"Name"`

	// Trigger AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scalingrule.html#cfn-elasticmapreduce-cluster-scalingrule-trigger
	Trigger AWSEMRCluster_ScalingTrigger `json:"Trigger"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRCluster_ScalingRule) AWSCloudFormationType() string {
	return "AWS::EMR::Cluster.ScalingRule"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRCluster_ScalingRule) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEMRCluster_ScalingRuleResources retrieves all AWSEMRCluster_ScalingRule items from a CloudFormation template
func GetAllAWSEMRCluster_ScalingRule(template *Template) map[string]*AWSEMRCluster_ScalingRule {

	results := map[string]*AWSEMRCluster_ScalingRule{}
	for name, resource := range template.Resources {
		result := &AWSEMRCluster_ScalingRule{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEMRCluster_ScalingRuleWithName retrieves all AWSEMRCluster_ScalingRule items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEMRCluster_ScalingRule(name string, template *Template) (*AWSEMRCluster_ScalingRule, error) {

	result := &AWSEMRCluster_ScalingRule{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEMRCluster_ScalingRule{}, errors.New("resource not found")

}
