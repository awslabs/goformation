package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ApiGateway::Method.IntegrationResponse AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-integration-integrationresponse.html
type AWSApiGatewayMethod_IntegrationResponse struct {

	// ResponseParameters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-integration-integrationresponse.html#cfn-apigateway-method-integration-integrationresponse-responseparameters

	ResponseParameters map[string]string `json:"ResponseParameters"`

	// ResponseTemplates AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-integration-integrationresponse.html#cfn-apigateway-method-integration-integrationresponse-responsetemplates

	ResponseTemplates map[string]string `json:"ResponseTemplates"`

	// SelectionPattern AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-integration-integrationresponse.html#cfn-apigateway-method-integration-integrationresponse-selectionpattern

	SelectionPattern string `json:"SelectionPattern"`

	// StatusCode AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-integration-integrationresponse.html#cfn-apigateway-method-integration-integrationresponse-statuscode

	StatusCode string `json:"StatusCode"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApiGatewayMethod_IntegrationResponse) AWSCloudFormationType() string {
	return "AWS::ApiGateway::Method.IntegrationResponse"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSApiGatewayMethod_IntegrationResponse) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSApiGatewayMethod_IntegrationResponseResources retrieves all AWSApiGatewayMethod_IntegrationResponse items from a CloudFormation template
func GetAllAWSApiGatewayMethod_IntegrationResponse(template *Template) map[string]*AWSApiGatewayMethod_IntegrationResponse {

	results := map[string]*AWSApiGatewayMethod_IntegrationResponse{}
	for name, resource := range template.Resources {
		result := &AWSApiGatewayMethod_IntegrationResponse{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSApiGatewayMethod_IntegrationResponseWithName retrieves all AWSApiGatewayMethod_IntegrationResponse items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSApiGatewayMethod_IntegrationResponse(name string, template *Template) (*AWSApiGatewayMethod_IntegrationResponse, error) {

	result := &AWSApiGatewayMethod_IntegrationResponse{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSApiGatewayMethod_IntegrationResponse{}, errors.New("resource not found")

}
