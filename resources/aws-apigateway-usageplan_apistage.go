package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ApiGateway::UsagePlan.ApiStage AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-apistage.html
type AWSApiGatewayUsagePlan_ApiStage struct {

	// ApiId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-apistage.html#cfn-apigateway-usageplan-apistage-apiid
	ApiId string `json:"ApiId"`

	// Stage AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-apistage.html#cfn-apigateway-usageplan-apistage-stage
	Stage string `json:"Stage"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApiGatewayUsagePlan_ApiStage) AWSCloudFormationType() string {
	return "AWS::ApiGateway::UsagePlan.ApiStage"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSApiGatewayUsagePlan_ApiStage) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSApiGatewayUsagePlan_ApiStageResources retrieves all AWSApiGatewayUsagePlan_ApiStage items from a CloudFormation template
func GetAllAWSApiGatewayUsagePlan_ApiStage(template *Template) map[string]*AWSApiGatewayUsagePlan_ApiStage {

	results := map[string]*AWSApiGatewayUsagePlan_ApiStage{}
	for name, resource := range template.Resources {
		result := &AWSApiGatewayUsagePlan_ApiStage{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSApiGatewayUsagePlan_ApiStageWithName retrieves all AWSApiGatewayUsagePlan_ApiStage items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSApiGatewayUsagePlan_ApiStage(name string, template *Template) (*AWSApiGatewayUsagePlan_ApiStage, error) {

	result := &AWSApiGatewayUsagePlan_ApiStage{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSApiGatewayUsagePlan_ApiStage{}, errors.New("resource not found")

}
