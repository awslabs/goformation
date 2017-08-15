package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::CloudFront::Distribution.Restrictions AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig-restrictions.html
type AWSCloudFrontDistribution_Restrictions struct {

	// GeoRestriction AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig-restrictions.html#cfn-cloudfront-distributionconfig-restrictions-georestriction

	GeoRestriction AWSCloudFrontDistribution_GeoRestriction `json:"GeoRestriction"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCloudFrontDistribution_Restrictions) AWSCloudFormationType() string {
	return "AWS::CloudFront::Distribution.Restrictions"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCloudFrontDistribution_Restrictions) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCloudFrontDistribution_RestrictionsResources retrieves all AWSCloudFrontDistribution_Restrictions items from a CloudFormation template
func GetAllAWSCloudFrontDistribution_Restrictions(template *Template) map[string]*AWSCloudFrontDistribution_Restrictions {

	results := map[string]*AWSCloudFrontDistribution_Restrictions{}
	for name, resource := range template.Resources {
		result := &AWSCloudFrontDistribution_Restrictions{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCloudFrontDistribution_RestrictionsWithName retrieves all AWSCloudFrontDistribution_Restrictions items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCloudFrontDistribution_Restrictions(name string, template *Template) (*AWSCloudFrontDistribution_Restrictions, error) {

	result := &AWSCloudFrontDistribution_Restrictions{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCloudFrontDistribution_Restrictions{}, errors.New("resource not found")

}
