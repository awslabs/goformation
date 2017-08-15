package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Redshift::ClusterParameterGroup.Parameter AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-property-redshift-clusterparametergroup-parameter.html
type AWSRedshiftClusterParameterGroup_Parameter struct {

	// ParameterName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-property-redshift-clusterparametergroup-parameter.html#cfn-redshift-clusterparametergroup-parameter-parametername

	ParameterName string `json:"ParameterName"`

	// ParameterValue AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-property-redshift-clusterparametergroup-parameter.html#cfn-redshift-clusterparametergroup-parameter-parametervalue

	ParameterValue string `json:"ParameterValue"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRedshiftClusterParameterGroup_Parameter) AWSCloudFormationType() string {
	return "AWS::Redshift::ClusterParameterGroup.Parameter"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSRedshiftClusterParameterGroup_Parameter) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSRedshiftClusterParameterGroup_ParameterResources retrieves all AWSRedshiftClusterParameterGroup_Parameter items from a CloudFormation template
func GetAllAWSRedshiftClusterParameterGroup_Parameter(template *Template) map[string]*AWSRedshiftClusterParameterGroup_Parameter {

	results := map[string]*AWSRedshiftClusterParameterGroup_Parameter{}
	for name, resource := range template.Resources {
		result := &AWSRedshiftClusterParameterGroup_Parameter{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSRedshiftClusterParameterGroup_ParameterWithName retrieves all AWSRedshiftClusterParameterGroup_Parameter items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSRedshiftClusterParameterGroup_Parameter(name string, template *Template) (*AWSRedshiftClusterParameterGroup_Parameter, error) {

	result := &AWSRedshiftClusterParameterGroup_Parameter{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSRedshiftClusterParameterGroup_Parameter{}, errors.New("resource not found")

}
