package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::OpsWorks::Layer.LifecycleEventConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-lifecycleeventconfiguration.html
type AWSOpsWorksLayer_LifecycleEventConfiguration struct {

	// ShutdownEventConfiguration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-lifecycleeventconfiguration.html#cfn-opsworks-layer-lifecycleconfiguration-shutdowneventconfiguration
	ShutdownEventConfiguration AWSOpsWorksLayer_ShutdownEventConfiguration `json:"ShutdownEventConfiguration"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSOpsWorksLayer_LifecycleEventConfiguration) AWSCloudFormationType() string {
	return "AWS::OpsWorks::Layer.LifecycleEventConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSOpsWorksLayer_LifecycleEventConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSOpsWorksLayer_LifecycleEventConfigurationResources retrieves all AWSOpsWorksLayer_LifecycleEventConfiguration items from a CloudFormation template
func GetAllAWSOpsWorksLayer_LifecycleEventConfiguration(template *Template) map[string]*AWSOpsWorksLayer_LifecycleEventConfiguration {

	results := map[string]*AWSOpsWorksLayer_LifecycleEventConfiguration{}
	for name, resource := range template.Resources {
		result := &AWSOpsWorksLayer_LifecycleEventConfiguration{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSOpsWorksLayer_LifecycleEventConfigurationWithName retrieves all AWSOpsWorksLayer_LifecycleEventConfiguration items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSOpsWorksLayer_LifecycleEventConfiguration(name string, template *Template) (*AWSOpsWorksLayer_LifecycleEventConfiguration, error) {

	result := &AWSOpsWorksLayer_LifecycleEventConfiguration{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSOpsWorksLayer_LifecycleEventConfiguration{}, errors.New("resource not found")

}
