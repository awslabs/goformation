package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::KinesisAnalytics::ApplicationReferenceDataSource.RecordColumn AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-recordcolumn.html
type AWSKinesisAnalyticsApplicationReferenceDataSource_RecordColumn struct {

	// Mapping AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-recordcolumn.html#cfn-kinesisanalytics-applicationreferencedatasource-recordcolumn-mapping
	Mapping string `json:"Mapping"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-recordcolumn.html#cfn-kinesisanalytics-applicationreferencedatasource-recordcolumn-name
	Name string `json:"Name"`

	// SqlType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-recordcolumn.html#cfn-kinesisanalytics-applicationreferencedatasource-recordcolumn-sqltype
	SqlType string `json:"SqlType"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplicationReferenceDataSource_RecordColumn) AWSCloudFormationType() string {
	return "AWS::KinesisAnalytics::ApplicationReferenceDataSource.RecordColumn"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisAnalyticsApplicationReferenceDataSource_RecordColumn) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSKinesisAnalyticsApplicationReferenceDataSource_RecordColumnResources retrieves all AWSKinesisAnalyticsApplicationReferenceDataSource_RecordColumn items from a CloudFormation template
func GetAllAWSKinesisAnalyticsApplicationReferenceDataSource_RecordColumn(template *Template) map[string]*AWSKinesisAnalyticsApplicationReferenceDataSource_RecordColumn {

	results := map[string]*AWSKinesisAnalyticsApplicationReferenceDataSource_RecordColumn{}
	for name, resource := range template.Resources {
		result := &AWSKinesisAnalyticsApplicationReferenceDataSource_RecordColumn{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSKinesisAnalyticsApplicationReferenceDataSource_RecordColumnWithName retrieves all AWSKinesisAnalyticsApplicationReferenceDataSource_RecordColumn items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSKinesisAnalyticsApplicationReferenceDataSource_RecordColumn(name string, template *Template) (*AWSKinesisAnalyticsApplicationReferenceDataSource_RecordColumn, error) {

	result := &AWSKinesisAnalyticsApplicationReferenceDataSource_RecordColumn{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSKinesisAnalyticsApplicationReferenceDataSource_RecordColumn{}, errors.New("resource not found")

}
