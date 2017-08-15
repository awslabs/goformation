package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::KinesisAnalytics::Application.MappingParameters AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-mappingparameters.html
type AWSKinesisAnalyticsApplication_MappingParameters struct {

	// CSVMappingParameters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-mappingparameters.html#cfn-kinesisanalytics-application-mappingparameters-csvmappingparameters
	CSVMappingParameters AWSKinesisAnalyticsApplication_CSVMappingParameters `json:"CSVMappingParameters"`

	// JSONMappingParameters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-mappingparameters.html#cfn-kinesisanalytics-application-mappingparameters-jsonmappingparameters
	JSONMappingParameters AWSKinesisAnalyticsApplication_JSONMappingParameters `json:"JSONMappingParameters"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplication_MappingParameters) AWSCloudFormationType() string {
	return "AWS::KinesisAnalytics::Application.MappingParameters"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisAnalyticsApplication_MappingParameters) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSKinesisAnalyticsApplication_MappingParametersResources retrieves all AWSKinesisAnalyticsApplication_MappingParameters items from a CloudFormation template
func GetAllAWSKinesisAnalyticsApplication_MappingParameters(template *Template) map[string]*AWSKinesisAnalyticsApplication_MappingParameters {

	results := map[string]*AWSKinesisAnalyticsApplication_MappingParameters{}
	for name, resource := range template.Resources {
		result := &AWSKinesisAnalyticsApplication_MappingParameters{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSKinesisAnalyticsApplication_MappingParametersWithName retrieves all AWSKinesisAnalyticsApplication_MappingParameters items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSKinesisAnalyticsApplication_MappingParameters(name string, template *Template) (*AWSKinesisAnalyticsApplication_MappingParameters, error) {

	result := &AWSKinesisAnalyticsApplication_MappingParameters{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSKinesisAnalyticsApplication_MappingParameters{}, errors.New("resource not found")

}
