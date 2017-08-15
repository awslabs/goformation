package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Route53::RecordSetGroup AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordsetgroup.html
type AWSRoute53RecordSetGroup struct {

	// Comment AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordsetgroup.html#cfn-route53-recordsetgroup-comment
	Comment string `json:"Comment"`

	// HostedZoneId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordsetgroup.html#cfn-route53-recordsetgroup-hostedzoneid
	HostedZoneId string `json:"HostedZoneId"`

	// HostedZoneName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordsetgroup.html#cfn-route53-recordsetgroup-hostedzonename
	HostedZoneName string `json:"HostedZoneName"`

	// RecordSets AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordsetgroup.html#cfn-route53-recordsetgroup-recordsets
	RecordSets []AWSRoute53RecordSetGroup_RecordSet `json:"RecordSets"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRoute53RecordSetGroup) AWSCloudFormationType() string {
	return "AWS::Route53::RecordSetGroup"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSRoute53RecordSetGroup) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSRoute53RecordSetGroupResources retrieves all AWSRoute53RecordSetGroup items from a CloudFormation template
func GetAllAWSRoute53RecordSetGroup(template *Template) map[string]*AWSRoute53RecordSetGroup {

	results := map[string]*AWSRoute53RecordSetGroup{}
	for name, resource := range template.Resources {
		result := &AWSRoute53RecordSetGroup{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSRoute53RecordSetGroupWithName retrieves all AWSRoute53RecordSetGroup items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSRoute53RecordSetGroup(name string, template *Template) (*AWSRoute53RecordSetGroup, error) {

	result := &AWSRoute53RecordSetGroup{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSRoute53RecordSetGroup{}, errors.New("resource not found")

}
