package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::OpsWorks::Instance.TimeBasedAutoScaling AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-timebasedautoscaling.html
type AWSOpsWorksInstance_TimeBasedAutoScaling struct {

	// Friday AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-timebasedautoscaling.html#cfn-opsworks-instance-timebasedautoscaling-friday
	Friday map[string]string `json:"Friday"`

	// Monday AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-timebasedautoscaling.html#cfn-opsworks-instance-timebasedautoscaling-monday
	Monday map[string]string `json:"Monday"`

	// Saturday AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-timebasedautoscaling.html#cfn-opsworks-instance-timebasedautoscaling-saturday
	Saturday map[string]string `json:"Saturday"`

	// Sunday AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-timebasedautoscaling.html#cfn-opsworks-instance-timebasedautoscaling-sunday
	Sunday map[string]string `json:"Sunday"`

	// Thursday AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-timebasedautoscaling.html#cfn-opsworks-instance-timebasedautoscaling-thursday
	Thursday map[string]string `json:"Thursday"`

	// Tuesday AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-timebasedautoscaling.html#cfn-opsworks-instance-timebasedautoscaling-tuesday
	Tuesday map[string]string `json:"Tuesday"`

	// Wednesday AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-opsworks-instance-timebasedautoscaling.html#cfn-opsworks-instance-timebasedautoscaling-wednesday
	Wednesday map[string]string `json:"Wednesday"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSOpsWorksInstance_TimeBasedAutoScaling) AWSCloudFormationType() string {
	return "AWS::OpsWorks::Instance.TimeBasedAutoScaling"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSOpsWorksInstance_TimeBasedAutoScaling) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSOpsWorksInstance_TimeBasedAutoScalingResources retrieves all AWSOpsWorksInstance_TimeBasedAutoScaling items from a CloudFormation template
func GetAllAWSOpsWorksInstance_TimeBasedAutoScaling(template *Template) map[string]*AWSOpsWorksInstance_TimeBasedAutoScaling {

	results := map[string]*AWSOpsWorksInstance_TimeBasedAutoScaling{}
	for name, resource := range template.Resources {
		result := &AWSOpsWorksInstance_TimeBasedAutoScaling{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSOpsWorksInstance_TimeBasedAutoScalingWithName retrieves all AWSOpsWorksInstance_TimeBasedAutoScaling items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSOpsWorksInstance_TimeBasedAutoScaling(name string, template *Template) (*AWSOpsWorksInstance_TimeBasedAutoScaling, error) {

	result := &AWSOpsWorksInstance_TimeBasedAutoScaling{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSOpsWorksInstance_TimeBasedAutoScaling{}, errors.New("resource not found")

}
