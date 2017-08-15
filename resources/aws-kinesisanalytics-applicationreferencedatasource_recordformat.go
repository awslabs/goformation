package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::KinesisAnalytics::ApplicationReferenceDataSource.RecordFormat AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-recordformat.html
type AWSKinesisAnalyticsApplicationReferenceDataSource_RecordFormat struct {

	// MappingParameters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-recordformat.html#cfn-kinesisanalytics-applicationreferencedatasource-recordformat-mappingparameters
	MappingParameters AWSKinesisAnalyticsApplicationReferenceDataSource_MappingParameters `json:"MappingParameters"`

	// RecordFormatType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-recordformat.html#cfn-kinesisanalytics-applicationreferencedatasource-recordformat-recordformattype
	RecordFormatType string `json:"RecordFormatType"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplicationReferenceDataSource_RecordFormat) AWSCloudFormationType() string {
	return "AWS::KinesisAnalytics::ApplicationReferenceDataSource.RecordFormat"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisAnalyticsApplicationReferenceDataSource_RecordFormat) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSKinesisAnalyticsApplicationReferenceDataSource_RecordFormatResources retrieves all AWSKinesisAnalyticsApplicationReferenceDataSource_RecordFormat items from a CloudFormation template
func GetAllAWSKinesisAnalyticsApplicationReferenceDataSource_RecordFormat(template *Template) map[string]*AWSKinesisAnalyticsApplicationReferenceDataSource_RecordFormat {

	results := map[string]*AWSKinesisAnalyticsApplicationReferenceDataSource_RecordFormat{}
	for name, resource := range template.Resources {
		result := &AWSKinesisAnalyticsApplicationReferenceDataSource_RecordFormat{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSKinesisAnalyticsApplicationReferenceDataSource_RecordFormatWithName retrieves all AWSKinesisAnalyticsApplicationReferenceDataSource_RecordFormat items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSKinesisAnalyticsApplicationReferenceDataSource_RecordFormat(name string, template *Template) (*AWSKinesisAnalyticsApplicationReferenceDataSource_RecordFormat, error) {

	result := &AWSKinesisAnalyticsApplicationReferenceDataSource_RecordFormat{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSKinesisAnalyticsApplicationReferenceDataSource_RecordFormat{}, errors.New("resource not found")

}
