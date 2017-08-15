package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::DMS::ReplicationSubnetGroup AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationsubnetgroup.html
type AWSDMSReplicationSubnetGroup struct {

	// ReplicationSubnetGroupDescription AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationsubnetgroup.html#cfn-dms-replicationsubnetgroup-replicationsubnetgroupdescription
	ReplicationSubnetGroupDescription string `json:"ReplicationSubnetGroupDescription"`

	// ReplicationSubnetGroupIdentifier AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationsubnetgroup.html#cfn-dms-replicationsubnetgroup-replicationsubnetgroupidentifier
	ReplicationSubnetGroupIdentifier string `json:"ReplicationSubnetGroupIdentifier"`

	// SubnetIds AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationsubnetgroup.html#cfn-dms-replicationsubnetgroup-subnetids
	SubnetIds []string `json:"SubnetIds"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-replicationsubnetgroup.html#cfn-dms-replicationsubnetgroup-tags
	Tags []Tag `json:"Tags"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDMSReplicationSubnetGroup) AWSCloudFormationType() string {
	return "AWS::DMS::ReplicationSubnetGroup"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSDMSReplicationSubnetGroup) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSDMSReplicationSubnetGroupResources retrieves all AWSDMSReplicationSubnetGroup items from a CloudFormation template
func GetAllAWSDMSReplicationSubnetGroupResources(template *Template) map[string]*AWSDMSReplicationSubnetGroup {

	results := map[string]*AWSDMSReplicationSubnetGroup{}
	for name, resource := range template.Resources {
		result := &AWSDMSReplicationSubnetGroup{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSDMSReplicationSubnetGroupWithName retrieves all AWSDMSReplicationSubnetGroup items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSDMSReplicationSubnetGroupWithName(name string, template *Template) (*AWSDMSReplicationSubnetGroup, error) {

	result := &AWSDMSReplicationSubnetGroup{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSDMSReplicationSubnetGroup{}, errors.New("resource not found")

}
