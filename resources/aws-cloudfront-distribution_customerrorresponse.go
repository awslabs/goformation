package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::CloudFront::Distribution.CustomErrorResponse AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig-customerrorresponse.html
type AWSCloudFrontDistribution_CustomErrorResponse struct {

	// ErrorCachingMinTTL AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig-customerrorresponse.html#cfn-cloudfront-distributionconfig-customerrorresponse-errorcachingminttl
	ErrorCachingMinTTL int64 `json:"ErrorCachingMinTTL"`

	// ErrorCode AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig-customerrorresponse.html#cfn-cloudfront-distributionconfig-customerrorresponse-errorcode
	ErrorCode int64 `json:"ErrorCode"`

	// ResponseCode AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig-customerrorresponse.html#cfn-cloudfront-distributionconfig-customerrorresponse-responsecode
	ResponseCode int64 `json:"ResponseCode"`

	// ResponsePagePath AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig-customerrorresponse.html#cfn-cloudfront-distributionconfig-customerrorresponse-responsepagepath
	ResponsePagePath string `json:"ResponsePagePath"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCloudFrontDistribution_CustomErrorResponse) AWSCloudFormationType() string {
	return "AWS::CloudFront::Distribution.CustomErrorResponse"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCloudFrontDistribution_CustomErrorResponse) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCloudFrontDistribution_CustomErrorResponseResources retrieves all AWSCloudFrontDistribution_CustomErrorResponse items from a CloudFormation template
func GetAllAWSCloudFrontDistribution_CustomErrorResponse(template *Template) map[string]*AWSCloudFrontDistribution_CustomErrorResponse {

	results := map[string]*AWSCloudFrontDistribution_CustomErrorResponse{}
	for name, resource := range template.Resources {
		result := &AWSCloudFrontDistribution_CustomErrorResponse{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCloudFrontDistribution_CustomErrorResponseWithName retrieves all AWSCloudFrontDistribution_CustomErrorResponse items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCloudFrontDistribution_CustomErrorResponse(name string, template *Template) (*AWSCloudFrontDistribution_CustomErrorResponse, error) {

	result := &AWSCloudFrontDistribution_CustomErrorResponse{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCloudFrontDistribution_CustomErrorResponse{}, errors.New("resource not found")

}
