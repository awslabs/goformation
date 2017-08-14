package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Logs::LogStream AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-logstream.html
type AWSLogsLogStream struct {

	// LogGroupName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-logstream.html#cfn-logs-logstream-loggroupname

	LogGroupName string `json:"LogGroupName"`

	// LogStreamName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-logstream.html#cfn-logs-logstream-logstreamname

	LogStreamName string `json:"LogStreamName"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSLogsLogStream) AWSCloudFormationType() string {
	return "AWS::Logs::LogStream"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSLogsLogStream) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSLogsLogStreamResources retrieves all AWSLogsLogStream items from a CloudFormation template
func GetAllAWSLogsLogStream(template *Template) map[string]*AWSLogsLogStream {

	results := map[string]*AWSLogsLogStream{}
	for name, resource := range template.Resources {
		result := &AWSLogsLogStream{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSLogsLogStreamWithName retrieves all AWSLogsLogStream items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSLogsLogStream(name string, template *Template) (*AWSLogsLogStream, error) {

	result := &AWSLogsLogStream{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSLogsLogStream{}, errors.New("resource not found")

}
