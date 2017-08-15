package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSKinesisAnalyticsApplicationOutput AWS CloudFormation Resource (AWS::KinesisAnalytics::ApplicationOutput)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalytics-applicationoutput.html
type AWSKinesisAnalyticsApplicationOutput struct {

	// ApplicationName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalytics-applicationoutput.html#cfn-kinesisanalytics-applicationoutput-applicationname
	ApplicationName string `json:"ApplicationName"`

	// Output AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalytics-applicationoutput.html#cfn-kinesisanalytics-applicationoutput-output
	Output AWSKinesisAnalyticsApplicationOutput_Output `json:"Output"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplicationOutput) AWSCloudFormationType() string {
	return "AWS::KinesisAnalytics::ApplicationOutput"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisAnalyticsApplicationOutput) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSKinesisAnalyticsApplicationOutputResources retrieves all AWSKinesisAnalyticsApplicationOutput items from a CloudFormation template
func GetAllAWSKinesisAnalyticsApplicationOutputResources(template *Template) map[string]*AWSKinesisAnalyticsApplicationOutput {

	results := map[string]*AWSKinesisAnalyticsApplicationOutput{}
	for name, resource := range template.Resources {
		result := &AWSKinesisAnalyticsApplicationOutput{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSKinesisAnalyticsApplicationOutputWithName retrieves all AWSKinesisAnalyticsApplicationOutput items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSKinesisAnalyticsApplicationOutputWithName(name string, template *Template) (*AWSKinesisAnalyticsApplicationOutput, error) {

	result := &AWSKinesisAnalyticsApplicationOutput{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSKinesisAnalyticsApplicationOutput{}, errors.New("resource not found")

}
