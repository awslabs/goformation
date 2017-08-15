package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EMR::Cluster.BootstrapActionConfig AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-bootstrapactionconfig.html
type AWSEMRCluster_BootstrapActionConfig struct {

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-bootstrapactionconfig.html#cfn-emr-cluster-bootstrapactionconfig-name
	Name string `json:"Name"`

	// ScriptBootstrapAction AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-bootstrapactionconfig.html#cfn-emr-cluster-bootstrapactionconfig-scriptbootstrapaction
	ScriptBootstrapAction AWSEMRCluster_ScriptBootstrapActionConfig `json:"ScriptBootstrapAction"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRCluster_BootstrapActionConfig) AWSCloudFormationType() string {
	return "AWS::EMR::Cluster.BootstrapActionConfig"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRCluster_BootstrapActionConfig) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEMRCluster_BootstrapActionConfigResources retrieves all AWSEMRCluster_BootstrapActionConfig items from a CloudFormation template
func GetAllAWSEMRCluster_BootstrapActionConfig(template *Template) map[string]*AWSEMRCluster_BootstrapActionConfig {

	results := map[string]*AWSEMRCluster_BootstrapActionConfig{}
	for name, resource := range template.Resources {
		result := &AWSEMRCluster_BootstrapActionConfig{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEMRCluster_BootstrapActionConfigWithName retrieves all AWSEMRCluster_BootstrapActionConfig items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEMRCluster_BootstrapActionConfig(name string, template *Template) (*AWSEMRCluster_BootstrapActionConfig, error) {

	result := &AWSEMRCluster_BootstrapActionConfig{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEMRCluster_BootstrapActionConfig{}, errors.New("resource not found")

}
