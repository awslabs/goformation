package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSApiGatewayClientCertificate AWS CloudFormation Resource (AWS::ApiGateway::ClientCertificate)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-clientcertificate.html
type AWSApiGatewayClientCertificate struct {

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-clientcertificate.html#cfn-apigateway-clientcertificate-description
	Description string `json:"Description"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApiGatewayClientCertificate) AWSCloudFormationType() string {
	return "AWS::ApiGateway::ClientCertificate"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSApiGatewayClientCertificate) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSApiGatewayClientCertificateResources retrieves all AWSApiGatewayClientCertificate items from a CloudFormation template
func GetAllAWSApiGatewayClientCertificateResources(template *Template) map[string]*AWSApiGatewayClientCertificate {

	results := map[string]*AWSApiGatewayClientCertificate{}
	for name, resource := range template.Resources {
		result := &AWSApiGatewayClientCertificate{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSApiGatewayClientCertificateWithName retrieves all AWSApiGatewayClientCertificate items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSApiGatewayClientCertificateWithName(name string, template *Template) (*AWSApiGatewayClientCertificate, error) {

	result := &AWSApiGatewayClientCertificate{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSApiGatewayClientCertificate{}, errors.New("resource not found")

}
