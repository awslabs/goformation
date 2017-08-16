package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSRDSDBParameterGroup AWS CloudFormation Resource (AWS::RDS::DBParameterGroup)
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
func (t *Template) GetAllAWSRDSDBParameterGroupResources() map[string]*AWSRDSDBParameterGroup {

	results := map[string]*AWSRDSDBParameterGroup{}
	for name, resource := range t.Resources {
		result := &AWSRDSDBParameterGroup{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSRDSDBParameterGroupWithName retrieves all AWSRDSDBParameterGroup items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSRDSDBParameterGroupWithName(name string) (*AWSRDSDBParameterGroup, error) {

	result := &AWSRDSDBParameterGroup{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSRDSDBParameterGroup{}, errors.New("resource not found")

}
