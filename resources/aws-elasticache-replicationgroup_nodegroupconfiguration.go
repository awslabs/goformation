package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ElastiCache::ReplicationGroup.NodeGroupConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-replicationgroup-nodegroupconfiguration.html
type AWSElastiCacheReplicationGroup_NodeGroupConfiguration struct {

	// PrimaryAvailabilityZone AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-replicationgroup-nodegroupconfiguration.html#cfn-elasticache-replicationgroup-nodegroupconfiguration-primaryavailabilityzone
	PrimaryAvailabilityZone string `json:"PrimaryAvailabilityZone"`

	// ReplicaAvailabilityZones AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-replicationgroup-nodegroupconfiguration.html#cfn-elasticache-replicationgroup-nodegroupconfiguration-replicaavailabilityzones
	ReplicaAvailabilityZones []string `json:"ReplicaAvailabilityZones"`

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
func (r *AWSElastiCacheReplicationGroup_NodeGroupConfiguration) AWSCloudFormationType() string {
	return "AWS::ElastiCache::ReplicationGroup.NodeGroupConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElastiCacheReplicationGroup_NodeGroupConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSElastiCacheReplicationGroup_NodeGroupConfigurationResources retrieves all AWSElastiCacheReplicationGroup_NodeGroupConfiguration items from a CloudFormation template
func GetAllAWSElastiCacheReplicationGroup_NodeGroupConfiguration(template *Template) map[string]*AWSElastiCacheReplicationGroup_NodeGroupConfiguration {

	results := map[string]*AWSElastiCacheReplicationGroup_NodeGroupConfiguration{}
	for name, resource := range template.Resources {
		result := &AWSElastiCacheReplicationGroup_NodeGroupConfiguration{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSElastiCacheReplicationGroup_NodeGroupConfigurationWithName retrieves all AWSElastiCacheReplicationGroup_NodeGroupConfiguration items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSElastiCacheReplicationGroup_NodeGroupConfiguration(name string, template *Template) (*AWSElastiCacheReplicationGroup_NodeGroupConfiguration, error) {

	result := &AWSElastiCacheReplicationGroup_NodeGroupConfiguration{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSElastiCacheReplicationGroup_NodeGroupConfiguration{}, errors.New("resource not found")

}
