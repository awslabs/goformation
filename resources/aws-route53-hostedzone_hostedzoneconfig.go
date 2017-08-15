package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Route53::HostedZone.HostedZoneConfig AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-hostedzoneconfig.html
type AWSRoute53HostedZone_HostedZoneConfig struct {

	// Comment AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-hostedzone-hostedzoneconfig.html#cfn-route53-hostedzone-hostedzoneconfig-comment

	Comment string `json:"Comment"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRoute53HostedZone_HostedZoneConfig) AWSCloudFormationType() string {
	return "AWS::Route53::HostedZone.HostedZoneConfig"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSRoute53HostedZone_HostedZoneConfig) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSRoute53HostedZone_HostedZoneConfigResources retrieves all AWSRoute53HostedZone_HostedZoneConfig items from a CloudFormation template
func GetAllAWSRoute53HostedZone_HostedZoneConfig(template *Template) map[string]*AWSRoute53HostedZone_HostedZoneConfig {

	results := map[string]*AWSRoute53HostedZone_HostedZoneConfig{}
	for name, resource := range template.Resources {
		result := &AWSRoute53HostedZone_HostedZoneConfig{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSRoute53HostedZone_HostedZoneConfigWithName retrieves all AWSRoute53HostedZone_HostedZoneConfig items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSRoute53HostedZone_HostedZoneConfig(name string, template *Template) (*AWSRoute53HostedZone_HostedZoneConfig, error) {

	result := &AWSRoute53HostedZone_HostedZoneConfig{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSRoute53HostedZone_HostedZoneConfig{}, errors.New("resource not found")

}
