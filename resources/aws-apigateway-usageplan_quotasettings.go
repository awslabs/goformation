package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ApiGateway::UsagePlan.QuotaSettings AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-quotasettings.html
type AWSApiGatewayUsagePlan_QuotaSettings struct {

	// Limit AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-quotasettings.html#cfn-apigateway-usageplan-quotasettings-limit
	Limit int64 `json:"Limit"`

	// Offset AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-quotasettings.html#cfn-apigateway-usageplan-quotasettings-offset
	Offset int64 `json:"Offset"`

	// Period AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apigateway-usageplan-quotasettings.html#cfn-apigateway-usageplan-quotasettings-period
	Period string `json:"Period"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApiGatewayUsagePlan_QuotaSettings) AWSCloudFormationType() string {
	return "AWS::ApiGateway::UsagePlan.QuotaSettings"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSApiGatewayUsagePlan_QuotaSettings) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSApiGatewayUsagePlan_QuotaSettingsResources retrieves all AWSApiGatewayUsagePlan_QuotaSettings items from a CloudFormation template
func GetAllAWSApiGatewayUsagePlan_QuotaSettings(template *Template) map[string]*AWSApiGatewayUsagePlan_QuotaSettings {

	results := map[string]*AWSApiGatewayUsagePlan_QuotaSettings{}
	for name, resource := range template.Resources {
		result := &AWSApiGatewayUsagePlan_QuotaSettings{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSApiGatewayUsagePlan_QuotaSettingsWithName retrieves all AWSApiGatewayUsagePlan_QuotaSettings items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSApiGatewayUsagePlan_QuotaSettings(name string, template *Template) (*AWSApiGatewayUsagePlan_QuotaSettings, error) {

	result := &AWSApiGatewayUsagePlan_QuotaSettings{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSApiGatewayUsagePlan_QuotaSettings{}, errors.New("resource not found")

}
