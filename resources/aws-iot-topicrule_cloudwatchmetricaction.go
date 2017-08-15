package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::IoT::TopicRule.CloudwatchMetricAction AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-cloudwatchmetric.html
type AWSIoTTopicRule_CloudwatchMetricAction struct {

	// MetricName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-cloudwatchmetric.html#cfn-iot-cloudwatchmetric-metricname
	MetricName string `json:"MetricName"`

	// MetricNamespace AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-cloudwatchmetric.html#cfn-iot-cloudwatchmetric-metricnamespace
	MetricNamespace string `json:"MetricNamespace"`

	// MetricTimestamp AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-cloudwatchmetric.html#cfn-iot-cloudwatchmetric-metrictimestamp
	MetricTimestamp string `json:"MetricTimestamp"`

	// MetricUnit AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-cloudwatchmetric.html#cfn-iot-cloudwatchmetric-metricunit
	MetricUnit string `json:"MetricUnit"`

	// MetricValue AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-cloudwatchmetric.html#cfn-iot-cloudwatchmetric-metricvalue
	MetricValue string `json:"MetricValue"`

	// RoleArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-cloudwatchmetric.html#cfn-iot-cloudwatchmetric-rolearn
	RoleArn string `json:"RoleArn"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTTopicRule_CloudwatchMetricAction) AWSCloudFormationType() string {
	return "AWS::IoT::TopicRule.CloudwatchMetricAction"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIoTTopicRule_CloudwatchMetricAction) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSIoTTopicRule_CloudwatchMetricActionResources retrieves all AWSIoTTopicRule_CloudwatchMetricAction items from a CloudFormation template
func GetAllAWSIoTTopicRule_CloudwatchMetricAction(template *Template) map[string]*AWSIoTTopicRule_CloudwatchMetricAction {

	results := map[string]*AWSIoTTopicRule_CloudwatchMetricAction{}
	for name, resource := range template.Resources {
		result := &AWSIoTTopicRule_CloudwatchMetricAction{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSIoTTopicRule_CloudwatchMetricActionWithName retrieves all AWSIoTTopicRule_CloudwatchMetricAction items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSIoTTopicRule_CloudwatchMetricAction(name string, template *Template) (*AWSIoTTopicRule_CloudwatchMetricAction, error) {

	result := &AWSIoTTopicRule_CloudwatchMetricAction{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSIoTTopicRule_CloudwatchMetricAction{}, errors.New("resource not found")

}
