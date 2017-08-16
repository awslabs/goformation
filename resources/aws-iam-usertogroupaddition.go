package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSIAMUserToGroupAddition AWS CloudFormation Resource (AWS::IAM::UserToGroupAddition)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-addusertogroup.html
type AWSIAMUserToGroupAddition struct {

	// GroupName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-addusertogroup.html#cfn-iam-addusertogroup-groupname
	GroupName string `json:"GroupName"`

	// Users AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-addusertogroup.html#cfn-iam-addusertogroup-users
	Users []string `json:"Users"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIAMUserToGroupAddition) AWSCloudFormationType() string {
	return "AWS::IAM::UserToGroupAddition"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIAMUserToGroupAddition) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSIAMUserToGroupAdditionResources retrieves all AWSIAMUserToGroupAddition items from a CloudFormation template
func (t *Template) GetAllAWSIAMUserToGroupAdditionResources() map[string]*AWSIAMUserToGroupAddition {

	results := map[string]*AWSIAMUserToGroupAddition{}
	for name, resource := range t.Resources {
		result := &AWSIAMUserToGroupAddition{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSIAMUserToGroupAdditionWithName retrieves all AWSIAMUserToGroupAddition items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIAMUserToGroupAdditionWithName(name string) (*AWSIAMUserToGroupAddition, error) {

	result := &AWSIAMUserToGroupAddition{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSIAMUserToGroupAddition{}, errors.New("resource not found")

}
