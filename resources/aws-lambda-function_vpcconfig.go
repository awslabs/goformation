package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Lambda::Function.VpcConfig AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-vpcconfig.html
type AWSLambdaFunction_VpcConfig struct {

	// SecurityGroupIds AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-vpcconfig.html#cfn-lambda-function-vpcconfig-securitygroupids

	SecurityGroupIds []string `json:"SecurityGroupIds"`

	// SubnetIds AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lambda-function-vpcconfig.html#cfn-lambda-function-vpcconfig-subnetids

	SubnetIds []string `json:"SubnetIds"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSLambdaFunction_VpcConfig) AWSCloudFormationType() string {
	return "AWS::Lambda::Function.VpcConfig"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSLambdaFunction_VpcConfig) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSLambdaFunction_VpcConfigResources retrieves all AWSLambdaFunction_VpcConfig items from a CloudFormation template
func GetAllAWSLambdaFunction_VpcConfig(template *Template) map[string]*AWSLambdaFunction_VpcConfig {

	results := map[string]*AWSLambdaFunction_VpcConfig{}
	for name, resource := range template.Resources {
		result := &AWSLambdaFunction_VpcConfig{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSLambdaFunction_VpcConfigWithName retrieves all AWSLambdaFunction_VpcConfig items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSLambdaFunction_VpcConfig(name string, template *Template) (*AWSLambdaFunction_VpcConfig, error) {

	result := &AWSLambdaFunction_VpcConfig{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSLambdaFunction_VpcConfig{}, errors.New("resource not found")

}
