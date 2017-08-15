package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ApiGateway::UsagePlanKey AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplankey.html
type AWSApiGatewayUsagePlanKey struct {

	// KeyId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplankey.html#cfn-apigateway-usageplankey-keyid
	KeyId string `json:"KeyId"`

	// KeyType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplankey.html#cfn-apigateway-usageplankey-keytype
	KeyType string `json:"KeyType"`

	// UsagePlanId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplankey.html#cfn-apigateway-usageplankey-usageplanid
	UsagePlanId string `json:"UsagePlanId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApiGatewayUsagePlanKey) AWSCloudFormationType() string {
	return "AWS::ApiGateway::UsagePlanKey"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSApiGatewayUsagePlanKey) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSApiGatewayUsagePlanKeyResources retrieves all AWSApiGatewayUsagePlanKey items from a CloudFormation template
func GetAllAWSApiGatewayUsagePlanKey(template *Template) map[string]*AWSApiGatewayUsagePlanKey {

	results := map[string]*AWSApiGatewayUsagePlanKey{}
	for name, resource := range template.Resources {
		result := &AWSApiGatewayUsagePlanKey{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSApiGatewayUsagePlanKeyWithName retrieves all AWSApiGatewayUsagePlanKey items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSApiGatewayUsagePlanKey(name string, template *Template) (*AWSApiGatewayUsagePlanKey, error) {

	result := &AWSApiGatewayUsagePlanKey{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSApiGatewayUsagePlanKey{}, errors.New("resource not found")

}
