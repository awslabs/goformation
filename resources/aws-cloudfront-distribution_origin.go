package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::CloudFront::Distribution.Origin AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-origin.html
type AWSCloudFrontDistribution_Origin struct {

	// CustomOriginConfig AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-origin.html#cfn-cloudfront-origin-customorigin
	CustomOriginConfig AWSCloudFrontDistribution_CustomOriginConfig `json:"CustomOriginConfig"`

	// DomainName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-origin.html#cfn-cloudfront-origin-domainname
	DomainName string `json:"DomainName"`

	// Id AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-origin.html#cfn-cloudfront-origin-id
	Id string `json:"Id"`

	// OriginCustomHeaders AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-origin.html#cfn-cloudfront-origin-origincustomheaders
	OriginCustomHeaders []AWSCloudFrontDistribution_OriginCustomHeader `json:"OriginCustomHeaders"`

	// OriginPath AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-origin.html#cfn-cloudfront-origin-originpath
	OriginPath string `json:"OriginPath"`

	// S3OriginConfig AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-origin.html#cfn-cloudfront-origin-s3origin
	S3OriginConfig AWSCloudFrontDistribution_S3OriginConfig `json:"S3OriginConfig"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCloudFrontDistribution_Origin) AWSCloudFormationType() string {
	return "AWS::CloudFront::Distribution.Origin"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCloudFrontDistribution_Origin) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCloudFrontDistribution_OriginResources retrieves all AWSCloudFrontDistribution_Origin items from a CloudFormation template
func GetAllAWSCloudFrontDistribution_Origin(template *Template) map[string]*AWSCloudFrontDistribution_Origin {

	results := map[string]*AWSCloudFrontDistribution_Origin{}
	for name, resource := range template.Resources {
		result := &AWSCloudFrontDistribution_Origin{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCloudFrontDistribution_OriginWithName retrieves all AWSCloudFrontDistribution_Origin items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCloudFrontDistribution_Origin(name string, template *Template) (*AWSCloudFrontDistribution_Origin, error) {

	result := &AWSCloudFrontDistribution_Origin{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCloudFrontDistribution_Origin{}, errors.New("resource not found")

}
