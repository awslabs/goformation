package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::IAM::Group.Policy AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-policy.html
type AWSIAMGroup_Policy struct {

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
func (r *AWSIAMGroup_Policy) AWSCloudFormationType() string {
	return "AWS::IAM::Group.Policy"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIAMGroup_Policy) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSIAMGroup_PolicyResources retrieves all AWSIAMGroup_Policy items from a CloudFormation template
func GetAllAWSIAMGroup_Policy(template *Template) map[string]*AWSIAMGroup_Policy {

	results := map[string]*AWSIAMGroup_Policy{}
	for name, resource := range template.Resources {
		result := &AWSIAMGroup_Policy{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSIAMGroup_PolicyWithName retrieves all AWSIAMGroup_Policy items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSIAMGroup_Policy(name string, template *Template) (*AWSIAMGroup_Policy, error) {

	result := &AWSIAMGroup_Policy{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSIAMGroup_Policy{}, errors.New("resource not found")

}
