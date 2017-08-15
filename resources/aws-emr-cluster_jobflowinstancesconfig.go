package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EMR::Cluster.JobFlowInstancesConfig AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html
type AWSEMRCluster_JobFlowInstancesConfig struct {

	// AdditionalMasterSecurityGroups AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-emr-cluster-jobflowinstancesconfig-additionalmastersecuritygroups
	AdditionalMasterSecurityGroups []string `json:"AdditionalMasterSecurityGroups"`

	// AdditionalSlaveSecurityGroups AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-emr-cluster-jobflowinstancesconfig-additionalslavesecuritygroups
	AdditionalSlaveSecurityGroups []string `json:"AdditionalSlaveSecurityGroups"`

	// CoreInstanceFleet AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-elasticmapreduce-cluster-jobflowinstancesconfig-coreinstancefleet
	CoreInstanceFleet AWSEMRCluster_InstanceFleetConfig `json:"CoreInstanceFleet"`

	// CoreInstanceGroup AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-emr-cluster-jobflowinstancesconfig-coreinstancegroup
	CoreInstanceGroup AWSEMRCluster_InstanceGroupConfig `json:"CoreInstanceGroup"`

	// Ec2KeyName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-emr-cluster-jobflowinstancesconfig-ec2keyname
	Ec2KeyName string `json:"Ec2KeyName"`

	// Ec2SubnetId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-emr-cluster-jobflowinstancesconfig-ec2subnetid
	Ec2SubnetId string `json:"Ec2SubnetId"`

	// EmrManagedMasterSecurityGroup AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-emr-cluster-jobflowinstancesconfig-emrmanagedmastersecuritygroup
	EmrManagedMasterSecurityGroup string `json:"EmrManagedMasterSecurityGroup"`

	// EmrManagedSlaveSecurityGroup AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-emr-cluster-jobflowinstancesconfig-emrmanagedslavesecuritygroup
	EmrManagedSlaveSecurityGroup string `json:"EmrManagedSlaveSecurityGroup"`

	// HadoopVersion AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-emr-cluster-jobflowinstancesconfig-hadoopversion
	HadoopVersion string `json:"HadoopVersion"`

	// MasterInstanceFleet AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-elasticmapreduce-cluster-jobflowinstancesconfig-masterinstancefleet
	MasterInstanceFleet AWSEMRCluster_InstanceFleetConfig `json:"MasterInstanceFleet"`

	// MasterInstanceGroup AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-emr-cluster-jobflowinstancesconfig-coreinstancegroup
	MasterInstanceGroup AWSEMRCluster_InstanceGroupConfig `json:"MasterInstanceGroup"`

	// Placement AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-emr-cluster-jobflowinstancesconfig-placement
	Placement AWSEMRCluster_PlacementType `json:"Placement"`

	// ServiceAccessSecurityGroup AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-emr-cluster-jobflowinstancesconfig-serviceaccesssecuritygroup
	ServiceAccessSecurityGroup string `json:"ServiceAccessSecurityGroup"`

	// TerminationProtected AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-emr-cluster-jobflowinstancesconfig-terminationprotected
	TerminationProtected bool `json:"TerminationProtected"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRCluster_JobFlowInstancesConfig) AWSCloudFormationType() string {
	return "AWS::EMR::Cluster.JobFlowInstancesConfig"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRCluster_JobFlowInstancesConfig) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEMRCluster_JobFlowInstancesConfigResources retrieves all AWSEMRCluster_JobFlowInstancesConfig items from a CloudFormation template
func GetAllAWSEMRCluster_JobFlowInstancesConfig(template *Template) map[string]*AWSEMRCluster_JobFlowInstancesConfig {

	results := map[string]*AWSEMRCluster_JobFlowInstancesConfig{}
	for name, resource := range template.Resources {
		result := &AWSEMRCluster_JobFlowInstancesConfig{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEMRCluster_JobFlowInstancesConfigWithName retrieves all AWSEMRCluster_JobFlowInstancesConfig items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEMRCluster_JobFlowInstancesConfig(name string, template *Template) (*AWSEMRCluster_JobFlowInstancesConfig, error) {

	result := &AWSEMRCluster_JobFlowInstancesConfig{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEMRCluster_JobFlowInstancesConfig{}, errors.New("resource not found")

}
