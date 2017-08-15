package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Cognito::IdentityPool.PushSync AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypool-pushsync.html
type AWSCognitoIdentityPool_PushSync struct {

	// ApplicationArns AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypool-pushsync.html#cfn-cognito-identitypool-pushsync-applicationarns

	ApplicationArns []string `json:"ApplicationArns"`

	// RoleArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypool-pushsync.html#cfn-cognito-identitypool-pushsync-rolearn

	RoleArn string `json:"RoleArn"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCognitoIdentityPool_PushSync) AWSCloudFormationType() string {
	return "AWS::Cognito::IdentityPool.PushSync"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSCognitoIdentityPool_PushSync) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSCognitoIdentityPool_PushSyncResources retrieves all AWSCognitoIdentityPool_PushSync items from a CloudFormation template
func GetAllAWSCognitoIdentityPool_PushSync(template *Template) map[string]*AWSCognitoIdentityPool_PushSync {

	results := map[string]*AWSCognitoIdentityPool_PushSync{}
	for name, resource := range template.Resources {
		result := &AWSCognitoIdentityPool_PushSync{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSCognitoIdentityPool_PushSyncWithName retrieves all AWSCognitoIdentityPool_PushSync items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSCognitoIdentityPool_PushSync(name string, template *Template) (*AWSCognitoIdentityPool_PushSync, error) {

	result := &AWSCognitoIdentityPool_PushSync{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSCognitoIdentityPool_PushSync{}, errors.New("resource not found")

}
