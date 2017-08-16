package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSConfigConfigRule AWS CloudFormation Resource (AWS::Config::ConfigRule)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configrule.html
type AWSConfigConfigRule struct {

	// ConfigRuleName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configrule.html#cfn-config-configrule-configrulename
	ConfigRuleName string `json:"ConfigRuleName"`

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configrule.html#cfn-config-configrule-description
	Description string `json:"Description"`

	// InputParameters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configrule.html#cfn-config-configrule-inputparameters
	InputParameters interface{} `json:"InputParameters"`

	// MaximumExecutionFrequency AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configrule.html#cfn-config-configrule-maximumexecutionfrequency
	MaximumExecutionFrequency string `json:"MaximumExecutionFrequency"`

	// Scope AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configrule.html#cfn-config-configrule-scope
	Scope AWSConfigConfigRule_Scope `json:"Scope"`

	// Source AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-configrule.html#cfn-config-configrule-source
	Source AWSConfigConfigRule_Source `json:"Source"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSConfigConfigRule) AWSCloudFormationType() string {
	return "AWS::Config::ConfigRule"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSConfigConfigRule) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSConfigConfigRuleResources retrieves all AWSConfigConfigRule items from a CloudFormation template
func (t *Template) GetAllAWSConfigConfigRuleResources() map[string]*AWSConfigConfigRule {

	results := map[string]*AWSConfigConfigRule{}
	for name, resource := range t.Resources {
		result := &AWSConfigConfigRule{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSConfigConfigRuleWithName retrieves all AWSConfigConfigRule items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSConfigConfigRuleWithName(name string) (*AWSConfigConfigRule, error) {

	result := &AWSConfigConfigRule{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSConfigConfigRule{}, errors.New("resource not found")

}
