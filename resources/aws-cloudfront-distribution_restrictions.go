package resources

// AWS::CloudFront::Distribution.Restrictions AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig-restrictions.html
type AWSCloudFrontDistributionRestrictions struct {
    
    // GeoRestriction AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distributionconfig-restrictions.html#cfn-cloudfront-distributionconfig-restrictions-georestriction
    GeoRestriction AWSCloudFrontDistributionRestrictionsGeoRestriction `json:"GeoRestriction"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCloudFrontDistributionRestrictions) AWSCloudFormationType() string {
    return "AWS::CloudFront::Distribution.Restrictions"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCloudFrontDistributionRestrictions) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}