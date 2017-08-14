package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::WAFRegional::WebACL AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-webacl.html
type AWSWAFRegionalWebACL struct {

	// DefaultAction AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-webacl.html#cfn-wafregional-webacl-defaultaction

	DefaultAction AWSWAFRegionalWebACL_Action `json:"DefaultAction"`

	// MetricName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-webacl.html#cfn-wafregional-webacl-metricname

	MetricName string `json:"MetricName"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-webacl.html#cfn-wafregional-webacl-name

	Name string `json:"Name"`

	// Rules AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-wafregional-webacl.html#cfn-wafregional-webacl-rules

	Rules []AWSWAFRegionalWebACL_Rule `json:"Rules"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFRegionalWebACL) AWSCloudFormationType() string {
	return "AWS::WAFRegional::WebACL"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFRegionalWebACL) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSWAFRegionalWebACLResources retrieves all AWSWAFRegionalWebACL items from a CloudFormation template
func GetAllAWSWAFRegionalWebACL(template *Template) map[string]*AWSWAFRegionalWebACL {

	results := map[string]*AWSWAFRegionalWebACL{}
	for name, resource := range template.Resources {
		result := &AWSWAFRegionalWebACL{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSWAFRegionalWebACLWithName retrieves all AWSWAFRegionalWebACL items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSWAFRegionalWebACL(name string, template *Template) (*AWSWAFRegionalWebACL, error) {

	result := &AWSWAFRegionalWebACL{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSWAFRegionalWebACL{}, errors.New("resource not found")

}
