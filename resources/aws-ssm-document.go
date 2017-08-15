package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSSSMDocument AWS CloudFormation Resource (AWS::SSM::Document)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-document.html
type AWSSSMDocument struct {

	// Content AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-document.html#cfn-ssm-document-content
	Content interface{} `json:"Content"`

	// DocumentType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-document.html#cfn-ssm-document-documenttype
	DocumentType string `json:"DocumentType"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSSMDocument) AWSCloudFormationType() string {
	return "AWS::SSM::Document"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSSSMDocument) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSSSMDocumentResources retrieves all AWSSSMDocument items from a CloudFormation template
func GetAllAWSSSMDocumentResources(template *Template) map[string]*AWSSSMDocument {

	results := map[string]*AWSSSMDocument{}
	for name, resource := range template.Resources {
		result := &AWSSSMDocument{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSSSMDocumentWithName retrieves all AWSSSMDocument items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSSSMDocumentWithName(name string, template *Template) (*AWSSSMDocument, error) {

	result := &AWSSSMDocument{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSSSMDocument{}, errors.New("resource not found")

}
