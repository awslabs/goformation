package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Config::ConfigRule.Scope AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-scope.html
type AWSConfigConfigRule_Scope struct {

	// ComplianceResourceId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-scope.html#cfn-config-configrule-scope-complianceresourceid

	ComplianceResourceId string `json:"ComplianceResourceId"`

	// ComplianceResourceTypes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-scope.html#cfn-config-configrule-scope-complianceresourcetypes

	ComplianceResourceTypes []string `json:"ComplianceResourceTypes"`

	// TagKey AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-scope.html#cfn-config-configrule-scope-tagkey

	TagKey string `json:"TagKey"`

	// TagValue AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-scope.html#cfn-config-configrule-scope-tagvalue

	TagValue string `json:"TagValue"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSConfigConfigRule_Scope) AWSCloudFormationType() string {
	return "AWS::Config::ConfigRule.Scope"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSConfigConfigRule_Scope) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSConfigConfigRule_ScopeResources retrieves all AWSConfigConfigRule_Scope items from a CloudFormation template
func GetAllAWSConfigConfigRule_Scope(template *Template) map[string]*AWSConfigConfigRule_Scope {

	results := map[string]*AWSConfigConfigRule_Scope{}
	for name, resource := range template.Resources {
		result := &AWSConfigConfigRule_Scope{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSConfigConfigRule_ScopeWithName retrieves all AWSConfigConfigRule_Scope items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSConfigConfigRule_Scope(name string, template *Template) (*AWSConfigConfigRule_Scope, error) {

	result := &AWSConfigConfigRule_Scope{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSConfigConfigRule_Scope{}, errors.New("resource not found")

}
