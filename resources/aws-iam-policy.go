package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::IAM::Policy AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-policy.html
type AWSIAMPolicy struct {

	// Groups AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-policy.html#cfn-iam-policy-groups
	Groups []string `json:"Groups"`

	// PolicyDocument AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-policy.html#cfn-iam-policy-policydocument
	PolicyDocument interface{} `json:"PolicyDocument"`

	// PolicyName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-policy.html#cfn-iam-policy-policyname
	PolicyName string `json:"PolicyName"`

	// Roles AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-policy.html#cfn-iam-policy-roles
	Roles []string `json:"Roles"`

	// Users AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-policy.html#cfn-iam-policy-users
	Users []string `json:"Users"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIAMPolicy) AWSCloudFormationType() string {
	return "AWS::IAM::Policy"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIAMPolicy) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSIAMPolicyResources retrieves all AWSIAMPolicy items from a CloudFormation template
func GetAllAWSIAMPolicyResources(template *Template) map[string]*AWSIAMPolicy {

	results := map[string]*AWSIAMPolicy{}
	for name, resource := range template.Resources {
		result := &AWSIAMPolicy{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSIAMPolicyWithName retrieves all AWSIAMPolicy items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSIAMPolicyWithName(name string, template *Template) (*AWSIAMPolicy, error) {

	result := &AWSIAMPolicy{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSIAMPolicy{}, errors.New("resource not found")

}
