package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ApiGateway::ApiKey.StageKey AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-apikey-stagekey.html
type AWSApiGatewayApiKey_StageKey struct {

	// RestApiId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-apikey-stagekey.html#cfn-apigateway-apikey-stagekey-restapiid

	RestApiId string `json:"RestApiId"`

	// StageName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-apikey-stagekey.html#cfn-apigateway-apikey-stagekey-stagename

	StageName string `json:"StageName"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApiGatewayApiKey_StageKey) AWSCloudFormationType() string {
	return "AWS::ApiGateway::ApiKey.StageKey"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSApiGatewayApiKey_StageKey) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSApiGatewayApiKey_StageKeyResources retrieves all AWSApiGatewayApiKey_StageKey items from a CloudFormation template
func GetAllAWSApiGatewayApiKey_StageKey(template *Template) map[string]*AWSApiGatewayApiKey_StageKey {

	results := map[string]*AWSApiGatewayApiKey_StageKey{}
	for name, resource := range template.Resources {
		result := &AWSApiGatewayApiKey_StageKey{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSApiGatewayApiKey_StageKeyWithName retrieves all AWSApiGatewayApiKey_StageKey items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSApiGatewayApiKey_StageKey(name string, template *Template) (*AWSApiGatewayApiKey_StageKey, error) {

	result := &AWSApiGatewayApiKey_StageKey{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSApiGatewayApiKey_StageKey{}, errors.New("resource not found")

}
