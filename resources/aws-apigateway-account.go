package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSApiGatewayAccount AWS CloudFormation Resource (AWS::ApiGateway::Account)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-account.html
type AWSApiGatewayAccount struct {

	// CloudWatchRoleArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-account.html#cfn-apigateway-account-cloudwatchrolearn
	CloudWatchRoleArn string `json:"CloudWatchRoleArn"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApiGatewayAccount) AWSCloudFormationType() string {
	return "AWS::ApiGateway::Account"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSApiGatewayAccount) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSApiGatewayAccountResources retrieves all AWSApiGatewayAccount items from a CloudFormation template
func (t *Template) GetAllAWSApiGatewayAccountResources() map[string]*AWSApiGatewayAccount {

	results := map[string]*AWSApiGatewayAccount{}
	for name, resource := range t.Resources {
		result := &AWSApiGatewayAccount{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSApiGatewayAccountWithName retrieves all AWSApiGatewayAccount items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayAccountWithName(name string) (*AWSApiGatewayAccount, error) {

	result := &AWSApiGatewayAccount{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSApiGatewayAccount{}, errors.New("resource not found")

}
