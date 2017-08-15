package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSApiGatewayApiKey AWS CloudFormation Resource (AWS::ApiGateway::ApiKey)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-apikey.html
type AWSApiGatewayApiKey struct {

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-apikey.html#cfn-apigateway-apigateway-apikey-description
	Description string `json:"Description"`

	// Enabled AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-apikey.html#cfn-apigateway-apigateway-apikey-enabled
	Enabled bool `json:"Enabled"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-apikey.html#cfn-apigateway-apigateway-apikey-name
	Name string `json:"Name"`

	// StageKeys AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-apikey.html#cfn-apigateway-apigateway-apikey-stagekeys
	StageKeys []AWSApiGatewayApiKey_StageKey `json:"StageKeys"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApiGatewayApiKey) AWSCloudFormationType() string {
	return "AWS::ApiGateway::ApiKey"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSApiGatewayApiKey) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSApiGatewayApiKeyResources retrieves all AWSApiGatewayApiKey items from a CloudFormation template
func GetAllAWSApiGatewayApiKeyResources(template *Template) map[string]*AWSApiGatewayApiKey {

	results := map[string]*AWSApiGatewayApiKey{}
	for name, resource := range template.Resources {
		result := &AWSApiGatewayApiKey{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSApiGatewayApiKeyWithName retrieves all AWSApiGatewayApiKey items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSApiGatewayApiKeyWithName(name string, template *Template) (*AWSApiGatewayApiKey, error) {

	result := &AWSApiGatewayApiKey{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSApiGatewayApiKey{}, errors.New("resource not found")

}
