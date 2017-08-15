package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ApiGateway::Method.MethodResponse AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-methodresponse.html
type AWSApiGatewayMethod_MethodResponse struct {

	// ResponseModels AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-methodresponse.html#cfn-apigateway-method-methodresponse-responsemodels
	ResponseModels map[string]string `json:"ResponseModels"`

	// ResponseParameters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-methodresponse.html#cfn-apigateway-method-methodresponse-responseparameters
	ResponseParameters map[string]bool `json:"ResponseParameters"`

	// StatusCode AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-method-methodresponse.html#cfn-apigateway-method-methodresponse-statuscode
	StatusCode string `json:"StatusCode"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApiGatewayMethod_MethodResponse) AWSCloudFormationType() string {
	return "AWS::ApiGateway::Method.MethodResponse"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSApiGatewayMethod_MethodResponse) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSApiGatewayMethod_MethodResponseResources retrieves all AWSApiGatewayMethod_MethodResponse items from a CloudFormation template
func GetAllAWSApiGatewayMethod_MethodResponse(template *Template) map[string]*AWSApiGatewayMethod_MethodResponse {

	results := map[string]*AWSApiGatewayMethod_MethodResponse{}
	for name, resource := range template.Resources {
		result := &AWSApiGatewayMethod_MethodResponse{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSApiGatewayMethod_MethodResponseWithName retrieves all AWSApiGatewayMethod_MethodResponse items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSApiGatewayMethod_MethodResponse(name string, template *Template) (*AWSApiGatewayMethod_MethodResponse, error) {

	result := &AWSApiGatewayMethod_MethodResponse{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSApiGatewayMethod_MethodResponse{}, errors.New("resource not found")

}
