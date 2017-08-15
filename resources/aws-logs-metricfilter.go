package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSLogsMetricFilter AWS CloudFormation Resource (AWS::Logs::MetricFilter)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-metricfilter.html
type AWSLogsMetricFilter struct {

	// FilterPattern AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-metricfilter.html#cfn-cwl-metricfilter-filterpattern
	FilterPattern string `json:"FilterPattern"`

	// LogGroupName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-metricfilter.html#cfn-cwl-metricfilter-loggroupname
	LogGroupName string `json:"LogGroupName"`

	// MetricTransformations AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-metricfilter.html#cfn-cwl-metricfilter-metrictransformations
	MetricTransformations []AWSLogsMetricFilter_MetricTransformation `json:"MetricTransformations"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSLogsMetricFilter) AWSCloudFormationType() string {
	return "AWS::Logs::MetricFilter"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSLogsMetricFilter) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSLogsMetricFilterResources retrieves all AWSLogsMetricFilter items from a CloudFormation template
func GetAllAWSLogsMetricFilterResources(template *Template) map[string]*AWSLogsMetricFilter {

	results := map[string]*AWSLogsMetricFilter{}
	for name, resource := range template.Resources {
		result := &AWSLogsMetricFilter{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSLogsMetricFilterWithName retrieves all AWSLogsMetricFilter items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSLogsMetricFilterWithName(name string, template *Template) (*AWSLogsMetricFilter, error) {

	result := &AWSLogsMetricFilter{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSLogsMetricFilter{}, errors.New("resource not found")

}
