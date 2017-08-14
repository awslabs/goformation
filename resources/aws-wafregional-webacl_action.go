package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::WAFRegional::WebACL.Action AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-webacl-action.html
type AWSWAFRegionalWebACL_Action struct {

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-webacl-action.html#cfn-wafregional-webacl-action-type

	Type string `json:"Type"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFRegionalWebACL_Action) AWSCloudFormationType() string {
	return "AWS::WAFRegional::WebACL.Action"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFRegionalWebACL_Action) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSWAFRegionalWebACL_ActionResources retrieves all AWSWAFRegionalWebACL_Action items from a CloudFormation template
func GetAllAWSWAFRegionalWebACL_Action(template *Template) map[string]*AWSWAFRegionalWebACL_Action {

	results := map[string]*AWSWAFRegionalWebACL_Action{}
	for name, resource := range template.Resources {
		result := &AWSWAFRegionalWebACL_Action{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSWAFRegionalWebACL_ActionWithName retrieves all AWSWAFRegionalWebACL_Action items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSWAFRegionalWebACL_Action(name string, template *Template) (*AWSWAFRegionalWebACL_Action, error) {

	result := &AWSWAFRegionalWebACL_Action{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSWAFRegionalWebACL_Action{}, errors.New("resource not found")

}
