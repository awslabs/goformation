package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Logs::SubscriptionFilter AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-subscriptionfilter.html
type AWSLogsSubscriptionFilter struct {

	// DestinationArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-subscriptionfilter.html#cfn-cwl-subscriptionfilter-destinationarn
	DestinationArn string `json:"DestinationArn"`

	// FilterPattern AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-subscriptionfilter.html#cfn-cwl-subscriptionfilter-filterpattern
	FilterPattern string `json:"FilterPattern"`

	// LogGroupName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-subscriptionfilter.html#cfn-cwl-subscriptionfilter-loggroupname
	LogGroupName string `json:"LogGroupName"`

	// RoleArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-subscriptionfilter.html#cfn-cwl-subscriptionfilter-rolearn
	RoleArn string `json:"RoleArn"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSLogsSubscriptionFilter) AWSCloudFormationType() string {
	return "AWS::Logs::SubscriptionFilter"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSLogsSubscriptionFilter) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSLogsSubscriptionFilterResources retrieves all AWSLogsSubscriptionFilter items from a CloudFormation template
func GetAllAWSLogsSubscriptionFilterResources(template *Template) map[string]*AWSLogsSubscriptionFilter {

	results := map[string]*AWSLogsSubscriptionFilter{}
	for name, resource := range template.Resources {
		result := &AWSLogsSubscriptionFilter{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSLogsSubscriptionFilterWithName retrieves all AWSLogsSubscriptionFilter items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSLogsSubscriptionFilterWithName(name string, template *Template) (*AWSLogsSubscriptionFilter, error) {

	result := &AWSLogsSubscriptionFilter{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSLogsSubscriptionFilter{}, errors.New("resource not found")

}
