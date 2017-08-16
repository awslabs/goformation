package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSDMSCertificate AWS CloudFormation Resource (AWS::DMS::Certificate)
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

// GetAllAWSDMSCertificateResources retrieves all AWSDMSCertificate items from a CloudFormation template
func (t *Template) GetAllAWSDMSCertificateResources() map[string]*AWSDMSCertificate {

	results := map[string]*AWSDMSCertificate{}
	for name, resource := range t.Resources {
		result := &AWSDMSCertificate{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSDMSCertificateWithName retrieves all AWSDMSCertificate items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSDMSCertificateWithName(name string) (*AWSDMSCertificate, error) {

	result := &AWSDMSCertificate{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSDMSCertificate{}, errors.New("resource not found")

}
