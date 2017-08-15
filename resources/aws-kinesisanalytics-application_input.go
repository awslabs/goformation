package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::KinesisAnalytics::Application.Input AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-input.html
type AWSKinesisAnalyticsApplication_Input struct {

	// InputParallelism AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-input.html#cfn-kinesisanalytics-application-input-inputparallelism

	InputParallelism AWSKinesisAnalyticsApplication_InputParallelism `json:"InputParallelism"`

	// InputSchema AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-input.html#cfn-kinesisanalytics-application-input-inputschema

	InputSchema AWSKinesisAnalyticsApplication_InputSchema `json:"InputSchema"`

	// KinesisFirehoseInput AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-input.html#cfn-kinesisanalytics-application-input-kinesisfirehoseinput

	KinesisFirehoseInput AWSKinesisAnalyticsApplication_KinesisFirehoseInput `json:"KinesisFirehoseInput"`

	// KinesisStreamsInput AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-input.html#cfn-kinesisanalytics-application-input-kinesisstreamsinput

	KinesisStreamsInput AWSKinesisAnalyticsApplication_KinesisStreamsInput `json:"KinesisStreamsInput"`

	// NamePrefix AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-input.html#cfn-kinesisanalytics-application-input-nameprefix

	NamePrefix string `json:"NamePrefix"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplication_Input) AWSCloudFormationType() string {
	return "AWS::KinesisAnalytics::Application.Input"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisAnalyticsApplication_Input) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSKinesisAnalyticsApplication_InputResources retrieves all AWSKinesisAnalyticsApplication_Input items from a CloudFormation template
func GetAllAWSKinesisAnalyticsApplication_Input(template *Template) map[string]*AWSKinesisAnalyticsApplication_Input {

	results := map[string]*AWSKinesisAnalyticsApplication_Input{}
	for name, resource := range template.Resources {
		result := &AWSKinesisAnalyticsApplication_Input{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSKinesisAnalyticsApplication_InputWithName retrieves all AWSKinesisAnalyticsApplication_Input items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSKinesisAnalyticsApplication_Input(name string, template *Template) (*AWSKinesisAnalyticsApplication_Input, error) {

	result := &AWSKinesisAnalyticsApplication_Input{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSKinesisAnalyticsApplication_Input{}, errors.New("resource not found")

}
