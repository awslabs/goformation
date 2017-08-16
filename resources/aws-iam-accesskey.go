package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSIAMAccessKey AWS CloudFormation Resource (AWS::IAM::AccessKey)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-accesskey.html
type AWSIAMAccessKey struct {

	// Serial AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-accesskey.html#cfn-iam-accesskey-serial
	Serial int `json:"Serial"`

	// Status AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-accesskey.html#cfn-iam-accesskey-status
	Status string `json:"Status"`

	// UserName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iam-accesskey.html#cfn-iam-accesskey-username
	UserName string `json:"UserName"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIAMAccessKey) AWSCloudFormationType() string {
	return "AWS::IAM::AccessKey"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIAMAccessKey) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSIAMAccessKeyResources retrieves all AWSIAMAccessKey items from a CloudFormation template
func (t *Template) GetAllAWSIAMAccessKeyResources() map[string]*AWSIAMAccessKey {

	results := map[string]*AWSIAMAccessKey{}
	for name, resource := range t.Resources {
		result := &AWSIAMAccessKey{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSIAMAccessKeyWithName retrieves all AWSIAMAccessKey items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSIAMAccessKeyWithName(name string) (*AWSIAMAccessKey, error) {

	result := &AWSIAMAccessKey{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSIAMAccessKey{}, errors.New("resource not found")

}
