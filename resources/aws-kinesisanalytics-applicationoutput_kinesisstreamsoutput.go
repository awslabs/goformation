package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::KinesisAnalytics::ApplicationOutput.KinesisStreamsOutput AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationoutput-kinesisstreamsoutput.html
type AWSKinesisAnalyticsApplicationOutput_KinesisStreamsOutput struct {

	// ResourceARN AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationoutput-kinesisstreamsoutput.html#cfn-kinesisanalytics-applicationoutput-kinesisstreamsoutput-resourcearn
	ResourceARN string `json:"ResourceARN"`

	// RoleARN AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationoutput-kinesisstreamsoutput.html#cfn-kinesisanalytics-applicationoutput-kinesisstreamsoutput-rolearn
	RoleARN string `json:"RoleARN"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplicationOutput_KinesisStreamsOutput) AWSCloudFormationType() string {
	return "AWS::KinesisAnalytics::ApplicationOutput.KinesisStreamsOutput"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisAnalyticsApplicationOutput_KinesisStreamsOutput) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSKinesisAnalyticsApplicationOutput_KinesisStreamsOutputResources retrieves all AWSKinesisAnalyticsApplicationOutput_KinesisStreamsOutput items from a CloudFormation template
func GetAllAWSKinesisAnalyticsApplicationOutput_KinesisStreamsOutput(template *Template) map[string]*AWSKinesisAnalyticsApplicationOutput_KinesisStreamsOutput {

	results := map[string]*AWSKinesisAnalyticsApplicationOutput_KinesisStreamsOutput{}
	for name, resource := range template.Resources {
		result := &AWSKinesisAnalyticsApplicationOutput_KinesisStreamsOutput{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSKinesisAnalyticsApplicationOutput_KinesisStreamsOutputWithName retrieves all AWSKinesisAnalyticsApplicationOutput_KinesisStreamsOutput items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSKinesisAnalyticsApplicationOutput_KinesisStreamsOutput(name string, template *Template) (*AWSKinesisAnalyticsApplicationOutput_KinesisStreamsOutput, error) {

	result := &AWSKinesisAnalyticsApplicationOutput_KinesisStreamsOutput{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSKinesisAnalyticsApplicationOutput_KinesisStreamsOutput{}, errors.New("resource not found")

}
