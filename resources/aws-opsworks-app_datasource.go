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
