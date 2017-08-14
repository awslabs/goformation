package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::OpsWorks::Layer.ShutdownEventConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-lifecycleeventconfiguration-shutdowneventconfiguration.html
type AWSOpsWorksLayer_ShutdownEventConfiguration struct {

	// DelayUntilElbConnectionsDrained AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-lifecycleeventconfiguration-shutdowneventconfiguration.html#cfn-opsworks-layer-lifecycleconfiguration-shutdowneventconfiguration-delayuntilelbconnectionsdrained

	DelayUntilElbConnectionsDrained bool `json:"DelayUntilElbConnectionsDrained"`

	// ExecutionTimeout AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-layer-lifecycleeventconfiguration-shutdowneventconfiguration.html#cfn-opsworks-layer-lifecycleconfiguration-shutdowneventconfiguration-executiontimeout

	ExecutionTimeout int64 `json:"ExecutionTimeout"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSOpsWorksLayer_ShutdownEventConfiguration) AWSCloudFormationType() string {
	return "AWS::OpsWorks::Layer.ShutdownEventConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSOpsWorksLayer_ShutdownEventConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSOpsWorksLayer_ShutdownEventConfigurationResources retrieves all AWSOpsWorksLayer_ShutdownEventConfiguration items from a CloudFormation template
func GetAllAWSOpsWorksLayer_ShutdownEventConfiguration(template *Template) map[string]*AWSOpsWorksLayer_ShutdownEventConfiguration {

	results := map[string]*AWSOpsWorksLayer_ShutdownEventConfiguration{}
	for name, resource := range template.Resources {
		result := &AWSOpsWorksLayer_ShutdownEventConfiguration{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSOpsWorksLayer_ShutdownEventConfigurationWithName retrieves all AWSOpsWorksLayer_ShutdownEventConfiguration items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSOpsWorksLayer_ShutdownEventConfiguration(name string, template *Template) (*AWSOpsWorksLayer_ShutdownEventConfiguration, error) {

	result := &AWSOpsWorksLayer_ShutdownEventConfiguration{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSOpsWorksLayer_ShutdownEventConfiguration{}, errors.New("resource not found")

}
