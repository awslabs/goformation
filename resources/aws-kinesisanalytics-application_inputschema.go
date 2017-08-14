package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::KinesisAnalytics::Application.InputSchema AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-inputschema.html
type AWSKinesisAnalyticsApplication_InputSchema struct {

	// RecordColumns AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-inputschema.html#cfn-kinesisanalytics-application-inputschema-recordcolumns

	RecordColumns []AWSKinesisAnalyticsApplication_RecordColumn `json:"RecordColumns"`

	// RecordEncoding AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-inputschema.html#cfn-kinesisanalytics-application-inputschema-recordencoding

	RecordEncoding string `json:"RecordEncoding"`

	// RecordFormat AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-application-inputschema.html#cfn-kinesisanalytics-application-inputschema-recordformat

	RecordFormat AWSKinesisAnalyticsApplication_RecordFormat `json:"RecordFormat"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplication_InputSchema) AWSCloudFormationType() string {
	return "AWS::KinesisAnalytics::Application.InputSchema"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisAnalyticsApplication_InputSchema) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSKinesisAnalyticsApplication_InputSchemaResources retrieves all AWSKinesisAnalyticsApplication_InputSchema items from a CloudFormation template
func GetAllAWSKinesisAnalyticsApplication_InputSchema(template *Template) map[string]*AWSKinesisAnalyticsApplication_InputSchema {

	results := map[string]*AWSKinesisAnalyticsApplication_InputSchema{}
	for name, resource := range template.Resources {
		result := &AWSKinesisAnalyticsApplication_InputSchema{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSKinesisAnalyticsApplication_InputSchemaWithName retrieves all AWSKinesisAnalyticsApplication_InputSchema items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSKinesisAnalyticsApplication_InputSchema(name string, template *Template) (*AWSKinesisAnalyticsApplication_InputSchema, error) {

	result := &AWSKinesisAnalyticsApplication_InputSchema{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSKinesisAnalyticsApplication_InputSchema{}, errors.New("resource not found")

}
