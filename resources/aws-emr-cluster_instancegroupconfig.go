package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EMR::Cluster.InstanceGroupConfig AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig-instancegroupconfig.html
type AWSEMRCluster_InstanceGroupConfig struct {

	// AutoScalingPolicy AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig-instancegroupconfig.html#cfn-elasticmapreduce-cluster-instancegroupconfig-autoscalingpolicy
	AutoScalingPolicy AWSEMRCluster_AutoScalingPolicy `json:"AutoScalingPolicy"`

	// BidPrice AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig-instancegroupconfig.html#cfn-emr-cluster-jobflowinstancesconfig-instancegroupconfig-bidprice
	BidPrice string `json:"BidPrice"`

	// Configurations AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig-instancegroupconfig.html#cfn-emr-cluster-jobflowinstancesconfig-instancegroupconfig-configurations
	Configurations []AWSEMRCluster_Configuration `json:"Configurations"`

	// EbsConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig-instancegroupconfig.html#cfn-emr-cluster-jobflowinstancesconfig-instancegroupconfigConfigurations-ebsconfiguration
	EbsConfiguration AWSEMRCluster_EbsConfiguration `json:"EbsConfiguration"`

	// InstanceCount AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig-instancegroupconfig.html#cfn-emr-cluster-jobflowinstancesconfig-instancegroupconfig-instancecount
	InstanceCount int64 `json:"InstanceCount"`

	// InstanceType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig-instancegroupconfig.html#cfn-emr-cluster-jobflowinstancesconfig-instancegroupconfig-instancetype
	InstanceType string `json:"InstanceType"`

	// Market AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig-instancegroupconfig.html#cfn-emr-cluster-jobflowinstancesconfig-instancegroupconfig-market
	Market string `json:"Market"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig-instancegroupconfig.html#cfn-emr-cluster-jobflowinstancesconfig-instancegroupconfig-name
	Name string `json:"Name"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRCluster_InstanceGroupConfig) AWSCloudFormationType() string {
	return "AWS::EMR::Cluster.InstanceGroupConfig"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRCluster_InstanceGroupConfig) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEMRCluster_InstanceGroupConfigResources retrieves all AWSEMRCluster_InstanceGroupConfig items from a CloudFormation template
func GetAllAWSEMRCluster_InstanceGroupConfig(template *Template) map[string]*AWSEMRCluster_InstanceGroupConfig {

	results := map[string]*AWSEMRCluster_InstanceGroupConfig{}
	for name, resource := range template.Resources {
		result := &AWSEMRCluster_InstanceGroupConfig{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEMRCluster_InstanceGroupConfigWithName retrieves all AWSEMRCluster_InstanceGroupConfig items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEMRCluster_InstanceGroupConfig(name string, template *Template) (*AWSEMRCluster_InstanceGroupConfig, error) {

	result := &AWSEMRCluster_InstanceGroupConfig{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEMRCluster_InstanceGroupConfig{}, errors.New("resource not found")

}
