package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::KinesisAnalytics::ApplicationOutput.KinesisFirehoseOutput AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationoutput-kinesisfirehoseoutput.html
type AWSKinesisAnalyticsApplicationOutput_KinesisFirehoseOutput struct {

	// ResourceARN AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationoutput-kinesisfirehoseoutput.html#cfn-kinesisanalytics-applicationoutput-kinesisfirehoseoutput-resourcearn
	ResourceARN string `json:"ResourceARN"`

	// RoleARN AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationoutput-kinesisfirehoseoutput.html#cfn-kinesisanalytics-applicationoutput-kinesisfirehoseoutput-rolearn
	RoleARN string `json:"RoleARN"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplicationOutput_KinesisFirehoseOutput) AWSCloudFormationType() string {
	return "AWS::KinesisAnalytics::ApplicationOutput.KinesisFirehoseOutput"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisAnalyticsApplicationOutput_KinesisFirehoseOutput) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSKinesisAnalyticsApplicationOutput_KinesisFirehoseOutputResources retrieves all AWSKinesisAnalyticsApplicationOutput_KinesisFirehoseOutput items from a CloudFormation template
func GetAllAWSKinesisAnalyticsApplicationOutput_KinesisFirehoseOutput(template *Template) map[string]*AWSKinesisAnalyticsApplicationOutput_KinesisFirehoseOutput {

	results := map[string]*AWSKinesisAnalyticsApplicationOutput_KinesisFirehoseOutput{}
	for name, resource := range template.Resources {
		result := &AWSKinesisAnalyticsApplicationOutput_KinesisFirehoseOutput{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSKinesisAnalyticsApplicationOutput_KinesisFirehoseOutputWithName retrieves all AWSKinesisAnalyticsApplicationOutput_KinesisFirehoseOutput items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSKinesisAnalyticsApplicationOutput_KinesisFirehoseOutput(name string, template *Template) (*AWSKinesisAnalyticsApplicationOutput_KinesisFirehoseOutput, error) {

	result := &AWSKinesisAnalyticsApplicationOutput_KinesisFirehoseOutput{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSKinesisAnalyticsApplicationOutput_KinesisFirehoseOutput{}, errors.New("resource not found")

}
