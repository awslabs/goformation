package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSIAMInstanceProfile AWS CloudFormation Resource (AWS::IAM::InstanceProfile)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-instanceprofile.html
type AWSIAMInstanceProfile struct {

	// InstanceProfileName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-instanceprofile.html#cfn-iam-instanceprofile-instanceprofilename
	InstanceProfileName string `json:"InstanceProfileName"`

	// Path AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-instanceprofile.html#cfn-iam-instanceprofile-path
	Path string `json:"Path"`

	// Roles AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iam-instanceprofile.html#cfn-iam-instanceprofile-roles
	Roles []string `json:"Roles"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIAMInstanceProfile) AWSCloudFormationType() string {
	return "AWS::IAM::InstanceProfile"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIAMInstanceProfile) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSIAMInstanceProfileResources retrieves all AWSIAMInstanceProfile items from a CloudFormation template
func (t *Template) GetAllAWSIAMInstanceProfileResources() map[string]*AWSIAMInstanceProfile {

	results := map[string]*AWSIAMInstanceProfile{}
	for name, resource := range t.Resources {
		result := &AWSIAMInstanceProfile{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSIAMInstanceProfileWithName retrieves all AWSIAMInstanceProfile items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIAMInstanceProfileWithName(name string) (*AWSIAMInstanceProfile, error) {

	result := &AWSIAMInstanceProfile{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSIAMInstanceProfile{}, errors.New("resource not found")

}
