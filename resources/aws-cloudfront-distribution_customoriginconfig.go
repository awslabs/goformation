package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::CloudFront::Distribution.CustomOriginConfig AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-customorigin.html
type AWSCloudFrontDistribution_CustomOriginConfig struct {

	// HTTPPort AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-customorigin.html#cfn-cloudfront-customorigin-httpport
	HTTPPort int64 `json:"HTTPPort"`

	// HTTPSPort AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-customorigin.html#cfn-cloudfront-customorigin-httpsport
	HTTPSPort int64 `json:"HTTPSPort"`

	// OriginProtocolPolicy AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-customorigin.html#cfn-cloudfront-customorigin-originprotocolpolicy
	OriginProtocolPolicy string `json:"OriginProtocolPolicy"`

	// OriginSSLProtocols AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-customorigin.html#cfn-cloudfront-customorigin-originsslprotocols
	OriginSSLProtocols []string `json:"OriginSSLProtocols"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCloudFrontDistribution_CustomOriginConfig) AWSCloudFormationType() string {
	return "AWS::CloudFront::Distribution.CustomOriginConfig"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCloudFrontDistribution_CustomOriginConfig) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCloudFrontDistribution_CustomOriginConfigResources retrieves all AWSCloudFrontDistribution_CustomOriginConfig items from a CloudFormation template
func GetAllAWSCloudFrontDistribution_CustomOriginConfig(template *Template) map[string]*AWSCloudFrontDistribution_CustomOriginConfig {

	results := map[string]*AWSCloudFrontDistribution_CustomOriginConfig{}
	for name, resource := range template.Resources {
		result := &AWSCloudFrontDistribution_CustomOriginConfig{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCloudFrontDistribution_CustomOriginConfigWithName retrieves all AWSCloudFrontDistribution_CustomOriginConfig items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCloudFrontDistribution_CustomOriginConfig(name string, template *Template) (*AWSCloudFrontDistribution_CustomOriginConfig, error) {

	result := &AWSCloudFrontDistribution_CustomOriginConfig{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCloudFrontDistribution_CustomOriginConfig{}, errors.New("resource not found")

}
