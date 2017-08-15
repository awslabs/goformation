package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ApiGateway::RestApi.S3Location AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-restapi-bodys3location.html
type AWSApiGatewayRestApi_S3Location struct {

	// Bucket AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-restapi-bodys3location.html#cfn-apigateway-restapi-s3location-bucket
	Bucket string `json:"Bucket"`

	// ETag AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-restapi-bodys3location.html#cfn-apigateway-restapi-s3location-etag
	ETag string `json:"ETag"`

	// Key AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-restapi-bodys3location.html#cfn-apigateway-restapi-s3location-key
	Key string `json:"Key"`

	// Version AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-restapi-bodys3location.html#cfn-apigateway-restapi-s3location-version
	Version string `json:"Version"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApiGatewayRestApi_S3Location) AWSCloudFormationType() string {
	return "AWS::ApiGateway::RestApi.S3Location"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSApiGatewayRestApi_S3Location) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSApiGatewayRestApi_S3LocationResources retrieves all AWSApiGatewayRestApi_S3Location items from a CloudFormation template
func GetAllAWSApiGatewayRestApi_S3Location(template *Template) map[string]*AWSApiGatewayRestApi_S3Location {

	results := map[string]*AWSApiGatewayRestApi_S3Location{}
	for name, resource := range template.Resources {
		result := &AWSApiGatewayRestApi_S3Location{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSApiGatewayRestApi_S3LocationWithName retrieves all AWSApiGatewayRestApi_S3Location items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSApiGatewayRestApi_S3Location(name string, template *Template) (*AWSApiGatewayRestApi_S3Location, error) {

	result := &AWSApiGatewayRestApi_S3Location{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSApiGatewayRestApi_S3Location{}, errors.New("resource not found")

}
