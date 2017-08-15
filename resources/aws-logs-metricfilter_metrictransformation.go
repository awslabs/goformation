package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Logs::MetricFilter.MetricTransformation AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-metricfilter-metrictransformation.html
type AWSLogsMetricFilter_MetricTransformation struct {

	// MetricName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-metricfilter-metrictransformation.html#cfn-cwl-metricfilter-metrictransformation-metricname
	MetricName string `json:"MetricName"`

	// MetricNamespace AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-metricfilter-metrictransformation.html#cfn-cwl-metricfilter-metrictransformation-metricnamespace
	MetricNamespace string `json:"MetricNamespace"`

	// MetricValue AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-logs-metricfilter-metrictransformation.html#cfn-cwl-metricfilter-metrictransformation-metricvalue
	MetricValue string `json:"MetricValue"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSLogsMetricFilter_MetricTransformation) AWSCloudFormationType() string {
	return "AWS::Logs::MetricFilter.MetricTransformation"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSLogsMetricFilter_MetricTransformation) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSLogsMetricFilter_MetricTransformationResources retrieves all AWSLogsMetricFilter_MetricTransformation items from a CloudFormation template
func GetAllAWSLogsMetricFilter_MetricTransformation(template *Template) map[string]*AWSLogsMetricFilter_MetricTransformation {

	results := map[string]*AWSLogsMetricFilter_MetricTransformation{}
	for name, resource := range template.Resources {
		result := &AWSLogsMetricFilter_MetricTransformation{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSLogsMetricFilter_MetricTransformationWithName retrieves all AWSLogsMetricFilter_MetricTransformation items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSLogsMetricFilter_MetricTransformation(name string, template *Template) (*AWSLogsMetricFilter_MetricTransformation, error) {

	result := &AWSLogsMetricFilter_MetricTransformation{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSLogsMetricFilter_MetricTransformation{}, errors.New("resource not found")

}
