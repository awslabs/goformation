package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSIAMUser AWS CloudFormation Resource (AWS::IAM::User)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user.html
type AWSIAMUser struct {

	// Groups AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user.html#cfn-iam-user-groups
	Groups []string `json:"Groups"`

	// LoginProfile AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user.html#cfn-iam-user-loginprofile
	LoginProfile AWSIAMUser_LoginProfile `json:"LoginProfile"`

	// ManagedPolicyArns AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user.html#cfn-iam-user-managepolicyarns
	ManagedPolicyArns []string `json:"ManagedPolicyArns"`

	// Path AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user.html#cfn-iam-user-path
	Path string `json:"Path"`

	// Policies AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user.html#cfn-iam-user-policies
	Policies []AWSIAMUser_Policy `json:"Policies"`

	// UserName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user.html#cfn-iam-user-username
	UserName string `json:"UserName"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIAMUser) AWSCloudFormationType() string {
	return "AWS::IAM::User"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIAMUser) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSIAMUserResources retrieves all AWSIAMUser items from a CloudFormation template
func (t *Template) GetAllAWSIAMUserResources() map[string]*AWSIAMUser {

	results := map[string]*AWSIAMUser{}
	for name, resource := range t.Resources {
		result := &AWSIAMUser{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSIAMUserWithName retrieves all AWSIAMUser items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIAMUserWithName(name string) (*AWSIAMUser, error) {

	result := &AWSIAMUser{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSIAMUser{}, errors.New("resource not found")

}
