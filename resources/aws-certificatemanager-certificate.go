package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSCertificateManagerCertificate AWS CloudFormation Resource (AWS::CertificateManager::Certificate)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-certificate.html
type AWSCertificateManagerCertificate struct {

	// DomainName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-certificate.html#cfn-certificatemanager-certificate-domainname
	DomainName string `json:"DomainName"`

	// DomainValidationOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-certificate.html#cfn-certificatemanager-certificate-domainvalidationoptions
	DomainValidationOptions []AWSCertificateManagerCertificate_DomainValidationOption `json:"DomainValidationOptions"`

	// SubjectAlternativeNames AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-certificate.html#cfn-certificatemanager-certificate-subjectalternativenames
	SubjectAlternativeNames []string `json:"SubjectAlternativeNames"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-certificatemanager-certificate.html#cfn-certificatemanager-certificate-tags
	Tags []Tag `json:"Tags"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCertificateManagerCertificate) AWSCloudFormationType() string {
	return "AWS::CertificateManager::Certificate"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCertificateManagerCertificate) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCertificateManagerCertificateResources retrieves all AWSCertificateManagerCertificate items from a CloudFormation template
func GetAllAWSCertificateManagerCertificateResources(template *Template) map[string]*AWSCertificateManagerCertificate {

	results := map[string]*AWSCertificateManagerCertificate{}
	for name, resource := range template.Resources {
		result := &AWSCertificateManagerCertificate{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCertificateManagerCertificateWithName retrieves all AWSCertificateManagerCertificate items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSCertificateManagerCertificateWithName(name string, template *Template) (*AWSCertificateManagerCertificate, error) {

	result := &AWSCertificateManagerCertificate{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCertificateManagerCertificate{}, errors.New("resource not found")

}
