package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Route53::HostedZone.VPC AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone-hostedzonevpcs.html
type AWSRoute53HostedZone_VPC struct {

	// VPCId AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone-hostedzonevpcs.html#cfn-route53-hostedzone-hostedzonevpcs-vpcid
	VPCId string `json:"VPCId"`

	// VPCRegion AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-hostedzone-hostedzonevpcs.html#cfn-route53-hostedzone-hostedzonevpcs-vpcregion
	VPCRegion string `json:"VPCRegion"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRoute53HostedZone_VPC) AWSCloudFormationType() string {
	return "AWS::Route53::HostedZone.VPC"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSRoute53HostedZone_VPC) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSRoute53HostedZone_VPCResources retrieves all AWSRoute53HostedZone_VPC items from a CloudFormation template
func GetAllAWSRoute53HostedZone_VPC(template *Template) map[string]*AWSRoute53HostedZone_VPC {

	results := map[string]*AWSRoute53HostedZone_VPC{}
	for name, resource := range template.Resources {
		result := &AWSRoute53HostedZone_VPC{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSRoute53HostedZone_VPCWithName retrieves all AWSRoute53HostedZone_VPC items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSRoute53HostedZone_VPC(name string, template *Template) (*AWSRoute53HostedZone_VPC, error) {

	result := &AWSRoute53HostedZone_VPC{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSRoute53HostedZone_VPC{}, errors.New("resource not found")

}
