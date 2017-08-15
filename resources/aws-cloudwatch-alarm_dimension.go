package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::CloudWatch::Alarm.Dimension AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cw-dimension.html
type AWSCloudWatchAlarm_Dimension struct {

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cw-dimension.html#cfn-cloudwatch-alarm-dimension-name
	Name string `json:"Name"`

	// Value AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cw-dimension.html#cfn-cloudwatch-alarm-dimension-value
	Value string `json:"Value"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCloudWatchAlarm_Dimension) AWSCloudFormationType() string {
	return "AWS::CloudWatch::Alarm.Dimension"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCloudWatchAlarm_Dimension) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCloudWatchAlarm_DimensionResources retrieves all AWSCloudWatchAlarm_Dimension items from a CloudFormation template
func GetAllAWSCloudWatchAlarm_Dimension(template *Template) map[string]*AWSCloudWatchAlarm_Dimension {

	results := map[string]*AWSCloudWatchAlarm_Dimension{}
	for name, resource := range template.Resources {
		result := &AWSCloudWatchAlarm_Dimension{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCloudWatchAlarm_DimensionWithName retrieves all AWSCloudWatchAlarm_Dimension items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCloudWatchAlarm_Dimension(name string, template *Template) (*AWSCloudWatchAlarm_Dimension, error) {

	result := &AWSCloudWatchAlarm_Dimension{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCloudWatchAlarm_Dimension{}, errors.New("resource not found")

}
