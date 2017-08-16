package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSRedshiftClusterSecurityGroup AWS CloudFormation Resource (AWS::Redshift::ClusterSecurityGroup)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersecuritygroup.html
type AWSRedshiftClusterSecurityGroup struct {

	// Description AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersecuritygroup.html#cfn-redshift-clustersecuritygroup-description
	Description string `json:"Description"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersecuritygroup.html#cfn-redshift-clustersecuritygroup-tags
	Tags []Tag `json:"Tags"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRedshiftClusterSecurityGroup) AWSCloudFormationType() string {
	return "AWS::Redshift::ClusterSecurityGroup"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSRedshiftClusterSecurityGroup) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSRedshiftClusterSecurityGroupResources retrieves all AWSRedshiftClusterSecurityGroup items from a CloudFormation template
func (t *Template) GetAllAWSRedshiftClusterSecurityGroupResources() map[string]*AWSRedshiftClusterSecurityGroup {

	results := map[string]*AWSRedshiftClusterSecurityGroup{}
	for name, resource := range t.Resources {
		result := &AWSRedshiftClusterSecurityGroup{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSRedshiftClusterSecurityGroupWithName retrieves all AWSRedshiftClusterSecurityGroup items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRedshiftClusterSecurityGroupWithName(name string) (*AWSRedshiftClusterSecurityGroup, error) {

	result := &AWSRedshiftClusterSecurityGroup{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSRedshiftClusterSecurityGroup{}, errors.New("resource not found")

}
