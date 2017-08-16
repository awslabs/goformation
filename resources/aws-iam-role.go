package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSIAMRole AWS CloudFormation Resource (AWS::IAM::Role)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html
type AWSIAMRole struct {

	// AssumeRolePolicyDocument AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html#cfn-iam-role-assumerolepolicydocument
	AssumeRolePolicyDocument interface{} `json:"AssumeRolePolicyDocument"`

	// ManagedPolicyArns AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html#cfn-iam-role-managepolicyarns
	ManagedPolicyArns []string `json:"ManagedPolicyArns"`

	// Path AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html#cfn-iam-role-path
	Path string `json:"Path"`

	// Policies AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html#cfn-iam-role-policies
	Policies []AWSIAMRole_Policy `json:"Policies"`

	// RoleName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-role.html#cfn-iam-role-rolename
	RoleName string `json:"RoleName"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIAMRole) AWSCloudFormationType() string {
	return "AWS::IAM::Role"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIAMRole) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSIAMRoleResources retrieves all AWSIAMRole items from a CloudFormation template
func (t *Template) GetAllAWSIAMRoleResources() map[string]*AWSIAMRole {

	results := map[string]*AWSIAMRole{}
	for name, resource := range t.Resources {
		result := &AWSIAMRole{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSIAMRoleWithName retrieves all AWSIAMRole items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIAMRoleWithName(name string) (*AWSIAMRole, error) {

	result := &AWSIAMRole{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSIAMRole{}, errors.New("resource not found")

}
