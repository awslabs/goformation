package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Redshift::ClusterSubnetGroup AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersubnetgroup.html
type AWSRedshiftClusterSubnetGroup struct {

	// Description AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersubnetgroup.html#cfn-redshift-clustersubnetgroup-description
	Description string `json:"Description"`

	// SubnetIds AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersubnetgroup.html#cfn-redshift-clustersubnetgroup-subnetids
	SubnetIds []string `json:"SubnetIds"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clustersubnetgroup.html#cfn-redshift-clustersubnetgroup-tags
	Tags []Tag `json:"Tags"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRedshiftClusterSubnetGroup) AWSCloudFormationType() string {
	return "AWS::Redshift::ClusterSubnetGroup"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSRedshiftClusterSubnetGroup) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSRedshiftClusterSubnetGroupResources retrieves all AWSRedshiftClusterSubnetGroup items from a CloudFormation template
func GetAllAWSRedshiftClusterSubnetGroup(template *Template) map[string]*AWSRedshiftClusterSubnetGroup {

	results := map[string]*AWSRedshiftClusterSubnetGroup{}
	for name, resource := range template.Resources {
		result := &AWSRedshiftClusterSubnetGroup{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSRedshiftClusterSubnetGroupWithName retrieves all AWSRedshiftClusterSubnetGroup items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSRedshiftClusterSubnetGroup(name string, template *Template) (*AWSRedshiftClusterSubnetGroup, error) {

	result := &AWSRedshiftClusterSubnetGroup{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSRedshiftClusterSubnetGroup{}, errors.New("resource not found")

}
