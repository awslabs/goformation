// Code generated by "go generate". Please don't change this file directly.

package serverless

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Api_DomainConfiguration AWS CloudFormation Resource (AWS::Serverless::Api.DomainConfiguration)
// See: https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/sam-property-api-domainconfiguration.html
type Api_DomainConfiguration struct {

	// BasePath AWS CloudFormation Property
	// Required: false
	// See: https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/sam-property-api-domainconfiguration.html#sam-api-domainconfiguration-basepath
	BasePath []string `json:"BasePath,omitempty"`

	// CertificateArn AWS CloudFormation Property
	// Required: true
	// See: https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/sam-property-api-domainconfiguration.html#sam-api-domainconfiguration-certificatearn
	CertificateArn string `json:"CertificateArn"`

	// DomainName AWS CloudFormation Property
	// Required: true
	// See: https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/sam-property-api-domainconfiguration.html#sam-api-domainconfiguration-domainname
	DomainName string `json:"DomainName"`

	// EndpointConfiguration AWS CloudFormation Property
	// Required: false
	// See: https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/sam-property-api-domainconfiguration.html#sam-api-domainconfiguration-endpointconfiguration
	EndpointConfiguration *string `json:"EndpointConfiguration,omitempty"`

	// MutualTlsAuthentication AWS CloudFormation Property
	// Required: false
	// See: https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/sam-property-api-domainconfiguration.html#sam-api-domainconfiguration-mutualtlsauthentication
	MutualTlsAuthentication *Api_MutualTlsAuthentication `json:"MutualTlsAuthentication,omitempty"`

	// OwnershipVerificationCertificateArn AWS CloudFormation Property
	// Required: false
	// See: https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/sam-property-api-domainconfiguration.html#sam-api-domainconfiguration-ownershipverificationcertificatearn
	OwnershipVerificationCertificateArn *string `json:"OwnershipVerificationCertificateArn,omitempty"`

	// Route53 AWS CloudFormation Property
	// Required: false
	// See: https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/sam-property-api-domainconfiguration.html#sam-api-domainconfiguration-route53
	Route53 *Api_Route53Configuration `json:"Route53,omitempty"`

	// SecurityPolicy AWS CloudFormation Property
	// Required: false
	// See: https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/sam-property-api-domainconfiguration.html#sam-api-domainconfiguration-securitypolicy
	SecurityPolicy *string `json:"SecurityPolicy,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *Api_DomainConfiguration) AWSCloudFormationType() string {
	return "AWS::Serverless::Api.DomainConfiguration"
}
