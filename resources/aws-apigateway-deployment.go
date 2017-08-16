package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSApiGatewayDeployment AWS CloudFormation Resource (AWS::ApiGateway::Deployment)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-deployment.html
type AWSApiGatewayDeployment struct {

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-deployment.html#cfn-apigateway-deployment-description
	Description string `json:"Description"`

	// RestApiId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-deployment.html#cfn-apigateway-deployment-restapiid
	RestApiId string `json:"RestApiId"`

	// StageDescription AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-deployment.html#cfn-apigateway-deployment-stagedescription
	StageDescription AWSApiGatewayDeployment_StageDescription `json:"StageDescription"`

	// StageName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-deployment.html#cfn-apigateway-deployment-stagename
	StageName string `json:"StageName"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApiGatewayDeployment) AWSCloudFormationType() string {
	return "AWS::ApiGateway::Deployment"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSApiGatewayDeployment) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSApiGatewayDeploymentResources retrieves all AWSApiGatewayDeployment items from a CloudFormation template
func (t *Template) GetAllAWSApiGatewayDeploymentResources() map[string]*AWSApiGatewayDeployment {

	results := map[string]*AWSApiGatewayDeployment{}
	for name, resource := range t.Resources {
		result := &AWSApiGatewayDeployment{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSApiGatewayDeploymentWithName retrieves all AWSApiGatewayDeployment items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSApiGatewayDeploymentWithName(name string) (*AWSApiGatewayDeployment, error) {

	result := &AWSApiGatewayDeployment{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSApiGatewayDeployment{}, errors.New("resource not found")

}
