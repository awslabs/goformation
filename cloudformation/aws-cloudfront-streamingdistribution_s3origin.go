package cloudformation

import (
	"encoding/json"
)

// AWSCloudFrontStreamingDistribution_S3Origin AWS CloudFormation Resource (AWS::CloudFront::StreamingDistribution.S3Origin)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-streamingdistribution-s3origin.html
type AWSCloudFrontStreamingDistribution_S3Origin struct {

	// DomainName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-streamingdistribution-s3origin.html#cfn-cloudfront-streamingdistribution-s3origin-domainname
	DomainName *Value `json:"DomainName,omitempty"`

	// OriginAccessIdentity AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-streamingdistribution-s3origin.html#cfn-cloudfront-streamingdistribution-s3origin-originaccessidentity
	OriginAccessIdentity *Value `json:"OriginAccessIdentity,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCloudFrontStreamingDistribution_S3Origin) AWSCloudFormationType() string {
	return "AWS::CloudFront::StreamingDistribution.S3Origin"
}

func (r *AWSCloudFrontStreamingDistribution_S3Origin) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
