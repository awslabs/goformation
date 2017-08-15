package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EMR::Cluster.Configuration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-configuration.html
type AWSEMRCluster_Configuration struct {

	// Classification AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-configuration.html#cfn-emr-cluster-configuration-classification
	Classification string `json:"Classification"`

	// ConfigurationProperties AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-configuration.html#cfn-emr-cluster-configuration-configurationproperties
	ConfigurationProperties map[string]string `json:"ConfigurationProperties"`

	// Configurations AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-configuration.html#cfn-emr-cluster-configuration-configurations
	Configurations []AWSEMRCluster_Configuration `json:"Configurations"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRCluster_Configuration) AWSCloudFormationType() string {
	return "AWS::EMR::Cluster.Configuration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRCluster_Configuration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEMRCluster_ConfigurationResources retrieves all AWSEMRCluster_Configuration items from a CloudFormation template
func GetAllAWSEMRCluster_Configuration(template *Template) map[string]*AWSEMRCluster_Configuration {

	results := map[string]*AWSEMRCluster_Configuration{}
	for name, resource := range template.Resources {
		result := &AWSEMRCluster_Configuration{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEMRCluster_ConfigurationWithName retrieves all AWSEMRCluster_Configuration items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEMRCluster_Configuration(name string, template *Template) (*AWSEMRCluster_Configuration, error) {

	result := &AWSEMRCluster_Configuration{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEMRCluster_Configuration{}, errors.New("resource not found")

}
