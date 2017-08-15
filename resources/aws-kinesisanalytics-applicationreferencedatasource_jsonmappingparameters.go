package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::KinesisAnalytics::ApplicationReferenceDataSource.JSONMappingParameters AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-jsonmappingparameters.html
type AWSKinesisAnalyticsApplicationReferenceDataSource_JSONMappingParameters struct {

	// RecordRowPath AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationreferencedatasource-jsonmappingparameters.html#cfn-kinesisanalytics-applicationreferencedatasource-jsonmappingparameters-recordrowpath

	RecordRowPath string `json:"RecordRowPath"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplicationReferenceDataSource_JSONMappingParameters) AWSCloudFormationType() string {
	return "AWS::KinesisAnalytics::ApplicationReferenceDataSource.JSONMappingParameters"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisAnalyticsApplicationReferenceDataSource_JSONMappingParameters) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSKinesisAnalyticsApplicationReferenceDataSource_JSONMappingParametersResources retrieves all AWSKinesisAnalyticsApplicationReferenceDataSource_JSONMappingParameters items from a CloudFormation template
func GetAllAWSKinesisAnalyticsApplicationReferenceDataSource_JSONMappingParameters(template *Template) map[string]*AWSKinesisAnalyticsApplicationReferenceDataSource_JSONMappingParameters {

	results := map[string]*AWSKinesisAnalyticsApplicationReferenceDataSource_JSONMappingParameters{}
	for name, resource := range template.Resources {
		result := &AWSKinesisAnalyticsApplicationReferenceDataSource_JSONMappingParameters{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSKinesisAnalyticsApplicationReferenceDataSource_JSONMappingParametersWithName retrieves all AWSKinesisAnalyticsApplicationReferenceDataSource_JSONMappingParameters items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSKinesisAnalyticsApplicationReferenceDataSource_JSONMappingParameters(name string, template *Template) (*AWSKinesisAnalyticsApplicationReferenceDataSource_JSONMappingParameters, error) {

	result := &AWSKinesisAnalyticsApplicationReferenceDataSource_JSONMappingParameters{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSKinesisAnalyticsApplicationReferenceDataSource_JSONMappingParameters{}, errors.New("resource not found")

}
