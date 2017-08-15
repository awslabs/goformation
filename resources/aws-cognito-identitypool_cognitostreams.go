package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Cognito::IdentityPool.CognitoStreams AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypool-cognitostreams.html
type AWSCognitoIdentityPool_CognitoStreams struct {

	// RoleArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypool-cognitostreams.html#cfn-cognito-identitypool-cognitostreams-rolearn
	RoleArn string `json:"RoleArn"`

	// StreamName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypool-cognitostreams.html#cfn-cognito-identitypool-cognitostreams-streamname
	StreamName string `json:"StreamName"`

	// StreamingStatus AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypool-cognitostreams.html#cfn-cognito-identitypool-cognitostreams-streamingstatus
	StreamingStatus string `json:"StreamingStatus"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoIdentityPool_CognitoStreams) AWSCloudFormationType() string {
	return "AWS::Cognito::IdentityPool.CognitoStreams"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCognitoIdentityPool_CognitoStreams) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCognitoIdentityPool_CognitoStreamsResources retrieves all AWSCognitoIdentityPool_CognitoStreams items from a CloudFormation template
func GetAllAWSCognitoIdentityPool_CognitoStreams(template *Template) map[string]*AWSCognitoIdentityPool_CognitoStreams {

	results := map[string]*AWSCognitoIdentityPool_CognitoStreams{}
	for name, resource := range template.Resources {
		result := &AWSCognitoIdentityPool_CognitoStreams{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCognitoIdentityPool_CognitoStreamsWithName retrieves all AWSCognitoIdentityPool_CognitoStreams items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCognitoIdentityPool_CognitoStreams(name string, template *Template) (*AWSCognitoIdentityPool_CognitoStreams, error) {

	result := &AWSCognitoIdentityPool_CognitoStreams{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCognitoIdentityPool_CognitoStreams{}, errors.New("resource not found")

}
