package cloudformation

import (
	"encoding/json"
)

// AWSCloudFrontStreamingDistribution_StreamingDistributionConfig AWS CloudFormation Resource (AWS::CloudFront::StreamingDistribution.StreamingDistributionConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-streamingdistribution-streamingdistributionconfig.html
type AWSCloudFrontStreamingDistribution_StreamingDistributionConfig struct {

	// Aliases AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-streamingdistribution-streamingdistributionconfig.html#cfn-cloudfront-streamingdistribution-streamingdistributionconfig-aliases
	Aliases []*Value `json:"Aliases,omitempty"`

	// Comment AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-streamingdistribution-streamingdistributionconfig.html#cfn-cloudfront-streamingdistribution-streamingdistributionconfig-comment
	Comment *Value `json:"Comment,omitempty"`

	// Enabled AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-streamingdistribution-streamingdistributionconfig.html#cfn-cloudfront-streamingdistribution-streamingdistributionconfig-enabled
	Enabled *Value `json:"Enabled,omitempty"`

	// Logging AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-streamingdistribution-streamingdistributionconfig.html#cfn-cloudfront-streamingdistribution-streamingdistributionconfig-logging
	Logging *AWSCloudFrontStreamingDistribution_Logging `json:"Logging,omitempty"`

	// PriceClass AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-streamingdistribution-streamingdistributionconfig.html#cfn-cloudfront-streamingdistribution-streamingdistributionconfig-priceclass
	PriceClass *Value `json:"PriceClass,omitempty"`

	// S3Origin AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-streamingdistribution-streamingdistributionconfig.html#cfn-cloudfront-streamingdistribution-streamingdistributionconfig-s3origin
	S3Origin *AWSCloudFrontStreamingDistribution_S3Origin `json:"S3Origin,omitempty"`

	// TrustedSigners AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-streamingdistribution-streamingdistributionconfig.html#cfn-cloudfront-streamingdistribution-streamingdistributionconfig-trustedsigners
	TrustedSigners *AWSCloudFrontStreamingDistribution_TrustedSigners `json:"TrustedSigners,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCloudFrontStreamingDistribution_StreamingDistributionConfig) AWSCloudFormationType() string {
	return "AWS::CloudFront::StreamingDistribution.StreamingDistributionConfig"
}

func (r *AWSCloudFrontStreamingDistribution_StreamingDistributionConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
