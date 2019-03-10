package resources

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/awslabs/goformation/cloudformation/policies"
)

// AWSElastiCacheReplicationGroup AWS CloudFormation Resource (AWS::ElastiCache::ReplicationGroup)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html
type AWSElastiCacheReplicationGroup struct {

	// AtRestEncryptionEnabled AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-atrestencryptionenabled
	AtRestEncryptionEnabled bool `json:"AtRestEncryptionEnabled,omitempty"`

	// AuthToken AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-authtoken
	AuthToken string `json:"AuthToken,omitempty"`

	// AutoMinorVersionUpgrade AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-autominorversionupgrade
	AutoMinorVersionUpgrade bool `json:"AutoMinorVersionUpgrade,omitempty"`

	// AutomaticFailoverEnabled AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-automaticfailoverenabled
	AutomaticFailoverEnabled bool `json:"AutomaticFailoverEnabled,omitempty"`

	// CacheNodeType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-cachenodetype
	CacheNodeType string `json:"CacheNodeType,omitempty"`

	// CacheParameterGroupName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-cacheparametergroupname
	CacheParameterGroupName string `json:"CacheParameterGroupName,omitempty"`

	// CacheSecurityGroupNames AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-cachesecuritygroupnames
	CacheSecurityGroupNames []string `json:"CacheSecurityGroupNames,omitempty"`

	// CacheSubnetGroupName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-cachesubnetgroupname
	CacheSubnetGroupName string `json:"CacheSubnetGroupName,omitempty"`

	// Engine AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-engine
	Engine string `json:"Engine,omitempty"`

	// EngineVersion AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-engineversion
	EngineVersion string `json:"EngineVersion,omitempty"`

	// NodeGroupConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-nodegroupconfiguration
	NodeGroupConfiguration []AWSElastiCacheReplicationGroup_NodeGroupConfiguration `json:"NodeGroupConfiguration,omitempty"`

	// NotificationTopicArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-notificationtopicarn
	NotificationTopicArn string `json:"NotificationTopicArn,omitempty"`

	// NumCacheClusters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-numcacheclusters
	NumCacheClusters int `json:"NumCacheClusters,omitempty"`

	// NumNodeGroups AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-numnodegroups
	NumNodeGroups int `json:"NumNodeGroups,omitempty"`

	// Port AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-port
	Port int `json:"Port,omitempty"`

	// PreferredCacheClusterAZs AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-preferredcacheclusterazs
	PreferredCacheClusterAZs []string `json:"PreferredCacheClusterAZs,omitempty"`

	// PreferredMaintenanceWindow AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-preferredmaintenancewindow
	PreferredMaintenanceWindow string `json:"PreferredMaintenanceWindow,omitempty"`

	// PrimaryClusterId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-primaryclusterid
	PrimaryClusterId string `json:"PrimaryClusterId,omitempty"`

	// ReplicasPerNodeGroup AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-replicaspernodegroup
	ReplicasPerNodeGroup int `json:"ReplicasPerNodeGroup,omitempty"`

	// ReplicationGroupDescription AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-replicationgroupdescription
	ReplicationGroupDescription string `json:"ReplicationGroupDescription,omitempty"`

	// ReplicationGroupId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-replicationgroupid
	ReplicationGroupId string `json:"ReplicationGroupId,omitempty"`

	// SecurityGroupIds AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-securitygroupids
	SecurityGroupIds []string `json:"SecurityGroupIds,omitempty"`

	// SnapshotArns AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-snapshotarns
	SnapshotArns []string `json:"SnapshotArns,omitempty"`

	// SnapshotName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-snapshotname
	SnapshotName string `json:"SnapshotName,omitempty"`

	// SnapshotRetentionLimit AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-snapshotretentionlimit
	SnapshotRetentionLimit int `json:"SnapshotRetentionLimit,omitempty"`

	// SnapshotWindow AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-snapshotwindow
	SnapshotWindow string `json:"SnapshotWindow,omitempty"`

	// SnapshottingClusterId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-snapshottingclusterid
	SnapshottingClusterId string `json:"SnapshottingClusterId,omitempty"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-tags
	Tags []Tag `json:"Tags,omitempty"`

	// TransitEncryptionEnabled AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticache-replicationgroup.html#cfn-elasticache-replicationgroup-transitencryptionenabled
	TransitEncryptionEnabled bool `json:"TransitEncryptionEnabled,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy policies.DeletionPolicy

	// _dependsOn stores the logical ID of the resources to be created before this resource
	_dependsOn []string

	// _metadata stores structured data associated with this resource
	_metadata map[string]interface{}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElastiCacheReplicationGroup) AWSCloudFormationType() string {
	return "AWS::ElastiCache::ReplicationGroup"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSElastiCacheReplicationGroup) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSElastiCacheReplicationGroup) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// Metadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSElastiCacheReplicationGroup) Metadata() map[string]interface{} {
	return r._metadata
}

// SetMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSElastiCacheReplicationGroup) SetMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSElastiCacheReplicationGroup) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r AWSElastiCacheReplicationGroup) MarshalJSON() ([]byte, error) {
	type Properties AWSElastiCacheReplicationGroup
	return json.Marshal(&struct {
		Type           string
		Properties     Properties
		DependsOn      []string                `json:"DependsOn,omitempty"`
		Metadata       map[string]interface{}  `json:"Metadata,omitempty"`
		DeletionPolicy policies.DeletionPolicy `json:"DeletionPolicy,omitempty"`
	}{
		Type:           r.AWSCloudFormationType(),
		Properties:     (Properties)(r),
		DependsOn:      r._dependsOn,
		Metadata:       r._metadata,
		DeletionPolicy: r._deletionPolicy,
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *AWSElastiCacheReplicationGroup) UnmarshalJSON(b []byte) error {
	type Properties AWSElastiCacheReplicationGroup
	res := &struct {
		Type       string
		Properties *Properties
		DependsOn  []string
		Metadata   map[string]interface{}
	}{}

	dec := json.NewDecoder(bytes.NewReader(b))
	dec.DisallowUnknownFields() // Force error if unknown field is found

	if err := dec.Decode(&res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}

	// If the resource has no Properties set, it could be nil
	if res.Properties != nil {
		*r = AWSElastiCacheReplicationGroup(*res.Properties)
	}
	if res.DependsOn != nil {
		r._dependsOn = res.DependsOn
	}
	if res.Metadata != nil {
		r._metadata = res.Metadata
	}

	return nil
}
