package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Route53::RecordSet.AliasTarget AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-aliastarget.html
type AWSRoute53RecordSet_AliasTarget struct {

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
func (r *AWSRoute53RecordSet_AliasTarget) AWSCloudFormationType() string {
	return "AWS::Route53::RecordSet.AliasTarget"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSRoute53RecordSet_AliasTarget) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSRoute53RecordSet_AliasTargetResources retrieves all AWSRoute53RecordSet_AliasTarget items from a CloudFormation template
func GetAllAWSRoute53RecordSet_AliasTarget(template *Template) map[string]*AWSRoute53RecordSet_AliasTarget {

	results := map[string]*AWSRoute53RecordSet_AliasTarget{}
	for name, resource := range template.Resources {
		result := &AWSRoute53RecordSet_AliasTarget{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSRoute53RecordSet_AliasTargetWithName retrieves all AWSRoute53RecordSet_AliasTarget items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSRoute53RecordSet_AliasTarget(name string, template *Template) (*AWSRoute53RecordSet_AliasTarget, error) {

	result := &AWSRoute53RecordSet_AliasTarget{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSRoute53RecordSet_AliasTarget{}, errors.New("resource not found")

}
