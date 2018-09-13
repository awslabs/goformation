package cloudformation

import (
	"encoding/json"
)

// AWSApiGatewayRestApi_EndpointConfiguration AWS CloudFormation Resource (AWS::ApiGateway::RestApi.EndpointConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-restapi-endpointconfiguration.html
type AWSApiGatewayRestApi_EndpointConfiguration struct {

	// Types AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-restapi-endpointconfiguration.html#cfn-apigateway-restapi-endpointconfiguration-types
	Types []*Value `json:"Types,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApiGatewayRestApi_EndpointConfiguration) AWSCloudFormationType() string {
	return "AWS::ApiGateway::RestApi.EndpointConfiguration"
}

func (r *AWSApiGatewayRestApi_EndpointConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
