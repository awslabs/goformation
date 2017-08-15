package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

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
	Tags []Tag `json:"Tags"`

	// VpcSecurityGroupIds AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationinstance.html#cfn-dms-replicationinstance-vpcsecuritygroupids
	VpcSecurityGroupIds []string `json:"VpcSecurityGroupIds"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDMSReplicationInstance) AWSCloudFormationType() string {
	return "AWS::DMS::ReplicationInstance"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSDMSReplicationInstance) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSDMSReplicationInstanceResources retrieves all AWSDMSReplicationInstance items from a CloudFormation template
func GetAllAWSDMSReplicationInstance(template *Template) map[string]*AWSDMSReplicationInstance {

	results := map[string]*AWSDMSReplicationInstance{}
	for name, resource := range template.Resources {
		result := &AWSDMSReplicationInstance{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSDMSReplicationInstanceWithName retrieves all AWSDMSReplicationInstance items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSDMSReplicationInstance(name string, template *Template) (*AWSDMSReplicationInstance, error) {

	result := &AWSDMSReplicationInstance{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSDMSReplicationInstance{}, errors.New("resource not found")

}
