package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::CloudFront::Distribution.ViewerCertificate AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig-viewercertificate.html
type AWSCloudFrontDistribution_ViewerCertificate struct {

	// AcmCertificateArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig-viewercertificate.html#cfn-cloudfront-distributionconfig-viewercertificate-acmcertificatearn

	AcmCertificateArn string `json:"AcmCertificateArn"`

	// CloudFrontDefaultCertificate AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig-viewercertificate.html#cfn-cloudfront-distributionconfig-viewercertificate-cloudfrontdefaultcertificate

	CloudFrontDefaultCertificate bool `json:"CloudFrontDefaultCertificate"`

	// IamCertificateId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig-viewercertificate.html#cfn-cloudfront-distributionconfig-viewercertificate-iamcertificateid

	IamCertificateId string `json:"IamCertificateId"`

	// MinimumProtocolVersion AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig-viewercertificate.html#cfn-cloudfront-distributionconfig-viewercertificate-sslsupportmethod

	MinimumProtocolVersion string `json:"MinimumProtocolVersion"`

	// SslSupportMethod AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig-viewercertificate.html#cfn-cloudfront-distributionconfig-viewercertificate-minimumprotocolversion

	SslSupportMethod string `json:"SslSupportMethod"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCloudFrontDistribution_ViewerCertificate) AWSCloudFormationType() string {
	return "AWS::CloudFront::Distribution.ViewerCertificate"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCloudFrontDistribution_ViewerCertificate) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCloudFrontDistribution_ViewerCertificateResources retrieves all AWSCloudFrontDistribution_ViewerCertificate items from a CloudFormation template
func GetAllAWSCloudFrontDistribution_ViewerCertificate(template *Template) map[string]*AWSCloudFrontDistribution_ViewerCertificate {

	results := map[string]*AWSCloudFrontDistribution_ViewerCertificate{}
	for name, resource := range template.Resources {
		result := &AWSCloudFrontDistribution_ViewerCertificate{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCloudFrontDistribution_ViewerCertificateWithName retrieves all AWSCloudFrontDistribution_ViewerCertificate items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCloudFrontDistribution_ViewerCertificate(name string, template *Template) (*AWSCloudFrontDistribution_ViewerCertificate, error) {

	result := &AWSCloudFrontDistribution_ViewerCertificate{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCloudFrontDistribution_ViewerCertificate{}, errors.New("resource not found")

}
