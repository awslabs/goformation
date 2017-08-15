package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Route53::HealthCheck.HealthCheckTag AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthchecktag.html
type AWSRoute53HealthCheck_HealthCheckTag struct {

	// Key AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthchecktag.html#cfn-route53-healthchecktags-key
	Key string `json:"Key"`

	// Value AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-healthchecktag.html#cfn-route53-healthchecktags-value
	Value string `json:"Value"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRoute53HealthCheck_HealthCheckTag) AWSCloudFormationType() string {
	return "AWS::Route53::HealthCheck.HealthCheckTag"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSRoute53HealthCheck_HealthCheckTag) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSRoute53HealthCheck_HealthCheckTagResources retrieves all AWSRoute53HealthCheck_HealthCheckTag items from a CloudFormation template
func GetAllAWSRoute53HealthCheck_HealthCheckTag(template *Template) map[string]*AWSRoute53HealthCheck_HealthCheckTag {

	results := map[string]*AWSRoute53HealthCheck_HealthCheckTag{}
	for name, resource := range template.Resources {
		result := &AWSRoute53HealthCheck_HealthCheckTag{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSRoute53HealthCheck_HealthCheckTagWithName retrieves all AWSRoute53HealthCheck_HealthCheckTag items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSRoute53HealthCheck_HealthCheckTag(name string, template *Template) (*AWSRoute53HealthCheck_HealthCheckTag, error) {

	result := &AWSRoute53HealthCheck_HealthCheckTag{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSRoute53HealthCheck_HealthCheckTag{}, errors.New("resource not found")

}
