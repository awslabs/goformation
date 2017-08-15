package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::RDS::DBParameterGroup AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbparametergroup.html
type AWSRDSDBParameterGroup struct {

	// Description AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbparametergroup.html#cfn-rds-dbparametergroup-description
	Description string `json:"Description"`

	// Family AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbparametergroup.html#cfn-rds-dbparametergroup-family
	Family string `json:"Family"`

	// Parameters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbparametergroup.html#cfn-rds-dbparametergroup-parameters
	Parameters map[string]string `json:"Parameters"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-rds-dbparametergroup.html#cfn-rds-dbparametergroup-tags
	Tags []Tag `json:"Tags"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRDSDBParameterGroup) AWSCloudFormationType() string {
	return "AWS::RDS::DBParameterGroup"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSRDSDBParameterGroup) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSRDSDBParameterGroupResources retrieves all AWSRDSDBParameterGroup items from a CloudFormation template
func GetAllAWSRDSDBParameterGroupResources(template *Template) map[string]*AWSRDSDBParameterGroup {

	results := map[string]*AWSRDSDBParameterGroup{}
	for name, resource := range template.Resources {
		result := &AWSRDSDBParameterGroup{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSRDSDBParameterGroupWithName retrieves all AWSRDSDBParameterGroup items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSRDSDBParameterGroupWithName(name string, template *Template) (*AWSRDSDBParameterGroup, error) {

	result := &AWSRDSDBParameterGroup{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSRDSDBParameterGroup{}, errors.New("resource not found")

}
