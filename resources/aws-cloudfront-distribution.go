package resources

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
