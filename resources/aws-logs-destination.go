package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Logs::Destination AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-destination.html
type AWSLogsDestination struct {

	// DestinationName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-destination.html#cfn-logs-destination-destinationname
	DestinationName string `json:"DestinationName"`

	// DestinationPolicy AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-destination.html#cfn-logs-destination-destinationpolicy
	DestinationPolicy string `json:"DestinationPolicy"`

	// RoleArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-destination.html#cfn-logs-destination-rolearn
	RoleArn string `json:"RoleArn"`

	// TargetArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-logs-destination.html#cfn-logs-destination-targetarn
	TargetArn string `json:"TargetArn"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSLogsDestination) AWSCloudFormationType() string {
	return "AWS::Logs::Destination"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSLogsDestination) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSLogsDestinationResources retrieves all AWSLogsDestination items from a CloudFormation template
func GetAllAWSLogsDestination(template *Template) map[string]*AWSLogsDestination {

	results := map[string]*AWSLogsDestination{}
	for name, resource := range template.Resources {
		result := &AWSLogsDestination{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSLogsDestinationWithName retrieves all AWSLogsDestination items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSLogsDestination(name string, template *Template) (*AWSLogsDestination, error) {

	result := &AWSLogsDestination{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSLogsDestination{}, errors.New("resource not found")

}
