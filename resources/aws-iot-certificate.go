package resources

// AWS::IoT::Certificate AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-certificate.html
type AWSIoTCertificate struct {
    
    // CertificateSigningRequest AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-certificate.html#cfn-iot-certificate-certificatesigningrequest
    CertificateSigningRequest string `json:"CertificateSigningRequest"`
    
    // Status AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-certificate.html#cfn-iot-certificate-status
    Status string `json:"Status"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTCertificate) AWSCloudFormationType() string {
    return "AWS::IoT::Certificate"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIoTCertificate) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}