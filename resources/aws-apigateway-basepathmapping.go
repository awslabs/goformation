package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ApiGateway::BasePathMapping AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-basepathmapping.html
type AWSApiGatewayBasePathMapping struct {

	// BasePath AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-basepathmapping.html#cfn-apigateway-basepathmapping-basepath
	BasePath string `json:"BasePath"`

	// DomainName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-basepathmapping.html#cfn-apigateway-basepathmapping-domainname
	DomainName string `json:"DomainName"`

	// RestApiId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-basepathmapping.html#cfn-apigateway-basepathmapping-restapiid
	RestApiId string `json:"RestApiId"`

	// Stage AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-basepathmapping.html#cfn-apigateway-basepathmapping-stage
	Stage string `json:"Stage"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApiGatewayBasePathMapping) AWSCloudFormationType() string {
	return "AWS::ApiGateway::BasePathMapping"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSApiGatewayBasePathMapping) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSApiGatewayBasePathMappingResources retrieves all AWSApiGatewayBasePathMapping items from a CloudFormation template
func GetAllAWSApiGatewayBasePathMappingResources(template *Template) map[string]*AWSApiGatewayBasePathMapping {

	results := map[string]*AWSApiGatewayBasePathMapping{}
	for name, resource := range template.Resources {
		result := &AWSApiGatewayBasePathMapping{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSApiGatewayBasePathMappingWithName retrieves all AWSApiGatewayBasePathMapping items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSApiGatewayBasePathMappingWithName(name string, template *Template) (*AWSApiGatewayBasePathMapping, error) {

	result := &AWSApiGatewayBasePathMapping{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSApiGatewayBasePathMapping{}, errors.New("resource not found")

}
