package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSKMSKey AWS CloudFormation Resource (AWS::KMS::Key)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html
type AWSKMSKey struct {

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-description
	Description string `json:"Description"`

	// EnableKeyRotation AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-enablekeyrotation
	EnableKeyRotation bool `json:"EnableKeyRotation"`

	// Enabled AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-enabled
	Enabled bool `json:"Enabled"`

	// KeyPolicy AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-keypolicy
	KeyPolicy interface{} `json:"KeyPolicy"`

	// KeyUsage AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-key.html#cfn-kms-key-keyusage
	KeyUsage string `json:"KeyUsage"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKMSKey) AWSCloudFormationType() string {
	return "AWS::KMS::Key"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKMSKey) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSKMSKeyResources retrieves all AWSKMSKey items from a CloudFormation template
func (t *Template) GetAllAWSKMSKeyResources() map[string]*AWSKMSKey {

	results := map[string]*AWSKMSKey{}
	for name, resource := range t.Resources {
		result := &AWSKMSKey{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSKMSKeyWithName retrieves all AWSKMSKey items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSKMSKeyWithName(name string) (*AWSKMSKey, error) {

	result := &AWSKMSKey{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSKMSKey{}, errors.New("resource not found")

}
