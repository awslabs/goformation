package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EMR::Cluster.ScalingConstraints AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scalingconstraints.html
type AWSEMRCluster_ScalingConstraints struct {

	// MaxCapacity AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scalingconstraints.html#cfn-elasticmapreduce-cluster-scalingconstraints-maxcapacity
	MaxCapacity int64 `json:"MaxCapacity"`

	// MinCapacity AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-scalingconstraints.html#cfn-elasticmapreduce-cluster-scalingconstraints-mincapacity
	MinCapacity int64 `json:"MinCapacity"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRCluster_ScalingConstraints) AWSCloudFormationType() string {
	return "AWS::EMR::Cluster.ScalingConstraints"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRCluster_ScalingConstraints) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEMRCluster_ScalingConstraintsResources retrieves all AWSEMRCluster_ScalingConstraints items from a CloudFormation template
func GetAllAWSEMRCluster_ScalingConstraints(template *Template) map[string]*AWSEMRCluster_ScalingConstraints {

	results := map[string]*AWSEMRCluster_ScalingConstraints{}
	for name, resource := range template.Resources {
		result := &AWSEMRCluster_ScalingConstraints{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEMRCluster_ScalingConstraintsWithName retrieves all AWSEMRCluster_ScalingConstraints items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEMRCluster_ScalingConstraints(name string, template *Template) (*AWSEMRCluster_ScalingConstraints, error) {

	result := &AWSEMRCluster_ScalingConstraints{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEMRCluster_ScalingConstraints{}, errors.New("resource not found")

}
