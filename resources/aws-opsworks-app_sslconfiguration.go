package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::OpsWorks::App.SslConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-sslconfiguration.html
type AWSOpsWorksApp_SslConfiguration struct {

	// Certificate AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-sslconfiguration.html#cfn-opsworks-app-sslconfig-certificate
	Certificate string `json:"Certificate"`

	// Chain AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-sslconfiguration.html#cfn-opsworks-app-sslconfig-chain
	Chain string `json:"Chain"`

	// PrivateKey AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-app-sslconfiguration.html#cfn-opsworks-app-sslconfig-privatekey
	PrivateKey string `json:"PrivateKey"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSOpsWorksApp_SslConfiguration) AWSCloudFormationType() string {
	return "AWS::OpsWorks::App.SslConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSOpsWorksApp_SslConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSOpsWorksApp_SslConfigurationResources retrieves all AWSOpsWorksApp_SslConfiguration items from a CloudFormation template
func GetAllAWSOpsWorksApp_SslConfiguration(template *Template) map[string]*AWSOpsWorksApp_SslConfiguration {

	results := map[string]*AWSOpsWorksApp_SslConfiguration{}
	for name, resource := range template.Resources {
		result := &AWSOpsWorksApp_SslConfiguration{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSOpsWorksApp_SslConfigurationWithName retrieves all AWSOpsWorksApp_SslConfiguration items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSOpsWorksApp_SslConfiguration(name string, template *Template) (*AWSOpsWorksApp_SslConfiguration, error) {

	result := &AWSOpsWorksApp_SslConfiguration{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSOpsWorksApp_SslConfiguration{}, errors.New("resource not found")

}
