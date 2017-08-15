package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::RDS::DBClusterParameterGroup AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbclusterparametergroup.html
type AWSRDSDBClusterParameterGroup struct {

	// Description AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbclusterparametergroup.html#cfn-rds-dbclusterparametergroup-description
	Description string `json:"Description"`

	// Family AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbclusterparametergroup.html#cfn-rds-dbclusterparametergroup-family
	Family string `json:"Family"`

	// Parameters AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbclusterparametergroup.html#cfn-rds-dbclusterparametergroup-parameters
	Parameters interface{} `json:"Parameters"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-dbclusterparametergroup.html#cfn-rds-dbclusterparametergroup-tags
	Tags []Tag `json:"Tags"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRDSDBClusterParameterGroup) AWSCloudFormationType() string {
	return "AWS::RDS::DBClusterParameterGroup"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSRDSDBClusterParameterGroup) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSRDSDBClusterParameterGroupResources retrieves all AWSRDSDBClusterParameterGroup items from a CloudFormation template
func GetAllAWSRDSDBClusterParameterGroup(template *Template) map[string]*AWSRDSDBClusterParameterGroup {

	results := map[string]*AWSRDSDBClusterParameterGroup{}
	for name, resource := range template.Resources {
		result := &AWSRDSDBClusterParameterGroup{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSRDSDBClusterParameterGroupWithName retrieves all AWSRDSDBClusterParameterGroup items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSRDSDBClusterParameterGroup(name string, template *Template) (*AWSRDSDBClusterParameterGroup, error) {

	result := &AWSRDSDBClusterParameterGroup{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSRDSDBClusterParameterGroup{}, errors.New("resource not found")

}
