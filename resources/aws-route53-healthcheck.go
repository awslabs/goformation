package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Route53::HealthCheck AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-healthcheck.html
type AWSRoute53HealthCheck struct {

	// HealthCheckConfig AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-healthcheck.html#cfn-route53-healthcheck-healthcheckconfig

	HealthCheckConfig AWSRoute53HealthCheck_HealthCheckConfig `json:"HealthCheckConfig"`

	// HealthCheckTags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-healthcheck.html#cfn-route53-healthcheck-healthchecktags

	HealthCheckTags []AWSRoute53HealthCheck_HealthCheckTag `json:"HealthCheckTags"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRoute53HealthCheck) AWSCloudFormationType() string {
	return "AWS::Route53::HealthCheck"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSRoute53HealthCheck) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSRoute53HealthCheckResources retrieves all AWSRoute53HealthCheck items from a CloudFormation template
func GetAllAWSRoute53HealthCheck(template *Template) map[string]*AWSRoute53HealthCheck {

	results := map[string]*AWSRoute53HealthCheck{}
	for name, resource := range template.Resources {
		result := &AWSRoute53HealthCheck{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSRoute53HealthCheckWithName retrieves all AWSRoute53HealthCheck items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSRoute53HealthCheck(name string, template *Template) (*AWSRoute53HealthCheck, error) {

	result := &AWSRoute53HealthCheck{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSRoute53HealthCheck{}, errors.New("resource not found")

}
