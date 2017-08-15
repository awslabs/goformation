package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::CloudFront::Distribution.S3OriginConfig AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-s3origin.html
type AWSCloudFrontDistribution_S3OriginConfig struct {

	// OriginAccessIdentity AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-s3origin.html#cfn-cloudfront-s3origin-originaccessidentity
	OriginAccessIdentity string `json:"OriginAccessIdentity"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCloudFrontDistribution_S3OriginConfig) AWSCloudFormationType() string {
	return "AWS::CloudFront::Distribution.S3OriginConfig"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCloudFrontDistribution_S3OriginConfig) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCloudFrontDistribution_S3OriginConfigResources retrieves all AWSCloudFrontDistribution_S3OriginConfig items from a CloudFormation template
func GetAllAWSCloudFrontDistribution_S3OriginConfig(template *Template) map[string]*AWSCloudFrontDistribution_S3OriginConfig {

	results := map[string]*AWSCloudFrontDistribution_S3OriginConfig{}
	for name, resource := range template.Resources {
		result := &AWSCloudFrontDistribution_S3OriginConfig{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCloudFrontDistribution_S3OriginConfigWithName retrieves all AWSCloudFrontDistribution_S3OriginConfig items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCloudFrontDistribution_S3OriginConfig(name string, template *Template) (*AWSCloudFrontDistribution_S3OriginConfig, error) {

	result := &AWSCloudFrontDistribution_S3OriginConfig{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCloudFrontDistribution_S3OriginConfig{}, errors.New("resource not found")

}
