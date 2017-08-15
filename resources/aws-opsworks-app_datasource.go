package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::OpsWorks::App.DataSource AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-datasource.html
type AWSOpsWorksApp_DataSource struct {

	// Arn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-datasource.html#cfn-opsworks-app-datasource-arn
	Arn string `json:"Arn"`

	// DatabaseName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-datasource.html#cfn-opsworks-app-datasource-databasename
	DatabaseName string `json:"DatabaseName"`

	// Type AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-datasource.html#cfn-opsworks-app-datasource-type
	Type string `json:"Type"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSOpsWorksApp_DataSource) AWSCloudFormationType() string {
	return "AWS::OpsWorks::App.DataSource"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSOpsWorksApp_DataSource) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSOpsWorksApp_DataSourceResources retrieves all AWSOpsWorksApp_DataSource items from a CloudFormation template
func GetAllAWSOpsWorksApp_DataSource(template *Template) map[string]*AWSOpsWorksApp_DataSource {

	results := map[string]*AWSOpsWorksApp_DataSource{}
	for name, resource := range template.Resources {
		result := &AWSOpsWorksApp_DataSource{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSOpsWorksApp_DataSourceWithName retrieves all AWSOpsWorksApp_DataSource items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSOpsWorksApp_DataSource(name string, template *Template) (*AWSOpsWorksApp_DataSource, error) {

	result := &AWSOpsWorksApp_DataSource{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSOpsWorksApp_DataSource{}, errors.New("resource not found")

}
