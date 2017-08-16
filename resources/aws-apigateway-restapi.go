package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSApiGatewayRestApi AWS CloudFormation Resource (AWS::ApiGateway::RestApi)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html
type AWSApiGatewayRestApi struct {

	// BinaryMediaTypes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html#cfn-apigateway-restapi-binarymediatypes
	BinaryMediaTypes []string `json:"BinaryMediaTypes"`

	// Body AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html#cfn-apigateway-restapi-body
	Body interface{} `json:"Body"`

	// BodyS3Location AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html#cfn-apigateway-restapi-bodys3location
	BodyS3Location AWSApiGatewayRestApi_S3Location `json:"BodyS3Location"`

	// CloneFrom AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html#cfn-apigateway-restapi-clonefrom
	CloneFrom string `json:"CloneFrom"`

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html#cfn-apigateway-restapi-description
	Description string `json:"Description"`

	// FailOnWarnings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html#cfn-apigateway-restapi-failonwarnings
	FailOnWarnings bool `json:"FailOnWarnings"`

	// Mode AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html#cfn-apigateway-restapi-mode
	Mode string `json:"Mode"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html#cfn-apigateway-restapi-name
	Name string `json:"Name"`

	// Parameters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-restapi.html#cfn-apigateway-restapi-parameters
	Parameters map[string]string `json:"Parameters"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApiGatewayRestApi) AWSCloudFormationType() string {
	return "AWS::ApiGateway::RestApi"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSApiGatewayRestApi) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSApiGatewayRestApiResources retrieves all AWSApiGatewayRestApi items from a CloudFormation template
func (t *Template) GetAllAWSApiGatewayRestApiResources() map[string]*AWSApiGatewayRestApi {

	results := map[string]*AWSApiGatewayRestApi{}
	for name, resource := range t.Resources {
		result := &AWSApiGatewayRestApi{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSApiGatewayRestApiWithName retrieves all AWSApiGatewayRestApi items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayRestApiWithName(name string) (*AWSApiGatewayRestApi, error) {

	result := &AWSApiGatewayRestApi{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSApiGatewayRestApi{}, errors.New("resource not found")

}
