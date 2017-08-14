package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::KinesisAnalytics::ApplicationReferenceDataSource.ReferenceDataSource AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-referencedatasource.html
type AWSKinesisAnalyticsApplicationReferenceDataSource_ReferenceDataSource struct {

	// ReferenceSchema AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-referencedatasource.html#cfn-kinesisanalytics-applicationreferencedatasource-referencedatasource-referenceschema

	ReferenceSchema AWSKinesisAnalyticsApplicationReferenceDataSource_ReferenceSchema `json:"ReferenceSchema"`

	// S3ReferenceDataSource AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-referencedatasource.html#cfn-kinesisanalytics-applicationreferencedatasource-referencedatasource-s3referencedatasource

	S3ReferenceDataSource AWSKinesisAnalyticsApplicationReferenceDataSource_S3ReferenceDataSource `json:"S3ReferenceDataSource"`

	// TableName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-referencedatasource.html#cfn-kinesisanalytics-applicationreferencedatasource-referencedatasource-tablename

	TableName string `json:"TableName"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplicationReferenceDataSource_ReferenceDataSource) AWSCloudFormationType() string {
	return "AWS::KinesisAnalytics::ApplicationReferenceDataSource.ReferenceDataSource"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisAnalyticsApplicationReferenceDataSource_ReferenceDataSource) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSKinesisAnalyticsApplicationReferenceDataSource_ReferenceDataSourceResources retrieves all AWSKinesisAnalyticsApplicationReferenceDataSource_ReferenceDataSource items from a CloudFormation template
func GetAllAWSKinesisAnalyticsApplicationReferenceDataSource_ReferenceDataSource(template *Template) map[string]*AWSKinesisAnalyticsApplicationReferenceDataSource_ReferenceDataSource {

	results := map[string]*AWSKinesisAnalyticsApplicationReferenceDataSource_ReferenceDataSource{}
	for name, resource := range template.Resources {
		result := &AWSKinesisAnalyticsApplicationReferenceDataSource_ReferenceDataSource{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSKinesisAnalyticsApplicationReferenceDataSource_ReferenceDataSourceWithName retrieves all AWSKinesisAnalyticsApplicationReferenceDataSource_ReferenceDataSource items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSKinesisAnalyticsApplicationReferenceDataSource_ReferenceDataSource(name string, template *Template) (*AWSKinesisAnalyticsApplicationReferenceDataSource_ReferenceDataSource, error) {

	result := &AWSKinesisAnalyticsApplicationReferenceDataSource_ReferenceDataSource{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSKinesisAnalyticsApplicationReferenceDataSource_ReferenceDataSource{}, errors.New("resource not found")

}
