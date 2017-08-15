package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ApiGateway::Deployment.StageDescription AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription.html
type AWSApiGatewayDeployment_StageDescription struct {

	// CacheClusterEnabled AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-cacheclusterenabled
	CacheClusterEnabled bool `json:"CacheClusterEnabled"`

	// CacheClusterSize AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-cacheclustersize
	CacheClusterSize string `json:"CacheClusterSize"`

	// CacheDataEncrypted AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-cachedataencrypted
	CacheDataEncrypted bool `json:"CacheDataEncrypted"`

	// CacheTtlInSeconds AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-cachettlinseconds
	CacheTtlInSeconds int64 `json:"CacheTtlInSeconds"`

	// CachingEnabled AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-cachingenabled
	CachingEnabled bool `json:"CachingEnabled"`

	// ClientCertificateId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-clientcertificateid
	ClientCertificateId string `json:"ClientCertificateId"`

	// DataTraceEnabled AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-datatraceenabled
	DataTraceEnabled bool `json:"DataTraceEnabled"`

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-description
	Description string `json:"Description"`

	// LoggingLevel AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-logginglevel
	LoggingLevel string `json:"LoggingLevel"`

	// MethodSettings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-methodsettings
	MethodSettings []AWSApiGatewayDeployment_MethodSetting `json:"MethodSettings"`

	// MetricsEnabled AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-metricsenabled
	MetricsEnabled bool `json:"MetricsEnabled"`

	// StageName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-stagename
	StageName string `json:"StageName"`

	// ThrottlingBurstLimit AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-throttlingburstlimit
	ThrottlingBurstLimit int64 `json:"ThrottlingBurstLimit"`

	// ThrottlingRateLimit AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-throttlingratelimit
	ThrottlingRateLimit float64 `json:"ThrottlingRateLimit"`

	// Variables AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-apitgateway-deployment-stagedescription.html#cfn-apigateway-deployment-stagedescription-variables
	Variables map[string]string `json:"Variables"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApiGatewayDeployment_StageDescription) AWSCloudFormationType() string {
	return "AWS::ApiGateway::Deployment.StageDescription"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSApiGatewayDeployment_StageDescription) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSApiGatewayDeployment_StageDescriptionResources retrieves all AWSApiGatewayDeployment_StageDescription items from a CloudFormation template
func GetAllAWSApiGatewayDeployment_StageDescription(template *Template) map[string]*AWSApiGatewayDeployment_StageDescription {

	results := map[string]*AWSApiGatewayDeployment_StageDescription{}
	for name, resource := range template.Resources {
		result := &AWSApiGatewayDeployment_StageDescription{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSApiGatewayDeployment_StageDescriptionWithName retrieves all AWSApiGatewayDeployment_StageDescription items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSApiGatewayDeployment_StageDescription(name string, template *Template) (*AWSApiGatewayDeployment_StageDescription, error) {

	result := &AWSApiGatewayDeployment_StageDescription{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSApiGatewayDeployment_StageDescription{}, errors.New("resource not found")

}
