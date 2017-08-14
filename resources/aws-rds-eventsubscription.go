package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::RDS::EventSubscription AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-eventsubscription.html
type AWSRDSEventSubscription struct {

	// Enabled AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-eventsubscription.html#cfn-rds-eventsubscription-enabled

	Enabled bool `json:"Enabled"`

	// EventCategories AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-eventsubscription.html#cfn-rds-eventsubscription-eventcategories

	EventCategories []string `json:"EventCategories"`

	// SnsTopicArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-eventsubscription.html#cfn-rds-eventsubscription-snstopicarn

	SnsTopicArn string `json:"SnsTopicArn"`

	// SourceIds AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-eventsubscription.html#cfn-rds-eventsubscription-sourceids

	SourceIds []string `json:"SourceIds"`

	// SourceType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-rds-eventsubscription.html#cfn-rds-eventsubscription-sourcetype

	SourceType string `json:"SourceType"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSRDSEventSubscription) AWSCloudFormationType() string {
	return "AWS::RDS::EventSubscription"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSRDSEventSubscription) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSRDSEventSubscriptionResources retrieves all AWSRDSEventSubscription items from a CloudFormation template
func GetAllAWSRDSEventSubscription(template *Template) map[string]*AWSRDSEventSubscription {

	results := map[string]*AWSRDSEventSubscription{}
	for name, resource := range template.Resources {
		result := &AWSRDSEventSubscription{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSRDSEventSubscriptionWithName retrieves all AWSRDSEventSubscription items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSRDSEventSubscription(name string, template *Template) (*AWSRDSEventSubscription, error) {

	result := &AWSRDSEventSubscription{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSRDSEventSubscription{}, errors.New("resource not found")

}
