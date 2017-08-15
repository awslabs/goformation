package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::OpsWorks::Stack.RdsDbInstance AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-rdsdbinstance.html
type AWSOpsWorksStack_RdsDbInstance struct {

	// DbPassword AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-rdsdbinstance.html#cfn-opsworks-stack-rdsdbinstance-dbpassword

	DbPassword string `json:"DbPassword"`

	// DbUser AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-rdsdbinstance.html#cfn-opsworks-stack-rdsdbinstance-dbuser

	DbUser string `json:"DbUser"`

	// RdsDbInstanceArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-rdsdbinstance.html#cfn-opsworks-stack-rdsdbinstance-rdsdbinstancearn

	RdsDbInstanceArn string `json:"RdsDbInstanceArn"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSOpsWorksStack_RdsDbInstance) AWSCloudFormationType() string {
	return "AWS::OpsWorks::Stack.RdsDbInstance"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSOpsWorksStack_RdsDbInstance) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSOpsWorksStack_RdsDbInstanceResources retrieves all AWSOpsWorksStack_RdsDbInstance items from a CloudFormation template
func GetAllAWSOpsWorksStack_RdsDbInstance(template *Template) map[string]*AWSOpsWorksStack_RdsDbInstance {

	results := map[string]*AWSOpsWorksStack_RdsDbInstance{}
	for name, resource := range template.Resources {
		result := &AWSOpsWorksStack_RdsDbInstance{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSOpsWorksStack_RdsDbInstanceWithName retrieves all AWSOpsWorksStack_RdsDbInstance items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSOpsWorksStack_RdsDbInstance(name string, template *Template) (*AWSOpsWorksStack_RdsDbInstance, error) {

	result := &AWSOpsWorksStack_RdsDbInstance{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSOpsWorksStack_RdsDbInstance{}, errors.New("resource not found")

}
