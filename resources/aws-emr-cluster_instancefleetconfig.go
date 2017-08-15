package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EMR::Cluster.InstanceFleetConfig AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-instancefleetconfig.html
type AWSEMRCluster_InstanceFleetConfig struct {

	// InstanceTypeConfigs AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-instancefleetconfig.html#cfn-elasticmapreduce-cluster-instancefleetconfig-instancetypeconfigs
	InstanceTypeConfigs []AWSEMRCluster_InstanceTypeConfig `json:"InstanceTypeConfigs"`

	// LaunchSpecifications AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-instancefleetconfig.html#cfn-elasticmapreduce-cluster-instancefleetconfig-launchspecifications
	LaunchSpecifications AWSEMRCluster_InstanceFleetProvisioningSpecifications `json:"LaunchSpecifications"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-instancefleetconfig.html#cfn-elasticmapreduce-cluster-instancefleetconfig-name
	Name string `json:"Name"`

	// TargetOnDemandCapacity AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-instancefleetconfig.html#cfn-elasticmapreduce-cluster-instancefleetconfig-targetondemandcapacity
	TargetOnDemandCapacity int64 `json:"TargetOnDemandCapacity"`

	// TargetSpotCapacity AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-instancefleetconfig.html#cfn-elasticmapreduce-cluster-instancefleetconfig-targetspotcapacity
	TargetSpotCapacity int64 `json:"TargetSpotCapacity"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRCluster_InstanceFleetConfig) AWSCloudFormationType() string {
	return "AWS::EMR::Cluster.InstanceFleetConfig"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRCluster_InstanceFleetConfig) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEMRCluster_InstanceFleetConfigResources retrieves all AWSEMRCluster_InstanceFleetConfig items from a CloudFormation template
func GetAllAWSEMRCluster_InstanceFleetConfig(template *Template) map[string]*AWSEMRCluster_InstanceFleetConfig {

	results := map[string]*AWSEMRCluster_InstanceFleetConfig{}
	for name, resource := range template.Resources {
		result := &AWSEMRCluster_InstanceFleetConfig{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEMRCluster_InstanceFleetConfigWithName retrieves all AWSEMRCluster_InstanceFleetConfig items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEMRCluster_InstanceFleetConfig(name string, template *Template) (*AWSEMRCluster_InstanceFleetConfig, error) {

	result := &AWSEMRCluster_InstanceFleetConfig{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEMRCluster_InstanceFleetConfig{}, errors.New("resource not found")

}
