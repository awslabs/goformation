package resources

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
