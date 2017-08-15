package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ApiGateway::UsagePlan AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplan.html
type AWSApiGatewayUsagePlan struct {

	// ApiStages AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplan.html#cfn-apigateway-usageplan-apistages
	ApiStages []AWSApiGatewayUsagePlan_ApiStage `json:"ApiStages"`

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplan.html#cfn-apigateway-usageplan-description
	Description string `json:"Description"`

	// Quota AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplan.html#cfn-apigateway-usageplan-quota
	Quota AWSApiGatewayUsagePlan_QuotaSettings `json:"Quota"`

	// Throttle AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplan.html#cfn-apigateway-usageplan-throttle
	Throttle AWSApiGatewayUsagePlan_ThrottleSettings `json:"Throttle"`

	// UsagePlanName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-usageplan.html#cfn-apigateway-usageplan-usageplanname
	UsagePlanName string `json:"UsagePlanName"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApiGatewayUsagePlan) AWSCloudFormationType() string {
	return "AWS::ApiGateway::UsagePlan"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSApiGatewayUsagePlan) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSApiGatewayUsagePlanResources retrieves all AWSApiGatewayUsagePlan items from a CloudFormation template
func GetAllAWSApiGatewayUsagePlanResources(template *Template) map[string]*AWSApiGatewayUsagePlan {

	results := map[string]*AWSApiGatewayUsagePlan{}
	for name, resource := range template.Resources {
		result := &AWSApiGatewayUsagePlan{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSApiGatewayUsagePlanWithName retrieves all AWSApiGatewayUsagePlan items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSApiGatewayUsagePlanWithName(name string, template *Template) (*AWSApiGatewayUsagePlan, error) {

	result := &AWSApiGatewayUsagePlan{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSApiGatewayUsagePlan{}, errors.New("resource not found")

}
