package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::IAM::ManagedPolicy AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-managedpolicy.html
type AWSIAMManagedPolicy struct {

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-managedpolicy.html#cfn-iam-managedpolicy-description
	Description string `json:"Description"`

	// Groups AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-managedpolicy.html#cfn-iam-managedpolicy-groups
	Groups []string `json:"Groups"`

	// ManagedPolicyName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-managedpolicy.html#cfn-iam-managedpolicy-managedpolicyname
	ManagedPolicyName string `json:"ManagedPolicyName"`

	// Path AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-managedpolicy.html#cfn-ec2-dhcpoptions-path
	Path string `json:"Path"`

	// PolicyDocument AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-managedpolicy.html#cfn-iam-managedpolicy-policydocument
	PolicyDocument interface{} `json:"PolicyDocument"`

	// Roles AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-managedpolicy.html#cfn-iam-managedpolicy-roles
	Roles []string `json:"Roles"`

	// Users AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-managedpolicy.html#cfn-iam-managedpolicy-users
	Users []string `json:"Users"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIAMManagedPolicy) AWSCloudFormationType() string {
	return "AWS::IAM::ManagedPolicy"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIAMManagedPolicy) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSIAMManagedPolicyResources retrieves all AWSIAMManagedPolicy items from a CloudFormation template
func GetAllAWSIAMManagedPolicy(template *Template) map[string]*AWSIAMManagedPolicy {

	results := map[string]*AWSIAMManagedPolicy{}
	for name, resource := range template.Resources {
		result := &AWSIAMManagedPolicy{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSIAMManagedPolicyWithName retrieves all AWSIAMManagedPolicy items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSIAMManagedPolicy(name string, template *Template) (*AWSIAMManagedPolicy, error) {

	result := &AWSIAMManagedPolicy{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSIAMManagedPolicy{}, errors.New("resource not found")

}
