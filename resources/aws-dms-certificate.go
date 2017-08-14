package resources

// AWS::DMS::Certificate AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-certificate.html
type AWSDMSCertificate struct {

	// CertificateIdentifier AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-certificate.html#cfn-dms-certificate-certificateidentifier

	CertificateIdentifier string `json:"CertificateIdentifier"`

	// CertificatePem AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-certificate.html#cfn-dms-certificate-certificatepem

	CertificatePem string `json:"CertificatePem"`

	// CertificateWallet AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-dms-certificate.html#cfn-dms-certificate-certificatewallet

	CertificateWallet string `json:"CertificateWallet"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDMSCertificate) AWSCloudFormationType() string {
	return "AWS::DMS::Certificate"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSDMSCertificate) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
