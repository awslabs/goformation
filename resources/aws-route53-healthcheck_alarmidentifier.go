package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Route53::HealthCheck.AlarmIdentifier AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-alarmidentifier.html
type AWSRoute53HealthCheck_AlarmIdentifier struct {

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-alarmidentifier.html#cfn-route53-healthcheck-alarmidentifier-name

	Name string `json:"Name"`

	// Region AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-route53-healthcheck-alarmidentifier.html#cfn-route53-healthcheck-alarmidentifier-region

	Region string `json:"Region"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRoute53HealthCheck_AlarmIdentifier) AWSCloudFormationType() string {
	return "AWS::Route53::HealthCheck.AlarmIdentifier"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSRoute53HealthCheck_AlarmIdentifier) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSRoute53HealthCheck_AlarmIdentifierResources retrieves all AWSRoute53HealthCheck_AlarmIdentifier items from a CloudFormation template
func GetAllAWSRoute53HealthCheck_AlarmIdentifier(template *Template) map[string]*AWSRoute53HealthCheck_AlarmIdentifier {

	results := map[string]*AWSRoute53HealthCheck_AlarmIdentifier{}
	for name, resource := range template.Resources {
		result := &AWSRoute53HealthCheck_AlarmIdentifier{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSRoute53HealthCheck_AlarmIdentifierWithName retrieves all AWSRoute53HealthCheck_AlarmIdentifier items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSRoute53HealthCheck_AlarmIdentifier(name string, template *Template) (*AWSRoute53HealthCheck_AlarmIdentifier, error) {

	result := &AWSRoute53HealthCheck_AlarmIdentifier{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSRoute53HealthCheck_AlarmIdentifier{}, errors.New("resource not found")

}
