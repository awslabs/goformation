package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::WAF::WebACL.WafAction AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-webacl-action.html
type AWSWAFWebACL_WafAction struct {

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-waf-webacl-action.html#cfn-waf-webacl-action-type
	Type string `json:"Type"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFWebACL_WafAction) AWSCloudFormationType() string {
	return "AWS::WAF::WebACL.WafAction"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFWebACL_WafAction) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSWAFWebACL_WafActionResources retrieves all AWSWAFWebACL_WafAction items from a CloudFormation template
func GetAllAWSWAFWebACL_WafAction(template *Template) map[string]*AWSWAFWebACL_WafAction {

	results := map[string]*AWSWAFWebACL_WafAction{}
	for name, resource := range template.Resources {
		result := &AWSWAFWebACL_WafAction{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSWAFWebACL_WafActionWithName retrieves all AWSWAFWebACL_WafAction items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSWAFWebACL_WafAction(name string, template *Template) (*AWSWAFWebACL_WafAction, error) {

	result := &AWSWAFWebACL_WafAction{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSWAFWebACL_WafAction{}, errors.New("resource not found")

}
