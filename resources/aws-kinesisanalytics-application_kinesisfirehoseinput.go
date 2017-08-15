package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::KinesisAnalytics::Application.KinesisFirehoseInput AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-kinesisfirehoseinput.html
type AWSKinesisAnalyticsApplication_KinesisFirehoseInput struct {

	// ResourceARN AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-kinesisfirehoseinput.html#cfn-kinesisanalytics-application-kinesisfirehoseinput-resourcearn
	ResourceARN string `json:"ResourceARN"`

	// RoleARN AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-kinesisfirehoseinput.html#cfn-kinesisanalytics-application-kinesisfirehoseinput-rolearn
	RoleARN string `json:"RoleARN"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplication_KinesisFirehoseInput) AWSCloudFormationType() string {
	return "AWS::KinesisAnalytics::Application.KinesisFirehoseInput"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisAnalyticsApplication_KinesisFirehoseInput) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSKinesisAnalyticsApplication_KinesisFirehoseInputResources retrieves all AWSKinesisAnalyticsApplication_KinesisFirehoseInput items from a CloudFormation template
func GetAllAWSKinesisAnalyticsApplication_KinesisFirehoseInput(template *Template) map[string]*AWSKinesisAnalyticsApplication_KinesisFirehoseInput {

	results := map[string]*AWSKinesisAnalyticsApplication_KinesisFirehoseInput{}
	for name, resource := range template.Resources {
		result := &AWSKinesisAnalyticsApplication_KinesisFirehoseInput{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSKinesisAnalyticsApplication_KinesisFirehoseInputWithName retrieves all AWSKinesisAnalyticsApplication_KinesisFirehoseInput items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSKinesisAnalyticsApplication_KinesisFirehoseInput(name string, template *Template) (*AWSKinesisAnalyticsApplication_KinesisFirehoseInput, error) {

	result := &AWSKinesisAnalyticsApplication_KinesisFirehoseInput{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSKinesisAnalyticsApplication_KinesisFirehoseInput{}, errors.New("resource not found")

}
