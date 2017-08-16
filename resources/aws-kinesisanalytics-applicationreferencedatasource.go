package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSKinesisAnalyticsApplicationReferenceDataSource AWS CloudFormation Resource (AWS::KinesisAnalytics::ApplicationReferenceDataSource)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalytics-applicationreferencedatasource.html
type AWSKinesisAnalyticsApplicationReferenceDataSource struct {

	// ApplicationName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalytics-applicationreferencedatasource.html#cfn-kinesisanalytics-applicationreferencedatasource-applicationname
	ApplicationName string `json:"ApplicationName"`

	// ReferenceDataSource AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kinesisanalytics-applicationreferencedatasource.html#cfn-kinesisanalytics-applicationreferencedatasource-referencedatasource
	ReferenceDataSource AWSKinesisAnalyticsApplicationReferenceDataSource_ReferenceDataSource `json:"ReferenceDataSource"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplicationReferenceDataSource) AWSCloudFormationType() string {
	return "AWS::KinesisAnalytics::ApplicationReferenceDataSource"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisAnalyticsApplicationReferenceDataSource) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSKinesisAnalyticsApplicationReferenceDataSourceResources retrieves all AWSKinesisAnalyticsApplicationReferenceDataSource items from a CloudFormation template
func (t *Template) GetAllAWSKinesisAnalyticsApplicationReferenceDataSourceResources() map[string]*AWSKinesisAnalyticsApplicationReferenceDataSource {

	results := map[string]*AWSKinesisAnalyticsApplicationReferenceDataSource{}
	for name, resource := range t.Resources {
		result := &AWSKinesisAnalyticsApplicationReferenceDataSource{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSKinesisAnalyticsApplicationReferenceDataSourceWithName retrieves all AWSKinesisAnalyticsApplicationReferenceDataSource items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSKinesisAnalyticsApplicationReferenceDataSourceWithName(name string) (*AWSKinesisAnalyticsApplicationReferenceDataSource, error) {

	result := &AWSKinesisAnalyticsApplicationReferenceDataSource{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSKinesisAnalyticsApplicationReferenceDataSource{}, errors.New("resource not found")

}
