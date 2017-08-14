package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::OpsWorks::Stack.StackConfigurationManager AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-stackconfigmanager.html
type AWSOpsWorksStack_StackConfigurationManager struct {

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-stackconfigmanager.html#cfn-opsworks-configmanager-name

	Name string `json:"Name"`

	// Version AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-stack-stackconfigmanager.html#cfn-opsworks-configmanager-version

	Version string `json:"Version"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSOpsWorksStack_StackConfigurationManager) AWSCloudFormationType() string {
	return "AWS::OpsWorks::Stack.StackConfigurationManager"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSOpsWorksStack_StackConfigurationManager) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSOpsWorksStack_StackConfigurationManagerResources retrieves all AWSOpsWorksStack_StackConfigurationManager items from a CloudFormation template
func GetAllAWSOpsWorksStack_StackConfigurationManager(template *Template) map[string]*AWSOpsWorksStack_StackConfigurationManager {

	results := map[string]*AWSOpsWorksStack_StackConfigurationManager{}
	for name, resource := range template.Resources {
		result := &AWSOpsWorksStack_StackConfigurationManager{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSOpsWorksStack_StackConfigurationManagerWithName retrieves all AWSOpsWorksStack_StackConfigurationManager items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSOpsWorksStack_StackConfigurationManager(name string, template *Template) (*AWSOpsWorksStack_StackConfigurationManager, error) {

	result := &AWSOpsWorksStack_StackConfigurationManager{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSOpsWorksStack_StackConfigurationManager{}, errors.New("resource not found")

}
