package resources

// AWS::DMS::ReplicationInstance AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationinstance.html
type AWSDMSReplicationInstance struct {

	// AllocatedStorage AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationinstance.html#cfn-dms-replicationinstance-allocatedstorage
	AllocatedStorage int64 `json:"AllocatedStorage"`

	// AllowMajorVersionUpgrade AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationinstance.html#cfn-dms-replicationinstance-allowmajorversionupgrade
	AllowMajorVersionUpgrade bool `json:"AllowMajorVersionUpgrade"`

	// AutoMinorVersionUpgrade AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationinstance.html#cfn-dms-replicationinstance-autominorversionupgrade
	AutoMinorVersionUpgrade bool `json:"AutoMinorVersionUpgrade"`

	// AvailabilityZone AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationinstance.html#cfn-dms-replicationinstance-availabilityzone
	AvailabilityZone string `json:"AvailabilityZone"`

	// EngineVersion AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationinstance.html#cfn-dms-replicationinstance-engineversion
	EngineVersion string `json:"EngineVersion"`

	// KmsKeyId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationinstance.html#cfn-dms-replicationinstance-kmskeyid
	KmsKeyId string `json:"KmsKeyId"`

	// MultiAZ AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationinstance.html#cfn-dms-replicationinstance-multiaz
	MultiAZ bool `json:"MultiAZ"`

	// PreferredMaintenanceWindow AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationinstance.html#cfn-dms-replicationinstance-preferredmaintenancewindow
	PreferredMaintenanceWindow string `json:"PreferredMaintenanceWindow"`

	// PubliclyAccessible AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationinstance.html#cfn-dms-replicationinstance-publiclyaccessible
	PubliclyAccessible bool `json:"PubliclyAccessible"`

	// ReplicationInstanceClass AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationinstance.html#cfn-dms-replicationinstance-replicationinstanceclass
	ReplicationInstanceClass string `json:"ReplicationInstanceClass"`

	// ReplicationInstanceIdentifier AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationinstance.html#cfn-dms-replicationinstance-replicationinstanceidentifier
	ReplicationInstanceIdentifier string `json:"ReplicationInstanceIdentifier"`

	// ReplicationSubnetGroupIdentifier AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationinstance.html#cfn-dms-replicationinstance-replicationsubnetgroupidentifier
	ReplicationSubnetGroupIdentifier string `json:"ReplicationSubnetGroupIdentifier"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationinstance.html#cfn-dms-replicationinstance-tags
	Tags []AWSDMSReplicationInstanceTag `json:"Tags"`

	// VpcSecurityGroupIds AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationinstance.html#cfn-dms-replicationinstance-vpcsecuritygroupids
	VpcSecurityGroupIds []AWSDMSReplicationInstancestring `json:"VpcSecurityGroupIds"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDMSReplicationInstance) AWSCloudFormationType() string {
	return "AWS::DMS::ReplicationInstance"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSDMSReplicationInstance) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
