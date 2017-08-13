package resources

// AWS::CertificateManager::Certificate.DomainValidationOption AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-certificatemanager-certificate-domainvalidationoption.html
type AWSCertificateManagerCertificateDomainValidationOption struct {
    
    // DomainName AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-certificatemanager-certificate-domainvalidationoption.html#cfn-certificatemanager-certificate-domainvalidationoptions-domainname
    DomainName string `json:"DomainName"`
    
    // ValidationDomain AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-certificatemanager-certificate-domainvalidationoption.html#cfn-certificatemanager-certificate-domainvalidationoption-validationdomain
    ValidationDomain string `json:"ValidationDomain"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCertificateManagerCertificateDomainValidationOption) AWSCloudFormationType() string {
    return "AWS::CertificateManager::Certificate.DomainValidationOption"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCertificateManagerCertificateDomainValidationOption) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}