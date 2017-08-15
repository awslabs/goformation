package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::IAM::User.LoginProfile AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user-loginprofile.html
type AWSIAMUser_LoginProfile struct {

	// Password AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user-loginprofile.html#cfn-iam-user-loginprofile-password
	Password string `json:"Password"`

	// PasswordResetRequired AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-user-loginprofile.html#cfn-iam-user-loginprofile-passwordresetrequired
	PasswordResetRequired bool `json:"PasswordResetRequired"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIAMUser_LoginProfile) AWSCloudFormationType() string {
	return "AWS::IAM::User.LoginProfile"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIAMUser_LoginProfile) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSIAMUser_LoginProfileResources retrieves all AWSIAMUser_LoginProfile items from a CloudFormation template
func GetAllAWSIAMUser_LoginProfile(template *Template) map[string]*AWSIAMUser_LoginProfile {

	results := map[string]*AWSIAMUser_LoginProfile{}
	for name, resource := range template.Resources {
		result := &AWSIAMUser_LoginProfile{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSIAMUser_LoginProfileWithName retrieves all AWSIAMUser_LoginProfile items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSIAMUser_LoginProfile(name string, template *Template) (*AWSIAMUser_LoginProfile, error) {

	result := &AWSIAMUser_LoginProfile{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSIAMUser_LoginProfile{}, errors.New("resource not found")

}
