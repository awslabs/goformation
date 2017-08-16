package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSRedshiftClusterParameterGroup AWS CloudFormation Resource (AWS::Redshift::ClusterParameterGroup)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clusterparametergroup.html
type AWSRedshiftClusterParameterGroup struct {

	// Description AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clusterparametergroup.html#cfn-redshift-clusterparametergroup-description
	Description string `json:"Description"`

	// ParameterGroupFamily AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clusterparametergroup.html#cfn-redshift-clusterparametergroup-parametergroupfamily
	ParameterGroupFamily string `json:"ParameterGroupFamily"`

	// Parameters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clusterparametergroup.html#cfn-redshift-clusterparametergroup-parameters
	Parameters []AWSRedshiftClusterParameterGroup_Parameter `json:"Parameters"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-redshift-clusterparametergroup.html#cfn-redshift-clusterparametergroup-tags
	Tags []Tag `json:"Tags"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRedshiftClusterParameterGroup) AWSCloudFormationType() string {
	return "AWS::Redshift::ClusterParameterGroup"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSRedshiftClusterParameterGroup) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSRedshiftClusterParameterGroupResources retrieves all AWSRedshiftClusterParameterGroup items from a CloudFormation template
func (t *Template) GetAllAWSRedshiftClusterParameterGroupResources() map[string]*AWSRedshiftClusterParameterGroup {

	results := map[string]*AWSRedshiftClusterParameterGroup{}
	for name, resource := range t.Resources {
		result := &AWSRedshiftClusterParameterGroup{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSRedshiftClusterParameterGroupWithName retrieves all AWSRedshiftClusterParameterGroup items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRedshiftClusterParameterGroupWithName(name string) (*AWSRedshiftClusterParameterGroup, error) {

	result := &AWSRedshiftClusterParameterGroup{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSRedshiftClusterParameterGroup{}, errors.New("resource not found")

}
