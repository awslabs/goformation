package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::CloudFront::Distribution AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution.html
type AWSCloudFrontDistribution struct {

	// DistributionConfig AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution.html#cfn-cloudfront-distribution-distributionconfig
	DistributionConfig AWSCloudFrontDistribution_DistributionConfig `json:"DistributionConfig"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCloudFrontDistribution) AWSCloudFormationType() string {
	return "AWS::CloudFront::Distribution"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCloudFrontDistribution) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCloudFrontDistributionResources retrieves all AWSCloudFrontDistribution items from a CloudFormation template
func GetAllAWSCloudFrontDistributionResources(template *Template) map[string]*AWSCloudFrontDistribution {

	results := map[string]*AWSCloudFrontDistribution{}
	for name, resource := range template.Resources {
		result := &AWSCloudFrontDistribution{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCloudFrontDistributionWithName retrieves all AWSCloudFrontDistribution items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSCloudFrontDistributionWithName(name string, template *Template) (*AWSCloudFrontDistribution, error) {

	result := &AWSCloudFrontDistribution{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCloudFrontDistribution{}, errors.New("resource not found")

}
