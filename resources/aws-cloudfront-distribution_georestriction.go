package resources

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
