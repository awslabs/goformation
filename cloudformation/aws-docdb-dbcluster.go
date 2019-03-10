package cloudformation

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// AWSDocDBDBCluster AWS CloudFormation Resource (AWS::DocDB::DBCluster)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html
type AWSDocDBDBCluster struct {

	// AvailabilityZones AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-availabilityzones
	AvailabilityZones []string `json:"AvailabilityZones,omitempty"`

	// BackupRetentionPeriod AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-backupretentionperiod
	BackupRetentionPeriod int `json:"BackupRetentionPeriod,omitempty"`

	// DBClusterIdentifier AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-dbclusteridentifier
	DBClusterIdentifier string `json:"DBClusterIdentifier,omitempty"`

	// DBClusterParameterGroupName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-dbclusterparametergroupname
	DBClusterParameterGroupName string `json:"DBClusterParameterGroupName,omitempty"`

	// DBSubnetGroupName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-dbsubnetgroupname
	DBSubnetGroupName string `json:"DBSubnetGroupName,omitempty"`

	// EngineVersion AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-engineversion
	EngineVersion string `json:"EngineVersion,omitempty"`

	// KmsKeyId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-kmskeyid
	KmsKeyId string `json:"KmsKeyId,omitempty"`

	// MasterUserPassword AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-masteruserpassword
	MasterUserPassword string `json:"MasterUserPassword,omitempty"`

	// MasterUsername AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-masterusername
	MasterUsername string `json:"MasterUsername,omitempty"`

	// Port AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-port
	Port int `json:"Port,omitempty"`

	// PreferredBackupWindow AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-preferredbackupwindow
	PreferredBackupWindow string `json:"PreferredBackupWindow,omitempty"`

	// PreferredMaintenanceWindow AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-preferredmaintenancewindow
	PreferredMaintenanceWindow string `json:"PreferredMaintenanceWindow,omitempty"`

	// SnapshotIdentifier AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-snapshotidentifier
	SnapshotIdentifier string `json:"SnapshotIdentifier,omitempty"`

	// StorageEncrypted AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-storageencrypted
	StorageEncrypted bool `json:"StorageEncrypted,omitempty"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-tags
	Tags []Tag `json:"Tags,omitempty"`

	// VpcSecurityGroupIds AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-docdb-dbcluster.html#cfn-docdb-dbcluster-vpcsecuritygroupids
	VpcSecurityGroupIds []string `json:"VpcSecurityGroupIds,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy

	// _dependsOn stores the logical ID of the resources to be created before this resource
	_dependsOn []string

	// _metadata stores structured data associated with this resource
	_metadata map[string]interface{}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDocDBDBCluster) AWSCloudFormationType() string {
	return "AWS::DocDB::DBCluster"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSDocDBDBCluster) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSDocDBDBCluster) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// Metadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSDocDBDBCluster) Metadata() map[string]interface{} {
	return r._metadata
}

// SetMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSDocDBDBCluster) SetMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSDocDBDBCluster) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r AWSDocDBDBCluster) MarshalJSON() ([]byte, error) {
	type Properties AWSDocDBDBCluster
	return json.Marshal(&struct {
		Type           string
		Properties     Properties
		DependsOn      []string               `json:"DependsOn,omitempty"`
		Metadata       map[string]interface{} `json:"Metadata,omitempty"`
		DeletionPolicy DeletionPolicy         `json:"DeletionPolicy,omitempty"`
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
func (r *AWSDocDBDBCluster) UnmarshalJSON(b []byte) error {
	type Properties AWSDocDBDBCluster
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
		*r = AWSDocDBDBCluster(*res.Properties)
	}
	if res.DependsOn != nil {
		r._dependsOn = res.DependsOn
	}
	if res.Metadata != nil {
		r._metadata = res.Metadata
	}

	return nil
}
