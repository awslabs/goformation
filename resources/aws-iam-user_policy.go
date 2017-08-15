package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::IAM::User.Policy AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-policy.html
type AWSIAMUser_Policy struct {

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
func (r *AWSIAMUser_Policy) AWSCloudFormationType() string {
	return "AWS::IAM::User.Policy"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIAMUser_Policy) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSIAMUser_PolicyResources retrieves all AWSIAMUser_Policy items from a CloudFormation template
func GetAllAWSIAMUser_Policy(template *Template) map[string]*AWSIAMUser_Policy {

	results := map[string]*AWSIAMUser_Policy{}
	for name, resource := range template.Resources {
		result := &AWSIAMUser_Policy{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSIAMUser_PolicyWithName retrieves all AWSIAMUser_Policy items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSIAMUser_Policy(name string, template *Template) (*AWSIAMUser_Policy, error) {

	result := &AWSIAMUser_Policy{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSIAMUser_Policy{}, errors.New("resource not found")

}
