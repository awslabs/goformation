package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::KMS::Alias AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-alias.html
type AWSKMSAlias struct {

	// AliasName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-alias.html#cfn-kms-alias-aliasname
	AliasName string `json:"AliasName"`

	// TargetKeyId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-kms-alias.html#cfn-kms-alias-targetkeyid
	TargetKeyId string `json:"TargetKeyId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSKMSAlias) AWSCloudFormationType() string {
	return "AWS::KMS::Alias"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSKMSAlias) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSKMSAliasResources retrieves all AWSKMSAlias items from a CloudFormation template
func GetAllAWSKMSAlias(template *Template) map[string]*AWSKMSAlias {

	results := map[string]*AWSKMSAlias{}
	for name, resource := range template.Resources {
		result := &AWSKMSAlias{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSKMSAliasWithName retrieves all AWSKMSAlias items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSKMSAlias(name string, template *Template) (*AWSKMSAlias, error) {

	result := &AWSKMSAlias{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSKMSAlias{}, errors.New("resource not found")

}
