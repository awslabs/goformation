package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::KinesisAnalytics::ApplicationReferenceDataSource.CSVMappingParameters AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-csvmappingparameters.html
type AWSKinesisAnalyticsApplicationReferenceDataSource_CSVMappingParameters struct {

	// RecordColumnDelimiter AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-csvmappingparameters.html#cfn-kinesisanalytics-applicationreferencedatasource-csvmappingparameters-recordcolumndelimiter

	RecordColumnDelimiter string `json:"RecordColumnDelimiter"`

	// RecordRowDelimiter AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-csvmappingparameters.html#cfn-kinesisanalytics-applicationreferencedatasource-csvmappingparameters-recordrowdelimiter

	RecordRowDelimiter string `json:"RecordRowDelimiter"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplicationReferenceDataSource_CSVMappingParameters) AWSCloudFormationType() string {
	return "AWS::KinesisAnalytics::ApplicationReferenceDataSource.CSVMappingParameters"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisAnalyticsApplicationReferenceDataSource_CSVMappingParameters) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSKinesisAnalyticsApplicationReferenceDataSource_CSVMappingParametersResources retrieves all AWSKinesisAnalyticsApplicationReferenceDataSource_CSVMappingParameters items from a CloudFormation template
func GetAllAWSKinesisAnalyticsApplicationReferenceDataSource_CSVMappingParameters(template *Template) map[string]*AWSKinesisAnalyticsApplicationReferenceDataSource_CSVMappingParameters {

	results := map[string]*AWSKinesisAnalyticsApplicationReferenceDataSource_CSVMappingParameters{}
	for name, resource := range template.Resources {
		result := &AWSKinesisAnalyticsApplicationReferenceDataSource_CSVMappingParameters{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSKinesisAnalyticsApplicationReferenceDataSource_CSVMappingParametersWithName retrieves all AWSKinesisAnalyticsApplicationReferenceDataSource_CSVMappingParameters items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSKinesisAnalyticsApplicationReferenceDataSource_CSVMappingParameters(name string, template *Template) (*AWSKinesisAnalyticsApplicationReferenceDataSource_CSVMappingParameters, error) {

	result := &AWSKinesisAnalyticsApplicationReferenceDataSource_CSVMappingParameters{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSKinesisAnalyticsApplicationReferenceDataSource_CSVMappingParameters{}, errors.New("resource not found")

}
