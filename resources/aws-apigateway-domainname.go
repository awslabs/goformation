package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSApiGatewayDomainName AWS CloudFormation Resource (AWS::ApiGateway::DomainName)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainname.html
type AWSApiGatewayDomainName struct {

	// CertificateArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainname.html#cfn-apigateway-domainname-certificatearn
	CertificateArn string `json:"CertificateArn"`

	// DomainName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-apigateway-domainname.html#cfn-apigateway-domainname-domainname
	DomainName string `json:"DomainName"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSApiGatewayDomainName) AWSCloudFormationType() string {
	return "AWS::ApiGateway::DomainName"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSApiGatewayDomainName) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSApiGatewayDomainNameResources retrieves all AWSApiGatewayDomainName items from a CloudFormation template
func GetAllAWSApiGatewayDomainNameResources(template *Template) map[string]*AWSApiGatewayDomainName {

	results := map[string]*AWSApiGatewayDomainName{}
	for name, resource := range template.Resources {
		result := &AWSApiGatewayDomainName{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSApiGatewayDomainNameWithName retrieves all AWSApiGatewayDomainName items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSApiGatewayDomainNameWithName(name string, template *Template) (*AWSApiGatewayDomainName, error) {

	result := &AWSApiGatewayDomainName{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSApiGatewayDomainName{}, errors.New("resource not found")

}
