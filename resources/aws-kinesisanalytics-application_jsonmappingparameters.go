package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::KinesisAnalytics::Application.JSONMappingParameters AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-jsonmappingparameters.html
type AWSKinesisAnalyticsApplication_JSONMappingParameters struct {

	// RecordRowPath AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-jsonmappingparameters.html#cfn-kinesisanalytics-application-jsonmappingparameters-recordrowpath
	RecordRowPath string `json:"RecordRowPath"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplication_JSONMappingParameters) AWSCloudFormationType() string {
	return "AWS::KinesisAnalytics::Application.JSONMappingParameters"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisAnalyticsApplication_JSONMappingParameters) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSKinesisAnalyticsApplication_JSONMappingParametersResources retrieves all AWSKinesisAnalyticsApplication_JSONMappingParameters items from a CloudFormation template
func GetAllAWSKinesisAnalyticsApplication_JSONMappingParameters(template *Template) map[string]*AWSKinesisAnalyticsApplication_JSONMappingParameters {

	results := map[string]*AWSKinesisAnalyticsApplication_JSONMappingParameters{}
	for name, resource := range template.Resources {
		result := &AWSKinesisAnalyticsApplication_JSONMappingParameters{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSKinesisAnalyticsApplication_JSONMappingParametersWithName retrieves all AWSKinesisAnalyticsApplication_JSONMappingParameters items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSKinesisAnalyticsApplication_JSONMappingParameters(name string, template *Template) (*AWSKinesisAnalyticsApplication_JSONMappingParameters, error) {

	result := &AWSKinesisAnalyticsApplication_JSONMappingParameters{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSKinesisAnalyticsApplication_JSONMappingParameters{}, errors.New("resource not found")

}
