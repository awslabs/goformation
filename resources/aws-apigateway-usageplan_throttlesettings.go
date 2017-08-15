package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ApiGateway::UsagePlan.ThrottleSettings AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-throttlesettings.html
type AWSApiGatewayUsagePlan_ThrottleSettings struct {

	// BurstLimit AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-throttlesettings.html#cfn-apigateway-usageplan-throttlesettings-burstlimit
	BurstLimit int64 `json:"BurstLimit"`

	// RateLimit AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-throttlesettings.html#cfn-apigateway-usageplan-throttlesettings-ratelimit
	RateLimit float64 `json:"RateLimit"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApiGatewayUsagePlan_ThrottleSettings) AWSCloudFormationType() string {
	return "AWS::ApiGateway::UsagePlan.ThrottleSettings"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSApiGatewayUsagePlan_ThrottleSettings) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSApiGatewayUsagePlan_ThrottleSettingsResources retrieves all AWSApiGatewayUsagePlan_ThrottleSettings items from a CloudFormation template
func GetAllAWSApiGatewayUsagePlan_ThrottleSettings(template *Template) map[string]*AWSApiGatewayUsagePlan_ThrottleSettings {

	results := map[string]*AWSApiGatewayUsagePlan_ThrottleSettings{}
	for name, resource := range template.Resources {
		result := &AWSApiGatewayUsagePlan_ThrottleSettings{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSApiGatewayUsagePlan_ThrottleSettingsWithName retrieves all AWSApiGatewayUsagePlan_ThrottleSettings items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSApiGatewayUsagePlan_ThrottleSettings(name string, template *Template) (*AWSApiGatewayUsagePlan_ThrottleSettings, error) {

	result := &AWSApiGatewayUsagePlan_ThrottleSettings{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSApiGatewayUsagePlan_ThrottleSettings{}, errors.New("resource not found")

}
