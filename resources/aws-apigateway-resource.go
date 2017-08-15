package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSApiGatewayResource AWS CloudFormation Resource (AWS::ApiGateway::Resource)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-resource.html
type AWSApiGatewayResource struct {

	// ParentId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-resource.html#cfn-apigateway-resource-parentid
	ParentId string `json:"ParentId"`

	// PathPart AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-resource.html#cfn-apigateway-resource-pathpart
	PathPart string `json:"PathPart"`

	// RestApiId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-resource.html#cfn-apigateway-resource-restapiid
	RestApiId string `json:"RestApiId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApiGatewayResource) AWSCloudFormationType() string {
	return "AWS::ApiGateway::Resource"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSApiGatewayResource) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSApiGatewayResourceResources retrieves all AWSApiGatewayResource items from a CloudFormation template
func GetAllAWSApiGatewayResourceResources(template *Template) map[string]*AWSApiGatewayResource {

	results := map[string]*AWSApiGatewayResource{}
	for name, resource := range template.Resources {
		result := &AWSApiGatewayResource{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSApiGatewayResourceWithName retrieves all AWSApiGatewayResource items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSApiGatewayResourceWithName(name string, template *Template) (*AWSApiGatewayResource, error) {

	result := &AWSApiGatewayResource{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSApiGatewayResource{}, errors.New("resource not found")

}
