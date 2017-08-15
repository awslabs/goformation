package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ApiGateway::Deployment.MethodSetting AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription-methodsetting.html
type AWSApiGatewayDeployment_MethodSetting struct {

	// CacheDataEncrypted AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription-methodsetting.html#cfn-apigateway-deployment-stagedescription-methodsetting-cachedataencrypted
	CacheDataEncrypted bool `json:"CacheDataEncrypted"`

	// CacheTtlInSeconds AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription-methodsetting.html#cfn-apigateway-deployment-stagedescription-methodsetting-cachettlinseconds
	CacheTtlInSeconds int64 `json:"CacheTtlInSeconds"`

	// CachingEnabled AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription-methodsetting.html#cfn-apigateway-deployment-stagedescription-methodsetting-cachingenabled
	CachingEnabled bool `json:"CachingEnabled"`

	// DataTraceEnabled AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription-methodsetting.html#cfn-apigateway-deployment-stagedescription-methodsetting-datatraceenabled
	DataTraceEnabled bool `json:"DataTraceEnabled"`

	// HttpMethod AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription-methodsetting.html#cfn-apigateway-deployment-stagedescription-methodsetting-httpmethod
	HttpMethod string `json:"HttpMethod"`

	// LoggingLevel AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription-methodsetting.html#cfn-apigateway-deployment-stagedescription-methodsetting-logginglevel
	LoggingLevel string `json:"LoggingLevel"`

	// MetricsEnabled AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription-methodsetting.html#cfn-apigateway-deployment-stagedescription-methodsetting-metricsenabled
	MetricsEnabled bool `json:"MetricsEnabled"`

	// ResourcePath AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription-methodsetting.html#cfn-apigateway-deployment-stagedescription-methodsetting-resourcepath
	ResourcePath string `json:"ResourcePath"`

	// ThrottlingBurstLimit AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription-methodsetting.html#cfn-apigateway-deployment-stagedescription-methodsetting-throttlingburstlimit
	ThrottlingBurstLimit int64 `json:"ThrottlingBurstLimit"`

	// ThrottlingRateLimit AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription-methodsetting.html#cfn-apigateway-deployment-stagedescription-methodsetting-throttlingratelimit
	ThrottlingRateLimit float64 `json:"ThrottlingRateLimit"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApiGatewayDeployment_MethodSetting) AWSCloudFormationType() string {
	return "AWS::ApiGateway::Deployment.MethodSetting"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSApiGatewayDeployment_MethodSetting) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSApiGatewayDeployment_MethodSettingResources retrieves all AWSApiGatewayDeployment_MethodSetting items from a CloudFormation template
func GetAllAWSApiGatewayDeployment_MethodSetting(template *Template) map[string]*AWSApiGatewayDeployment_MethodSetting {

	results := map[string]*AWSApiGatewayDeployment_MethodSetting{}
	for name, resource := range template.Resources {
		result := &AWSApiGatewayDeployment_MethodSetting{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSApiGatewayDeployment_MethodSettingWithName retrieves all AWSApiGatewayDeployment_MethodSetting items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSApiGatewayDeployment_MethodSetting(name string, template *Template) (*AWSApiGatewayDeployment_MethodSetting, error) {

	result := &AWSApiGatewayDeployment_MethodSetting{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSApiGatewayDeployment_MethodSetting{}, errors.New("resource not found")

}
