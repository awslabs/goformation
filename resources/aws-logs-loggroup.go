package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSLogsLogGroup AWS CloudFormation Resource (AWS::Logs::LogGroup)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loggroup.html
type AWSLogsLogGroup struct {

	// LogGroupName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loggroup.html#cfn-cwl-loggroup-loggroupname
	LogGroupName string `json:"LogGroupName"`

	// RetentionInDays AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-loggroup.html#cfn-cwl-loggroup-retentionindays
	RetentionInDays int `json:"RetentionInDays"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSLogsLogGroup) AWSCloudFormationType() string {
	return "AWS::Logs::LogGroup"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSLogsLogGroup) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSLogsLogGroupResources retrieves all AWSLogsLogGroup items from a CloudFormation template
func (t *Template) GetAllAWSLogsLogGroupResources() map[string]*AWSLogsLogGroup {

	results := map[string]*AWSLogsLogGroup{}
	for name, resource := range t.Resources {
		result := &AWSLogsLogGroup{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSLogsLogGroupWithName retrieves all AWSLogsLogGroup items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSLogsLogGroupWithName(name string) (*AWSLogsLogGroup, error) {

	result := &AWSLogsLogGroup{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSLogsLogGroup{}, errors.New("resource not found")

}
