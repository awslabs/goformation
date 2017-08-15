package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::OpsWorks::Stack.ChefConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-chefconfiguration.html
type AWSOpsWorksStack_ChefConfiguration struct {

	// BerkshelfVersion AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-chefconfiguration.html#cfn-opsworks-chefconfiguration-berkshelfversion
	BerkshelfVersion string `json:"BerkshelfVersion"`

	// ManageBerkshelf AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-chefconfiguration.html#cfn-opsworks-chefconfiguration-berkshelfversion
	ManageBerkshelf bool `json:"ManageBerkshelf"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSOpsWorksStack_ChefConfiguration) AWSCloudFormationType() string {
	return "AWS::OpsWorks::Stack.ChefConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSOpsWorksStack_ChefConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSOpsWorksStack_ChefConfigurationResources retrieves all AWSOpsWorksStack_ChefConfiguration items from a CloudFormation template
func GetAllAWSOpsWorksStack_ChefConfiguration(template *Template) map[string]*AWSOpsWorksStack_ChefConfiguration {

	results := map[string]*AWSOpsWorksStack_ChefConfiguration{}
	for name, resource := range template.Resources {
		result := &AWSOpsWorksStack_ChefConfiguration{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSOpsWorksStack_ChefConfigurationWithName retrieves all AWSOpsWorksStack_ChefConfiguration items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSOpsWorksStack_ChefConfiguration(name string, template *Template) (*AWSOpsWorksStack_ChefConfiguration, error) {

	result := &AWSOpsWorksStack_ChefConfiguration{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSOpsWorksStack_ChefConfiguration{}, errors.New("resource not found")

}
