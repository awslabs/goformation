package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::IAM::Role.Policy AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-policy.html
type AWSIAMRole_Policy struct {

	// PolicyDocument AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-policy.html#cfn-iam-policies-policydocument
	PolicyDocument interface{} `json:"PolicyDocument"`

	// PolicyName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-policy.html#cfn-iam-policies-policyname
	PolicyName string `json:"PolicyName"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIAMRole_Policy) AWSCloudFormationType() string {
	return "AWS::IAM::Role.Policy"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIAMRole_Policy) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSIAMRole_PolicyResources retrieves all AWSIAMRole_Policy items from a CloudFormation template
func GetAllAWSIAMRole_Policy(template *Template) map[string]*AWSIAMRole_Policy {

	results := map[string]*AWSIAMRole_Policy{}
	for name, resource := range template.Resources {
		result := &AWSIAMRole_Policy{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSIAMRole_PolicyWithName retrieves all AWSIAMRole_Policy items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSIAMRole_Policy(name string, template *Template) (*AWSIAMRole_Policy, error) {

	result := &AWSIAMRole_Policy{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSIAMRole_Policy{}, errors.New("resource not found")

}
