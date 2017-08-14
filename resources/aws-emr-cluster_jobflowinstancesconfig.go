package resources

// AWS::EMR::Cluster.JobFlowInstancesConfig AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html
type AWSEMRClusterJobFlowInstancesConfig struct {

	// AdditionalMasterSecurityGroups AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-emr-cluster-jobflowinstancesconfig-additionalmastersecuritygroups
	AdditionalMasterSecurityGroups []AWSEMRClusterJobFlowInstancesConfigstring `json:"AdditionalMasterSecurityGroups"`

	// AdditionalSlaveSecurityGroups AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-emr-cluster-jobflowinstancesconfig-additionalslavesecuritygroups
	AdditionalSlaveSecurityGroups []AWSEMRClusterJobFlowInstancesConfigstring `json:"AdditionalSlaveSecurityGroups"`

	// CoreInstanceFleet AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-elasticmapreduce-cluster-jobflowinstancesconfig-coreinstancefleet
	CoreInstanceFleet AWSEMRClusterJobFlowInstancesConfigInstanceFleetConfig `json:"CoreInstanceFleet"`

	// CoreInstanceGroup AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-emr-cluster-jobflowinstancesconfig-coreinstancegroup
	CoreInstanceGroup AWSEMRClusterJobFlowInstancesConfigInstanceGroupConfig `json:"CoreInstanceGroup"`

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
	MasterInstanceFleet AWSEMRClusterJobFlowInstancesConfigInstanceFleetConfig `json:"MasterInstanceFleet"`

	// MasterInstanceGroup AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-emr-cluster-jobflowinstancesconfig-coreinstancegroup
	MasterInstanceGroup AWSEMRClusterJobFlowInstancesConfigInstanceGroupConfig `json:"MasterInstanceGroup"`

	// Placement AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-emr-cluster-jobflowinstancesconfig.html#cfn-emr-cluster-jobflowinstancesconfig-placement
	Placement AWSEMRClusterJobFlowInstancesConfigPlacementType `json:"Placement"`

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
func (r *AWSEMRClusterJobFlowInstancesConfig) AWSCloudFormationType() string {
	return "AWS::EMR::Cluster.JobFlowInstancesConfig"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRClusterJobFlowInstancesConfig) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
