package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSWAFWebACL AWS CloudFormation Resource (AWS::WAF::WebACL)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-webacl.html
type AWSWAFWebACL struct {

	// DefaultAction AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-webacl.html#cfn-waf-webacl-defaultaction
	DefaultAction AWSWAFWebACL_WafAction `json:"DefaultAction"`

	// MetricName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-webacl.html#cfn-waf-webacl-metricname
	MetricName string `json:"MetricName"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-webacl.html#cfn-waf-webacl-name
	Name string `json:"Name"`

	// Rules AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-waf-webacl.html#cfn-waf-webacl-rules
	Rules []AWSWAFWebACL_ActivatedRule `json:"Rules"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFWebACL) AWSCloudFormationType() string {
	return "AWS::WAF::WebACL"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSWAFWebACL) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSWAFWebACLResources retrieves all AWSWAFWebACL items from a CloudFormation template
func (t *Template) GetAllAWSWAFWebACLResources() map[string]*AWSWAFWebACL {

	results := map[string]*AWSWAFWebACL{}
	for name, resource := range t.Resources {
		result := &AWSWAFWebACL{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSWAFWebACLWithName retrieves all AWSWAFWebACL items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSWAFWebACLWithName(name string) (*AWSWAFWebACL, error) {

	result := &AWSWAFWebACL{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSWAFWebACL{}, errors.New("resource not found")

}
