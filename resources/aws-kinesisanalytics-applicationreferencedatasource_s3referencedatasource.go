package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::KinesisAnalytics::ApplicationReferenceDataSource.S3ReferenceDataSource AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-s3referencedatasource.html
type AWSKinesisAnalyticsApplicationReferenceDataSource_S3ReferenceDataSource struct {

	// BucketARN AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-s3referencedatasource.html#cfn-kinesisanalytics-applicationreferencedatasource-s3referencedatasource-bucketarn
	BucketARN string `json:"BucketARN"`

	// FileKey AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-s3referencedatasource.html#cfn-kinesisanalytics-applicationreferencedatasource-s3referencedatasource-filekey
	FileKey string `json:"FileKey"`

	// ReferenceRoleARN AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-s3referencedatasource.html#cfn-kinesisanalytics-applicationreferencedatasource-s3referencedatasource-referencerolearn
	ReferenceRoleARN string `json:"ReferenceRoleARN"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplicationReferenceDataSource_S3ReferenceDataSource) AWSCloudFormationType() string {
	return "AWS::KinesisAnalytics::ApplicationReferenceDataSource.S3ReferenceDataSource"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisAnalyticsApplicationReferenceDataSource_S3ReferenceDataSource) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSKinesisAnalyticsApplicationReferenceDataSource_S3ReferenceDataSourceResources retrieves all AWSKinesisAnalyticsApplicationReferenceDataSource_S3ReferenceDataSource items from a CloudFormation template
func GetAllAWSKinesisAnalyticsApplicationReferenceDataSource_S3ReferenceDataSource(template *Template) map[string]*AWSKinesisAnalyticsApplicationReferenceDataSource_S3ReferenceDataSource {

	results := map[string]*AWSKinesisAnalyticsApplicationReferenceDataSource_S3ReferenceDataSource{}
	for name, resource := range template.Resources {
		result := &AWSKinesisAnalyticsApplicationReferenceDataSource_S3ReferenceDataSource{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSKinesisAnalyticsApplicationReferenceDataSource_S3ReferenceDataSourceWithName retrieves all AWSKinesisAnalyticsApplicationReferenceDataSource_S3ReferenceDataSource items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSKinesisAnalyticsApplicationReferenceDataSource_S3ReferenceDataSource(name string, template *Template) (*AWSKinesisAnalyticsApplicationReferenceDataSource_S3ReferenceDataSource, error) {

	result := &AWSKinesisAnalyticsApplicationReferenceDataSource_S3ReferenceDataSource{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSKinesisAnalyticsApplicationReferenceDataSource_S3ReferenceDataSource{}, errors.New("resource not found")

}
