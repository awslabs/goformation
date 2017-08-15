package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EMR::Cluster.MetricDimension AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-metricdimension.html
type AWSEMRCluster_MetricDimension struct {

	// Key AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-metricdimension.html#cfn-elasticmapreduce-cluster-metricdimension-key
	Key string `json:"Key"`

	// Value AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-metricdimension.html#cfn-elasticmapreduce-cluster-metricdimension-value
	Value string `json:"Value"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRCluster_MetricDimension) AWSCloudFormationType() string {
	return "AWS::EMR::Cluster.MetricDimension"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRCluster_MetricDimension) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEMRCluster_MetricDimensionResources retrieves all AWSEMRCluster_MetricDimension items from a CloudFormation template
func GetAllAWSEMRCluster_MetricDimension(template *Template) map[string]*AWSEMRCluster_MetricDimension {

	results := map[string]*AWSEMRCluster_MetricDimension{}
	for name, resource := range template.Resources {
		result := &AWSEMRCluster_MetricDimension{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEMRCluster_MetricDimensionWithName retrieves all AWSEMRCluster_MetricDimension items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEMRCluster_MetricDimension(name string, template *Template) (*AWSEMRCluster_MetricDimension, error) {

	result := &AWSEMRCluster_MetricDimension{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEMRCluster_MetricDimension{}, errors.New("resource not found")

}
