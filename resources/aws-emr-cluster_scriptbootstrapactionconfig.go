package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EMR::Cluster.ScriptBootstrapActionConfig AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-bootstrapactionconfig-scriptbootstrapactionconfig.html
type AWSEMRCluster_ScriptBootstrapActionConfig struct {

	// Args AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-bootstrapactionconfig-scriptbootstrapactionconfig.html#cfn-emr-cluster-bootstrapactionconfig-scriptbootstrapaction-args
	Args []string `json:"Args"`

	// Path AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-bootstrapactionconfig-scriptbootstrapactionconfig.html#cfn-emr-cluster-bootstrapactionconfig-scriptbootstrapaction-path
	Path string `json:"Path"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRCluster_ScriptBootstrapActionConfig) AWSCloudFormationType() string {
	return "AWS::EMR::Cluster.ScriptBootstrapActionConfig"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRCluster_ScriptBootstrapActionConfig) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEMRCluster_ScriptBootstrapActionConfigResources retrieves all AWSEMRCluster_ScriptBootstrapActionConfig items from a CloudFormation template
func GetAllAWSEMRCluster_ScriptBootstrapActionConfig(template *Template) map[string]*AWSEMRCluster_ScriptBootstrapActionConfig {

	results := map[string]*AWSEMRCluster_ScriptBootstrapActionConfig{}
	for name, resource := range template.Resources {
		result := &AWSEMRCluster_ScriptBootstrapActionConfig{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEMRCluster_ScriptBootstrapActionConfigWithName retrieves all AWSEMRCluster_ScriptBootstrapActionConfig items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEMRCluster_ScriptBootstrapActionConfig(name string, template *Template) (*AWSEMRCluster_ScriptBootstrapActionConfig, error) {

	result := &AWSEMRCluster_ScriptBootstrapActionConfig{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEMRCluster_ScriptBootstrapActionConfig{}, errors.New("resource not found")

}
