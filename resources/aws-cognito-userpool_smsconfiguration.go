package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Cognito::UserPool.SmsConfiguration AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-smsconfiguration.html
type AWSCognitoUserPool_SmsConfiguration struct {

	// ExternalId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-smsconfiguration.html#cfn-cognito-userpool-smsconfiguration-externalid
	ExternalId string `json:"ExternalId"`

	// SnsCallerArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-userpool-smsconfiguration.html#cfn-cognito-userpool-smsconfiguration-snscallerarn
	SnsCallerArn string `json:"SnsCallerArn"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoUserPool_SmsConfiguration) AWSCloudFormationType() string {
	return "AWS::Cognito::UserPool.SmsConfiguration"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCognitoUserPool_SmsConfiguration) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
