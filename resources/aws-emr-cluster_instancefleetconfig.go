package resources

// AWSEMRCluster_InstanceFleetConfig AWS CloudFormation Resource (AWS::EMR::Cluster.InstanceFleetConfig)
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
	TargetOnDemandCapacity int `json:"TargetOnDemandCapacity"`

	// TargetSpotCapacity AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticmapreduce-cluster-instancefleetconfig.html#cfn-elasticmapreduce-cluster-instancefleetconfig-targetspotcapacity
	TargetSpotCapacity int `json:"TargetSpotCapacity"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRCluster_InstanceFleetConfig) AWSCloudFormationType() string {
	return "AWS::EMR::Cluster.InstanceFleetConfig"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRCluster_InstanceFleetConfig) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
