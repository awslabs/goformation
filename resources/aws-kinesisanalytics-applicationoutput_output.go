package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::KinesisAnalytics::ApplicationOutput.Output AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationoutput-output.html
type AWSKinesisAnalyticsApplicationOutput_Output struct {

	// DestinationSchema AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationoutput-output.html#cfn-kinesisanalytics-applicationoutput-output-destinationschema
	DestinationSchema AWSKinesisAnalyticsApplicationOutput_DestinationSchema `json:"DestinationSchema"`

	// KinesisFirehoseOutput AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationoutput-output.html#cfn-kinesisanalytics-applicationoutput-output-kinesisfirehoseoutput
	KinesisFirehoseOutput AWSKinesisAnalyticsApplicationOutput_KinesisFirehoseOutput `json:"KinesisFirehoseOutput"`

	// KinesisStreamsOutput AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationoutput-output.html#cfn-kinesisanalytics-applicationoutput-output-kinesisstreamsoutput
	KinesisStreamsOutput AWSKinesisAnalyticsApplicationOutput_KinesisStreamsOutput `json:"KinesisStreamsOutput"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationoutput-output.html#cfn-kinesisanalytics-applicationoutput-output-name
	Name string `json:"Name"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplicationOutput_Output) AWSCloudFormationType() string {
	return "AWS::KinesisAnalytics::ApplicationOutput.Output"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisAnalyticsApplicationOutput_Output) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSKinesisAnalyticsApplicationOutput_OutputResources retrieves all AWSKinesisAnalyticsApplicationOutput_Output items from a CloudFormation template
func GetAllAWSKinesisAnalyticsApplicationOutput_Output(template *Template) map[string]*AWSKinesisAnalyticsApplicationOutput_Output {

	results := map[string]*AWSKinesisAnalyticsApplicationOutput_Output{}
	for name, resource := range template.Resources {
		result := &AWSKinesisAnalyticsApplicationOutput_Output{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSKinesisAnalyticsApplicationOutput_OutputWithName retrieves all AWSKinesisAnalyticsApplicationOutput_Output items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSKinesisAnalyticsApplicationOutput_Output(name string, template *Template) (*AWSKinesisAnalyticsApplicationOutput_Output, error) {

	result := &AWSKinesisAnalyticsApplicationOutput_Output{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSKinesisAnalyticsApplicationOutput_Output{}, errors.New("resource not found")

}
