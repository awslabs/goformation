package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EMR::Cluster.EbsConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration.html
type AWSEMRCluster_EbsConfiguration struct {

	// EbsBlockDeviceConfigs AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration.html#cfn-emr-ebsconfiguration-ebsblockdeviceconfigs

	EbsBlockDeviceConfigs []AWSEMRCluster_EbsBlockDeviceConfig `json:"EbsBlockDeviceConfigs"`

	// EbsOptimized AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-ebsconfiguration.html#cfn-emr-ebsconfiguration-ebsoptimized

	EbsOptimized bool `json:"EbsOptimized"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRCluster_EbsConfiguration) AWSCloudFormationType() string {
	return "AWS::EMR::Cluster.EbsConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRCluster_EbsConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEMRCluster_EbsConfigurationResources retrieves all AWSEMRCluster_EbsConfiguration items from a CloudFormation template
func GetAllAWSEMRCluster_EbsConfiguration(template *Template) map[string]*AWSEMRCluster_EbsConfiguration {

	results := map[string]*AWSEMRCluster_EbsConfiguration{}
	for name, resource := range template.Resources {
		result := &AWSEMRCluster_EbsConfiguration{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEMRCluster_EbsConfigurationWithName retrieves all AWSEMRCluster_EbsConfiguration items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEMRCluster_EbsConfiguration(name string, template *Template) (*AWSEMRCluster_EbsConfiguration, error) {

	result := &AWSEMRCluster_EbsConfiguration{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEMRCluster_EbsConfiguration{}, errors.New("resource not found")

}
