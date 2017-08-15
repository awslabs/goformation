package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::DMS::Endpoint AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html
type AWSDMSEndpoint struct {

	// CertificateArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html#cfn-dms-endpoint-certificatearn
	CertificateArn string `json:"CertificateArn"`

	// DatabaseName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html#cfn-dms-endpoint-databasename
	DatabaseName string `json:"DatabaseName"`

	// DynamoDbSettings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html#cfn-dms-endpoint-dynamodbsettings
	DynamoDbSettings AWSDMSEndpoint_DynamoDbSettings `json:"DynamoDbSettings"`

	// EndpointIdentifier AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html#cfn-dms-endpoint-endpointidentifier
	EndpointIdentifier string `json:"EndpointIdentifier"`

	// EndpointType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html#cfn-dms-endpoint-endpointtype
	EndpointType string `json:"EndpointType"`

	// EngineName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html#cfn-dms-endpoint-enginename
	EngineName string `json:"EngineName"`

	// ExtraConnectionAttributes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html#cfn-dms-endpoint-extraconnectionattributes
	ExtraConnectionAttributes string `json:"ExtraConnectionAttributes"`

	// KmsKeyId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html#cfn-dms-endpoint-kmskeyid
	KmsKeyId string `json:"KmsKeyId"`

	// MongoDbSettings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html#cfn-dms-endpoint-mongodbsettings
	MongoDbSettings AWSDMSEndpoint_MongoDbSettings `json:"MongoDbSettings"`

	// Password AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html#cfn-dms-endpoint-password
	Password string `json:"Password"`

	// Port AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html#cfn-dms-endpoint-port
	Port int64 `json:"Port"`

	// S3Settings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html#cfn-dms-endpoint-s3settings
	S3Settings AWSDMSEndpoint_S3Settings `json:"S3Settings"`

	// ServerName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html#cfn-dms-endpoint-servername
	ServerName string `json:"ServerName"`

	// SslMode AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html#cfn-dms-endpoint-sslmode
	SslMode string `json:"SslMode"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html#cfn-dms-endpoint-tags
	Tags []Tag `json:"Tags"`

	// Username AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-endpoint.html#cfn-dms-endpoint-username
	Username string `json:"Username"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDMSEndpoint) AWSCloudFormationType() string {
	return "AWS::DMS::Endpoint"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSDMSEndpoint) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSDMSEndpointResources retrieves all AWSDMSEndpoint items from a CloudFormation template
func GetAllAWSDMSEndpoint(template *Template) map[string]*AWSDMSEndpoint {

	results := map[string]*AWSDMSEndpoint{}
	for name, resource := range template.Resources {
		result := &AWSDMSEndpoint{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSDMSEndpointWithName retrieves all AWSDMSEndpoint items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSDMSEndpoint(name string, template *Template) (*AWSDMSEndpoint, error) {

	result := &AWSDMSEndpoint{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSDMSEndpoint{}, errors.New("resource not found")

}
