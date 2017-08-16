package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSIAMGroup AWS CloudFormation Resource (AWS::IAM::Group)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-group.html
type AWSIAMGroup struct {

	// GroupName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-group.html#cfn-iam-group-groupname
	GroupName string `json:"GroupName"`

	// ManagedPolicyArns AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-group.html#cfn-iam-group-managepolicyarns
	ManagedPolicyArns []string `json:"ManagedPolicyArns"`

	// Path AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-group.html#cfn-iam-group-path
	Path string `json:"Path"`

	// Policies AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-group.html#cfn-iam-group-policies
	Policies []AWSIAMGroup_Policy `json:"Policies"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIAMGroup) AWSCloudFormationType() string {
	return "AWS::IAM::Group"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIAMGroup) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSIAMGroupResources retrieves all AWSIAMGroup items from a CloudFormation template
func (t *Template) GetAllAWSIAMGroupResources() map[string]*AWSIAMGroup {

	results := map[string]*AWSIAMGroup{}
	for name, resource := range t.Resources {
		result := &AWSIAMGroup{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSIAMGroupWithName retrieves all AWSIAMGroup items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIAMGroupWithName(name string) (*AWSIAMGroup, error) {

	result := &AWSIAMGroup{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSIAMGroup{}, errors.New("resource not found")

}
