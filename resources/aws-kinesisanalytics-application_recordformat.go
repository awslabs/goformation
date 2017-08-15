package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::KinesisAnalytics::Application.RecordFormat AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-recordformat.html
type AWSKinesisAnalyticsApplication_RecordFormat struct {

	// MappingParameters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-recordformat.html#cfn-kinesisanalytics-application-recordformat-mappingparameters
	MappingParameters AWSKinesisAnalyticsApplication_MappingParameters `json:"MappingParameters"`

	// RecordFormatType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-recordformat.html#cfn-kinesisanalytics-application-recordformat-recordformattype
	RecordFormatType string `json:"RecordFormatType"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplication_RecordFormat) AWSCloudFormationType() string {
	return "AWS::KinesisAnalytics::Application.RecordFormat"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisAnalyticsApplication_RecordFormat) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSKinesisAnalyticsApplication_RecordFormatResources retrieves all AWSKinesisAnalyticsApplication_RecordFormat items from a CloudFormation template
func GetAllAWSKinesisAnalyticsApplication_RecordFormat(template *Template) map[string]*AWSKinesisAnalyticsApplication_RecordFormat {

	results := map[string]*AWSKinesisAnalyticsApplication_RecordFormat{}
	for name, resource := range template.Resources {
		result := &AWSKinesisAnalyticsApplication_RecordFormat{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSKinesisAnalyticsApplication_RecordFormatWithName retrieves all AWSKinesisAnalyticsApplication_RecordFormat items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSKinesisAnalyticsApplication_RecordFormat(name string, template *Template) (*AWSKinesisAnalyticsApplication_RecordFormat, error) {

	result := &AWSKinesisAnalyticsApplication_RecordFormat{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSKinesisAnalyticsApplication_RecordFormat{}, errors.New("resource not found")

}
