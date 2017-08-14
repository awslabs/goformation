package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::KinesisAnalytics::Application.CSVMappingParameters AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-csvmappingparameters.html
type AWSKinesisAnalyticsApplication_CSVMappingParameters struct {

	// RecordColumnDelimiter AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-csvmappingparameters.html#cfn-kinesisanalytics-application-csvmappingparameters-recordcolumndelimiter

	RecordColumnDelimiter string `json:"RecordColumnDelimiter"`

	// RecordRowDelimiter AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-csvmappingparameters.html#cfn-kinesisanalytics-application-csvmappingparameters-recordrowdelimiter

	RecordRowDelimiter string `json:"RecordRowDelimiter"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplication_CSVMappingParameters) AWSCloudFormationType() string {
	return "AWS::KinesisAnalytics::Application.CSVMappingParameters"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisAnalyticsApplication_CSVMappingParameters) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSKinesisAnalyticsApplication_CSVMappingParametersResources retrieves all AWSKinesisAnalyticsApplication_CSVMappingParameters items from a CloudFormation template
func GetAllAWSKinesisAnalyticsApplication_CSVMappingParameters(template *Template) map[string]*AWSKinesisAnalyticsApplication_CSVMappingParameters {

	results := map[string]*AWSKinesisAnalyticsApplication_CSVMappingParameters{}
	for name, resource := range template.Resources {
		result := &AWSKinesisAnalyticsApplication_CSVMappingParameters{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSKinesisAnalyticsApplication_CSVMappingParametersWithName retrieves all AWSKinesisAnalyticsApplication_CSVMappingParameters items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSKinesisAnalyticsApplication_CSVMappingParameters(name string, template *Template) (*AWSKinesisAnalyticsApplication_CSVMappingParameters, error) {

	result := &AWSKinesisAnalyticsApplication_CSVMappingParameters{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSKinesisAnalyticsApplication_CSVMappingParameters{}, errors.New("resource not found")

}
