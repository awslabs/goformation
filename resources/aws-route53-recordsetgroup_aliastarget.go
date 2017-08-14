package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Route53::RecordSetGroup.AliasTarget AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-aliastarget.html
type AWSRoute53RecordSetGroup_AliasTarget struct {

	// DNSName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-aliastarget.html#cfn-route53-aliastarget-dnshostname

	DNSName string `json:"DNSName"`

	// EvaluateTargetHealth AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-aliastarget.html#cfn-route53-aliastarget-evaluatetargethealth

	EvaluateTargetHealth bool `json:"EvaluateTargetHealth"`

	// HostedZoneId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-aliastarget.html#cfn-route53-aliastarget-hostedzoneid

	HostedZoneId string `json:"HostedZoneId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRoute53RecordSetGroup_AliasTarget) AWSCloudFormationType() string {
	return "AWS::Route53::RecordSetGroup.AliasTarget"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSRoute53RecordSetGroup_AliasTarget) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSRoute53RecordSetGroup_AliasTargetResources retrieves all AWSRoute53RecordSetGroup_AliasTarget items from a CloudFormation template
func GetAllAWSRoute53RecordSetGroup_AliasTarget(template *Template) map[string]*AWSRoute53RecordSetGroup_AliasTarget {

	results := map[string]*AWSRoute53RecordSetGroup_AliasTarget{}
	for name, resource := range template.Resources {
		result := &AWSRoute53RecordSetGroup_AliasTarget{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSRoute53RecordSetGroup_AliasTargetWithName retrieves all AWSRoute53RecordSetGroup_AliasTarget items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSRoute53RecordSetGroup_AliasTarget(name string, template *Template) (*AWSRoute53RecordSetGroup_AliasTarget, error) {

	result := &AWSRoute53RecordSetGroup_AliasTarget{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSRoute53RecordSetGroup_AliasTarget{}, errors.New("resource not found")

}
