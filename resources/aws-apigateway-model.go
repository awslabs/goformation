package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ApiGateway::Model AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-model.html
type AWSApiGatewayModel struct {

	// ContentType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-model.html#cfn-apigateway-model-contenttype
	ContentType string `json:"ContentType"`

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-model.html#cfn-apigateway-model-description
	Description string `json:"Description"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-model.html#cfn-apigateway-model-name
	Name string `json:"Name"`

	// RestApiId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-model.html#cfn-apigateway-model-restapiid
	RestApiId string `json:"RestApiId"`

	// Schema AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-model.html#cfn-apigateway-model-schema
	Schema interface{} `json:"Schema"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApiGatewayModel) AWSCloudFormationType() string {
	return "AWS::ApiGateway::Model"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSApiGatewayModel) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSApiGatewayModelResources retrieves all AWSApiGatewayModel items from a CloudFormation template
func GetAllAWSApiGatewayModelResources(template *Template) map[string]*AWSApiGatewayModel {

	results := map[string]*AWSApiGatewayModel{}
	for name, resource := range template.Resources {
		result := &AWSApiGatewayModel{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSApiGatewayModelWithName retrieves all AWSApiGatewayModel items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSApiGatewayModelWithName(name string, template *Template) (*AWSApiGatewayModel, error) {

	result := &AWSApiGatewayModel{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSApiGatewayModel{}, errors.New("resource not found")

}
