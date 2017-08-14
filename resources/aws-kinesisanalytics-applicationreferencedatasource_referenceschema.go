package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::KinesisAnalytics::ApplicationReferenceDataSource.ReferenceSchema AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-referenceschema.html
type AWSKinesisAnalyticsApplicationReferenceDataSource_ReferenceSchema struct {

	// RecordColumns AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-referenceschema.html#cfn-kinesisanalytics-applicationreferencedatasource-referenceschema-recordcolumns

	RecordColumns []AWSKinesisAnalyticsApplicationReferenceDataSource_RecordColumn `json:"RecordColumns"`

	// RecordEncoding AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-referenceschema.html#cfn-kinesisanalytics-applicationreferencedatasource-referenceschema-recordencoding

	RecordEncoding string `json:"RecordEncoding"`

	// RecordFormat AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-referenceschema.html#cfn-kinesisanalytics-applicationreferencedatasource-referenceschema-recordformat

	RecordFormat AWSKinesisAnalyticsApplicationReferenceDataSource_RecordFormat `json:"RecordFormat"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplicationReferenceDataSource_ReferenceSchema) AWSCloudFormationType() string {
	return "AWS::KinesisAnalytics::ApplicationReferenceDataSource.ReferenceSchema"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisAnalyticsApplicationReferenceDataSource_ReferenceSchema) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSKinesisAnalyticsApplicationReferenceDataSource_ReferenceSchemaResources retrieves all AWSKinesisAnalyticsApplicationReferenceDataSource_ReferenceSchema items from a CloudFormation template
func GetAllAWSKinesisAnalyticsApplicationReferenceDataSource_ReferenceSchema(template *Template) map[string]*AWSKinesisAnalyticsApplicationReferenceDataSource_ReferenceSchema {

	results := map[string]*AWSKinesisAnalyticsApplicationReferenceDataSource_ReferenceSchema{}
	for name, resource := range template.Resources {
		result := &AWSKinesisAnalyticsApplicationReferenceDataSource_ReferenceSchema{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSKinesisAnalyticsApplicationReferenceDataSource_ReferenceSchemaWithName retrieves all AWSKinesisAnalyticsApplicationReferenceDataSource_ReferenceSchema items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSKinesisAnalyticsApplicationReferenceDataSource_ReferenceSchema(name string, template *Template) (*AWSKinesisAnalyticsApplicationReferenceDataSource_ReferenceSchema, error) {

	result := &AWSKinesisAnalyticsApplicationReferenceDataSource_ReferenceSchema{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSKinesisAnalyticsApplicationReferenceDataSource_ReferenceSchema{}, errors.New("resource not found")

}
