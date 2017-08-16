package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSCloudFrontDistribution AWS CloudFormation Resource (AWS::CloudFront::Distribution)
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
func (t *Template) GetAllAWSCloudFrontDistributionResources() map[string]*AWSCloudFrontDistribution {

	results := map[string]*AWSCloudFrontDistribution{}
	for name, resource := range t.Resources {
		result := &AWSCloudFrontDistribution{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCloudFrontDistributionWithName retrieves all AWSCloudFrontDistribution items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSCloudFrontDistributionWithName(name string) (*AWSCloudFrontDistribution, error) {

	result := &AWSCloudFrontDistribution{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCloudFrontDistribution{}, errors.New("resource not found")

}
