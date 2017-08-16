package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSEMRInstanceFleetConfig AWS CloudFormation Resource (AWS::EMR::InstanceFleetConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-instancefleetconfig.html
type AWSEMRInstanceFleetConfig struct {

	// ClusterId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-instancefleetconfig.html#cfn-elasticmapreduce-instancefleetconfig-clusterid
	ClusterId string `json:"ClusterId"`

	// InstanceFleetType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-instancefleetconfig.html#cfn-elasticmapreduce-instancefleetconfig-instancefleettype
	InstanceFleetType string `json:"InstanceFleetType"`

	// InstanceTypeConfigs AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-instancefleetconfig.html#cfn-elasticmapreduce-instancefleetconfig-instancetypeconfigs
	InstanceTypeConfigs []AWSEMRInstanceFleetConfig_InstanceTypeConfig `json:"InstanceTypeConfigs"`

	// LaunchSpecifications AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-instancefleetconfig.html#cfn-elasticmapreduce-instancefleetconfig-launchspecifications
	LaunchSpecifications AWSEMRInstanceFleetConfig_InstanceFleetProvisioningSpecifications `json:"LaunchSpecifications"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-instancefleetconfig.html#cfn-elasticmapreduce-instancefleetconfig-name
	Name string `json:"Name"`

	// TargetOnDemandCapacity AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-instancefleetconfig.html#cfn-elasticmapreduce-instancefleetconfig-targetondemandcapacity
	TargetOnDemandCapacity int `json:"TargetOnDemandCapacity"`

	// TargetSpotCapacity AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-instancefleetconfig.html#cfn-elasticmapreduce-instancefleetconfig-targetspotcapacity
	TargetSpotCapacity int `json:"TargetSpotCapacity"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRInstanceFleetConfig) AWSCloudFormationType() string {
	return "AWS::EMR::InstanceFleetConfig"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRInstanceFleetConfig) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEMRInstanceFleetConfigResources retrieves all AWSEMRInstanceFleetConfig items from a CloudFormation template
func (t *Template) GetAllAWSEMRInstanceFleetConfigResources() map[string]*AWSEMRInstanceFleetConfig {

	results := map[string]*AWSEMRInstanceFleetConfig{}
	for name, resource := range t.Resources {
		result := &AWSEMRInstanceFleetConfig{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEMRInstanceFleetConfigWithName retrieves all AWSEMRInstanceFleetConfig items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSEMRInstanceFleetConfigWithName(name string) (*AWSEMRInstanceFleetConfig, error) {

	result := &AWSEMRInstanceFleetConfig{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEMRInstanceFleetConfig{}, errors.New("resource not found")

}
