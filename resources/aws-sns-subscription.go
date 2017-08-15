package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::SNS::Subscription AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html
type AWSSNSSubscription struct {

	// Endpoint AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html#cfn-sns-endpoint
	Endpoint string `json:"Endpoint"`

	// Protocol AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html#cfn-sns-protocol
	Protocol string `json:"Protocol"`

	// TopicArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html#topicarn
	TopicArn string `json:"TopicArn"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSNSSubscription) AWSCloudFormationType() string {
	return "AWS::SNS::Subscription"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSSNSSubscription) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSSNSSubscriptionResources retrieves all AWSSNSSubscription items from a CloudFormation template
func GetAllAWSSNSSubscription(template *Template) map[string]*AWSSNSSubscription {

	results := map[string]*AWSSNSSubscription{}
	for name, resource := range template.Resources {
		result := &AWSSNSSubscription{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSSNSSubscriptionWithName retrieves all AWSSNSSubscription items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSSNSSubscription(name string, template *Template) (*AWSSNSSubscription, error) {

	result := &AWSSNSSubscription{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSSNSSubscription{}, errors.New("resource not found")

}
