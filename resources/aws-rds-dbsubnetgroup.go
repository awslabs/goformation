package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSRDSDBSubnetGroup AWS CloudFormation Resource (AWS::RDS::DBSubnetGroup)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbsubnet-group.html
type AWSRDSDBSubnetGroup struct {

	// DBSubnetGroupDescription AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbsubnet-group.html#cfn-rds-dbsubnetgroup-dbsubnetgroupdescription
	DBSubnetGroupDescription string `json:"DBSubnetGroupDescription"`

	// SubnetIds AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbsubnet-group.html#cfn-rds-dbsubnetgroup-subnetids
	SubnetIds []string `json:"SubnetIds"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbsubnet-group.html#cfn-rds-dbsubnetgroup-tags
	Tags []Tag `json:"Tags"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRDSDBSubnetGroup) AWSCloudFormationType() string {
	return "AWS::RDS::DBSubnetGroup"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSRDSDBSubnetGroup) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSRDSDBSubnetGroupResources retrieves all AWSRDSDBSubnetGroup items from a CloudFormation template
func (t *Template) GetAllAWSRDSDBSubnetGroupResources() map[string]*AWSRDSDBSubnetGroup {

	results := map[string]*AWSRDSDBSubnetGroup{}
	for name, resource := range t.Resources {
		result := &AWSRDSDBSubnetGroup{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSRDSDBSubnetGroupWithName retrieves all AWSRDSDBSubnetGroup items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRDSDBSubnetGroupWithName(name string) (*AWSRDSDBSubnetGroup, error) {

	result := &AWSRDSDBSubnetGroup{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSRDSDBSubnetGroup{}, errors.New("resource not found")

}
