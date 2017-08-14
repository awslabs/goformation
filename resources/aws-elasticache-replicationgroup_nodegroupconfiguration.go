package resources

// AWS::ElastiCache::ReplicationGroup.NodeGroupConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-replicationgroup-nodegroupconfiguration.html
type AWSElastiCacheReplicationGroupNodeGroupConfiguration struct {

	// PrimaryAvailabilityZone AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-replicationgroup-nodegroupconfiguration.html#cfn-elasticache-replicationgroup-nodegroupconfiguration-primaryavailabilityzone
	PrimaryAvailabilityZone string `json:"PrimaryAvailabilityZone"`

	// ReplicaAvailabilityZones AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-replicationgroup-nodegroupconfiguration.html#cfn-elasticache-replicationgroup-nodegroupconfiguration-replicaavailabilityzones
	ReplicaAvailabilityZones []AWSElastiCacheReplicationGroupNodeGroupConfigurationstring `json:"ReplicaAvailabilityZones"`

	// ReplicaCount AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-replicationgroup-nodegroupconfiguration.html#cfn-elasticache-replicationgroup-nodegroupconfiguration-replicacount
	ReplicaCount int64 `json:"ReplicaCount"`

	// Slots AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-replicationgroup-nodegroupconfiguration.html#cfn-elasticache-replicationgroup-nodegroupconfiguration-slots
	Slots string `json:"Slots"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElastiCacheReplicationGroupNodeGroupConfiguration) AWSCloudFormationType() string {
	return "AWS::ElastiCache::ReplicationGroup.NodeGroupConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElastiCacheReplicationGroupNodeGroupConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
