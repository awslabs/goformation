package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Route53::HostedZone.HostedZoneTag AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-hostedzonetags.html
type AWSRoute53HostedZone_HostedZoneTag struct {

	// Key AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-hostedzonetags.html#cfn-route53-hostedzonetags-key
	Key string `json:"Key"`

	// Value AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-hostedzonetags.html#cfn-route53-hostedzonetags-value
	Value string `json:"Value"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRoute53HostedZone_HostedZoneTag) AWSCloudFormationType() string {
	return "AWS::Route53::HostedZone.HostedZoneTag"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSRoute53HostedZone_HostedZoneTag) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSRoute53HostedZone_HostedZoneTagResources retrieves all AWSRoute53HostedZone_HostedZoneTag items from a CloudFormation template
func GetAllAWSRoute53HostedZone_HostedZoneTag(template *Template) map[string]*AWSRoute53HostedZone_HostedZoneTag {

	results := map[string]*AWSRoute53HostedZone_HostedZoneTag{}
	for name, resource := range template.Resources {
		result := &AWSRoute53HostedZone_HostedZoneTag{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSRoute53HostedZone_HostedZoneTagWithName retrieves all AWSRoute53HostedZone_HostedZoneTag items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSRoute53HostedZone_HostedZoneTag(name string, template *Template) (*AWSRoute53HostedZone_HostedZoneTag, error) {

	result := &AWSRoute53HostedZone_HostedZoneTag{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSRoute53HostedZone_HostedZoneTag{}, errors.New("resource not found")

}
