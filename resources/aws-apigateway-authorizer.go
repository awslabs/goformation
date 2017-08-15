package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ApiGateway::Authorizer AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-authorizer.html
type AWSApiGatewayAuthorizer struct {

	// AuthorizerCredentials AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-authorizer.html#cfn-apigateway-authorizer-authorizercredentials
	AuthorizerCredentials string `json:"AuthorizerCredentials"`

	// AuthorizerResultTtlInSeconds AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-authorizer.html#cfn-apigateway-authorizer-authorizerresultttlinseconds
	AuthorizerResultTtlInSeconds int64 `json:"AuthorizerResultTtlInSeconds"`

	// AuthorizerUri AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-authorizer.html#cfn-apigateway-authorizer-authorizeruri
	AuthorizerUri string `json:"AuthorizerUri"`

	// IdentitySource AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-authorizer.html#cfn-apigateway-authorizer-identitysource
	IdentitySource string `json:"IdentitySource"`

	// IdentityValidationExpression AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-authorizer.html#cfn-apigateway-authorizer-identityvalidationexpression
	IdentityValidationExpression string `json:"IdentityValidationExpression"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-authorizer.html#cfn-apigateway-authorizer-name
	Name string `json:"Name"`

	// ProviderARNs AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-authorizer.html#cfn-apigateway-authorizer-providerarns
	ProviderARNs []string `json:"ProviderARNs"`

	// RestApiId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-authorizer.html#cfn-apigateway-authorizer-restapiid
	RestApiId string `json:"RestApiId"`

	// Type AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-authorizer.html#cfn-apigateway-authorizer-type
	Type string `json:"Type"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApiGatewayAuthorizer) AWSCloudFormationType() string {
	return "AWS::ApiGateway::Authorizer"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSApiGatewayAuthorizer) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSApiGatewayAuthorizerResources retrieves all AWSApiGatewayAuthorizer items from a CloudFormation template
func GetAllAWSApiGatewayAuthorizer(template *Template) map[string]*AWSApiGatewayAuthorizer {

	results := map[string]*AWSApiGatewayAuthorizer{}
	for name, resource := range template.Resources {
		result := &AWSApiGatewayAuthorizer{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSApiGatewayAuthorizerWithName retrieves all AWSApiGatewayAuthorizer items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSApiGatewayAuthorizer(name string, template *Template) (*AWSApiGatewayAuthorizer, error) {

	result := &AWSApiGatewayAuthorizer{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSApiGatewayAuthorizer{}, errors.New("resource not found")

}
