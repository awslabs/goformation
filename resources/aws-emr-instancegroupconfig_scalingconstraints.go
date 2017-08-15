package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EMR::InstanceGroupConfig.ScalingConstraints AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-scalingconstraints.html
type AWSEMRInstanceGroupConfig_ScalingConstraints struct {

	// MaxCapacity AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-scalingconstraints.html#cfn-elasticmapreduce-instancegroupconfig-scalingconstraints-maxcapacity
	MaxCapacity int64 `json:"MaxCapacity"`

	// MinCapacity AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-instancegroupconfig-scalingconstraints.html#cfn-elasticmapreduce-instancegroupconfig-scalingconstraints-mincapacity
	MinCapacity int64 `json:"MinCapacity"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRInstanceGroupConfig_ScalingConstraints) AWSCloudFormationType() string {
	return "AWS::EMR::InstanceGroupConfig.ScalingConstraints"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRInstanceGroupConfig_ScalingConstraints) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEMRInstanceGroupConfig_ScalingConstraintsResources retrieves all AWSEMRInstanceGroupConfig_ScalingConstraints items from a CloudFormation template
func GetAllAWSEMRInstanceGroupConfig_ScalingConstraints(template *Template) map[string]*AWSEMRInstanceGroupConfig_ScalingConstraints {

	results := map[string]*AWSEMRInstanceGroupConfig_ScalingConstraints{}
	for name, resource := range template.Resources {
		result := &AWSEMRInstanceGroupConfig_ScalingConstraints{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEMRInstanceGroupConfig_ScalingConstraintsWithName retrieves all AWSEMRInstanceGroupConfig_ScalingConstraints items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEMRInstanceGroupConfig_ScalingConstraints(name string, template *Template) (*AWSEMRInstanceGroupConfig_ScalingConstraints, error) {

	result := &AWSEMRInstanceGroupConfig_ScalingConstraints{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEMRInstanceGroupConfig_ScalingConstraints{}, errors.New("resource not found")

}
