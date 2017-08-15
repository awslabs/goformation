package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::KinesisAnalytics::Application.KinesisStreamsInput AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-kinesisstreamsinput.html
type AWSKinesisAnalyticsApplication_KinesisStreamsInput struct {

	// ResourceARN AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-kinesisstreamsinput.html#cfn-kinesisanalytics-application-kinesisstreamsinput-resourcearn

	ResourceARN string `json:"ResourceARN"`

	// RoleARN AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-kinesisstreamsinput.html#cfn-kinesisanalytics-application-kinesisstreamsinput-rolearn

	RoleARN string `json:"RoleARN"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplication_KinesisStreamsInput) AWSCloudFormationType() string {
	return "AWS::KinesisAnalytics::Application.KinesisStreamsInput"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisAnalyticsApplication_KinesisStreamsInput) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSKinesisAnalyticsApplication_KinesisStreamsInputResources retrieves all AWSKinesisAnalyticsApplication_KinesisStreamsInput items from a CloudFormation template
func GetAllAWSKinesisAnalyticsApplication_KinesisStreamsInput(template *Template) map[string]*AWSKinesisAnalyticsApplication_KinesisStreamsInput {

	results := map[string]*AWSKinesisAnalyticsApplication_KinesisStreamsInput{}
	for name, resource := range template.Resources {
		result := &AWSKinesisAnalyticsApplication_KinesisStreamsInput{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSKinesisAnalyticsApplication_KinesisStreamsInputWithName retrieves all AWSKinesisAnalyticsApplication_KinesisStreamsInput items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSKinesisAnalyticsApplication_KinesisStreamsInput(name string, template *Template) (*AWSKinesisAnalyticsApplication_KinesisStreamsInput, error) {

	result := &AWSKinesisAnalyticsApplication_KinesisStreamsInput{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSKinesisAnalyticsApplication_KinesisStreamsInput{}, errors.New("resource not found")

}
