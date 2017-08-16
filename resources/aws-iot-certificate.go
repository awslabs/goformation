package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSIoTCertificate AWS CloudFormation Resource (AWS::IoT::Certificate)
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

// GetAllAWSIoTCertificateResources retrieves all AWSIoTCertificate items from a CloudFormation template
func (t *Template) GetAllAWSIoTCertificateResources() map[string]*AWSIoTCertificate {

	results := map[string]*AWSIoTCertificate{}
	for name, resource := range t.Resources {
		result := &AWSIoTCertificate{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSIoTCertificateWithName retrieves all AWSIoTCertificate items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIoTCertificateWithName(name string) (*AWSIoTCertificate, error) {

	result := &AWSIoTCertificate{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSIoTCertificate{}, errors.New("resource not found")

}
