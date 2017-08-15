package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::DMS::Endpoint.S3Settings AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html
type AWSDMSEndpoint_S3Settings struct {

	// BucketFolder AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-bucketfolder
	BucketFolder string `json:"BucketFolder"`

	// BucketName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-bucketname
	BucketName string `json:"BucketName"`

	// CompressionType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-compressiontype
	CompressionType string `json:"CompressionType"`

	// CsvDelimiter AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-csvdelimiter
	CsvDelimiter string `json:"CsvDelimiter"`

	// CsvRowDelimiter AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-csvrowdelimiter
	CsvRowDelimiter string `json:"CsvRowDelimiter"`

	// ExternalTableDefinition AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-externaltabledefinition
	ExternalTableDefinition string `json:"ExternalTableDefinition"`

	// ServiceAccessRoleArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-s3settings.html#cfn-dms-endpoint-s3settings-serviceaccessrolearn
	ServiceAccessRoleArn string `json:"ServiceAccessRoleArn"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDMSEndpoint_S3Settings) AWSCloudFormationType() string {
	return "AWS::DMS::Endpoint.S3Settings"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSDMSEndpoint_S3Settings) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSDMSEndpoint_S3SettingsResources retrieves all AWSDMSEndpoint_S3Settings items from a CloudFormation template
func GetAllAWSDMSEndpoint_S3Settings(template *Template) map[string]*AWSDMSEndpoint_S3Settings {

	results := map[string]*AWSDMSEndpoint_S3Settings{}
	for name, resource := range template.Resources {
		result := &AWSDMSEndpoint_S3Settings{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSDMSEndpoint_S3SettingsWithName retrieves all AWSDMSEndpoint_S3Settings items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSDMSEndpoint_S3Settings(name string, template *Template) (*AWSDMSEndpoint_S3Settings, error) {

	result := &AWSDMSEndpoint_S3Settings{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSDMSEndpoint_S3Settings{}, errors.New("resource not found")

}
