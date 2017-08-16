package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSOpsWorksUserProfile AWS CloudFormation Resource (AWS::OpsWorks::UserProfile)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-userprofile.html
type AWSOpsWorksUserProfile struct {

	// AllowSelfManagement AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-userprofile.html#cfn-opsworks-userprofile-allowselfmanagement
	AllowSelfManagement bool `json:"AllowSelfManagement"`

	// IamUserArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-userprofile.html#cfn-opsworks-userprofile-iamuserarn
	IamUserArn string `json:"IamUserArn"`

	// SshPublicKey AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-userprofile.html#cfn-opsworks-userprofile-sshpublickey
	SshPublicKey string `json:"SshPublicKey"`

	// SshUsername AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-opsworks-userprofile.html#cfn-opsworks-userprofile-sshusername
	SshUsername string `json:"SshUsername"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSOpsWorksUserProfile) AWSCloudFormationType() string {
	return "AWS::OpsWorks::UserProfile"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSOpsWorksUserProfile) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSOpsWorksUserProfileResources retrieves all AWSOpsWorksUserProfile items from a CloudFormation template
func (t *Template) GetAllAWSOpsWorksUserProfileResources() map[string]*AWSOpsWorksUserProfile {

	results := map[string]*AWSOpsWorksUserProfile{}
	for name, resource := range t.Resources {
		result := &AWSOpsWorksUserProfile{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSOpsWorksUserProfileWithName retrieves all AWSOpsWorksUserProfile items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSOpsWorksUserProfileWithName(name string) (*AWSOpsWorksUserProfile, error) {

	result := &AWSOpsWorksUserProfile{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSOpsWorksUserProfile{}, errors.New("resource not found")

}
