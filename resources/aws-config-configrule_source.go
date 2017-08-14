package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Config::ConfigRule.Source AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-source.html
type AWSConfigConfigRule_Source struct {

	// Owner AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-source.html#cfn-config-configrule-source-owner

	Owner string `json:"Owner"`

	// SourceDetails AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-source.html#cfn-config-configrule-source-sourcedetails

	SourceDetails []AWSConfigConfigRule_SourceDetail `json:"SourceDetails"`

	// SourceIdentifier AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-source.html#cfn-config-configrule-source-sourceidentifier

	SourceIdentifier string `json:"SourceIdentifier"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSConfigConfigRule_Source) AWSCloudFormationType() string {
	return "AWS::Config::ConfigRule.Source"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSConfigConfigRule_Source) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSConfigConfigRule_SourceResources retrieves all AWSConfigConfigRule_Source items from a CloudFormation template
func GetAllAWSConfigConfigRule_Source(template *Template) map[string]*AWSConfigConfigRule_Source {

	results := map[string]*AWSConfigConfigRule_Source{}
	for name, resource := range template.Resources {
		result := &AWSConfigConfigRule_Source{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSConfigConfigRule_SourceWithName retrieves all AWSConfigConfigRule_Source items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSConfigConfigRule_Source(name string, template *Template) (*AWSConfigConfigRule_Source, error) {

	result := &AWSConfigConfigRule_Source{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSConfigConfigRule_Source{}, errors.New("resource not found")

}
