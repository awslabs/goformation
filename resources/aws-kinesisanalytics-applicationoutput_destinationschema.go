package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::KinesisAnalytics::ApplicationOutput.DestinationSchema AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationoutput-destinationschema.html
type AWSKinesisAnalyticsApplicationOutput_DestinationSchema struct {

	// RecordFormatType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kinesisanalytics-applicationoutput-destinationschema.html#cfn-kinesisanalytics-applicationoutput-destinationschema-recordformattype

	RecordFormatType string `json:"RecordFormatType"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKinesisAnalyticsApplicationOutput_DestinationSchema) AWSCloudFormationType() string {
	return "AWS::KinesisAnalytics::ApplicationOutput.DestinationSchema"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKinesisAnalyticsApplicationOutput_DestinationSchema) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSKinesisAnalyticsApplicationOutput_DestinationSchemaResources retrieves all AWSKinesisAnalyticsApplicationOutput_DestinationSchema items from a CloudFormation template
func GetAllAWSKinesisAnalyticsApplicationOutput_DestinationSchema(template *Template) map[string]*AWSKinesisAnalyticsApplicationOutput_DestinationSchema {

	results := map[string]*AWSKinesisAnalyticsApplicationOutput_DestinationSchema{}
	for name, resource := range template.Resources {
		result := &AWSKinesisAnalyticsApplicationOutput_DestinationSchema{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSKinesisAnalyticsApplicationOutput_DestinationSchemaWithName retrieves all AWSKinesisAnalyticsApplicationOutput_DestinationSchema items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSKinesisAnalyticsApplicationOutput_DestinationSchema(name string, template *Template) (*AWSKinesisAnalyticsApplicationOutput_DestinationSchema, error) {

	result := &AWSKinesisAnalyticsApplicationOutput_DestinationSchema{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSKinesisAnalyticsApplicationOutput_DestinationSchema{}, errors.New("resource not found")

}
