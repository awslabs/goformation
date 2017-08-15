package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::EMR::SecurityConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-securityconfiguration.html
type AWSEMRSecurityConfiguration struct {

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-securityconfiguration.html#cfn-emr-securityconfiguration-name

	Name string `json:"Name"`

	// SecurityConfiguration AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-emr-securityconfiguration.html#cfn-emr-securityconfiguration-securityconfiguration

	SecurityConfiguration interface{} `json:"SecurityConfiguration"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEMRSecurityConfiguration) AWSCloudFormationType() string {
	return "AWS::EMR::SecurityConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEMRSecurityConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEMRSecurityConfigurationResources retrieves all AWSEMRSecurityConfiguration items from a CloudFormation template
func GetAllAWSEMRSecurityConfiguration(template *Template) map[string]*AWSEMRSecurityConfiguration {

	results := map[string]*AWSEMRSecurityConfiguration{}
	for name, resource := range template.Resources {
		result := &AWSEMRSecurityConfiguration{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEMRSecurityConfigurationWithName retrieves all AWSEMRSecurityConfiguration items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEMRSecurityConfiguration(name string, template *Template) (*AWSEMRSecurityConfiguration, error) {

	result := &AWSEMRSecurityConfiguration{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEMRSecurityConfiguration{}, errors.New("resource not found")

}
