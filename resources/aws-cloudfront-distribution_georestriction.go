package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::CloudFront::Distribution.GeoRestriction AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig-restrictions-georestriction.html
type AWSCloudFrontDistribution_GeoRestriction struct {

	// Locations AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig-restrictions-georestriction.html#cfn-cloudfront-distributionconfig-restrictions-georestriction-locations
	Locations []string `json:"Locations"`

	// RestrictionType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig-restrictions-georestriction.html#cfn-cloudfront-distributionconfig-restrictions-georestriction-restrictiontype
	RestrictionType string `json:"RestrictionType"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCloudFrontDistribution_GeoRestriction) AWSCloudFormationType() string {
	return "AWS::CloudFront::Distribution.GeoRestriction"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCloudFrontDistribution_GeoRestriction) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCloudFrontDistribution_GeoRestrictionResources retrieves all AWSCloudFrontDistribution_GeoRestriction items from a CloudFormation template
func GetAllAWSCloudFrontDistribution_GeoRestriction(template *Template) map[string]*AWSCloudFrontDistribution_GeoRestriction {

	results := map[string]*AWSCloudFrontDistribution_GeoRestriction{}
	for name, resource := range template.Resources {
		result := &AWSCloudFrontDistribution_GeoRestriction{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCloudFrontDistribution_GeoRestrictionWithName retrieves all AWSCloudFrontDistribution_GeoRestriction items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCloudFrontDistribution_GeoRestriction(name string, template *Template) (*AWSCloudFrontDistribution_GeoRestriction, error) {

	result := &AWSCloudFrontDistribution_GeoRestriction{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCloudFrontDistribution_GeoRestriction{}, errors.New("resource not found")

}
