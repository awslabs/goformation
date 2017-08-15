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
