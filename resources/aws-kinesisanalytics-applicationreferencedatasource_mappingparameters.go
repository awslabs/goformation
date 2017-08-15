package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::KinesisAnalytics::ApplicationReferenceDataSource.MappingParameters AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-mappingparameters.html
type AWSKinesisAnalyticsApplicationReferenceDataSource_MappingParameters struct {

	// CSVMappingParameters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-mappingparameters.html#cfn-kinesisanalytics-applicationreferencedatasource-mappingparameters-csvmappingparameters
	CSVMappingParameters AWSKinesisAnalyticsApplicationReferenceDataSource_CSVMappingParameters `json:"CSVMappingParameters"`

	// JSONMappingParameters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-mappingparameters.html#cfn-kinesisanalytics-applicationreferencedatasource-mappingparameters-jsonmappingparameters
	JSONMappingParameters AWSKinesisAnalyticsApplicationReferenceDataSource_JSONMappingParameters `json:"JSONMappingParameters"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplicationReferenceDataSource_MappingParameters) AWSCloudFormationType() string {
	return "AWS::KinesisAnalytics::ApplicationReferenceDataSource.MappingParameters"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisAnalyticsApplicationReferenceDataSource_MappingParameters) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSKinesisAnalyticsApplicationReferenceDataSource_MappingParametersResources retrieves all AWSKinesisAnalyticsApplicationReferenceDataSource_MappingParameters items from a CloudFormation template
func GetAllAWSKinesisAnalyticsApplicationReferenceDataSource_MappingParameters(template *Template) map[string]*AWSKinesisAnalyticsApplicationReferenceDataSource_MappingParameters {

	results := map[string]*AWSKinesisAnalyticsApplicationReferenceDataSource_MappingParameters{}
	for name, resource := range template.Resources {
		result := &AWSKinesisAnalyticsApplicationReferenceDataSource_MappingParameters{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSKinesisAnalyticsApplicationReferenceDataSource_MappingParametersWithName retrieves all AWSKinesisAnalyticsApplicationReferenceDataSource_MappingParameters items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSKinesisAnalyticsApplicationReferenceDataSource_MappingParameters(name string, template *Template) (*AWSKinesisAnalyticsApplicationReferenceDataSource_MappingParameters, error) {

	result := &AWSKinesisAnalyticsApplicationReferenceDataSource_MappingParameters{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSKinesisAnalyticsApplicationReferenceDataSource_MappingParameters{}, errors.New("resource not found")

}
