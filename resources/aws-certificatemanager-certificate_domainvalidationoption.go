package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::CertificateManager::Certificate.DomainValidationOption AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-certificatemanager-certificate-domainvalidationoption.html
type AWSCertificateManagerCertificate_DomainValidationOption struct {

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
func (r *AWSCertificateManagerCertificate_DomainValidationOption) AWSCloudFormationType() string {
	return "AWS::CertificateManager::Certificate.DomainValidationOption"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCertificateManagerCertificate_DomainValidationOption) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCertificateManagerCertificate_DomainValidationOptionResources retrieves all AWSCertificateManagerCertificate_DomainValidationOption items from a CloudFormation template
func GetAllAWSCertificateManagerCertificate_DomainValidationOption(template *Template) map[string]*AWSCertificateManagerCertificate_DomainValidationOption {

	results := map[string]*AWSCertificateManagerCertificate_DomainValidationOption{}
	for name, resource := range template.Resources {
		result := &AWSCertificateManagerCertificate_DomainValidationOption{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCertificateManagerCertificate_DomainValidationOptionWithName retrieves all AWSCertificateManagerCertificate_DomainValidationOption items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCertificateManagerCertificate_DomainValidationOption(name string, template *Template) (*AWSCertificateManagerCertificate_DomainValidationOption, error) {

	result := &AWSCertificateManagerCertificate_DomainValidationOption{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCertificateManagerCertificate_DomainValidationOption{}, errors.New("resource not found")

}
