package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::SNS::Topic.Subscription AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-subscription.html
type AWSSNSTopic_Subscription struct {

	// Endpoint AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-subscription.html#cfn-sns-topic-subscription-endpoint

	Endpoint string `json:"Endpoint"`

	// Protocol AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-subscription.html#cfn-sns-topic-subscription-protocol

	Protocol string `json:"Protocol"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSNSTopic_Subscription) AWSCloudFormationType() string {
	return "AWS::SNS::Topic.Subscription"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSSNSTopic_Subscription) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSSNSTopic_SubscriptionResources retrieves all AWSSNSTopic_Subscription items from a CloudFormation template
func GetAllAWSSNSTopic_Subscription(template *Template) map[string]*AWSSNSTopic_Subscription {

	results := map[string]*AWSSNSTopic_Subscription{}
	for name, resource := range template.Resources {
		result := &AWSSNSTopic_Subscription{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSSNSTopic_SubscriptionWithName retrieves all AWSSNSTopic_Subscription items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSSNSTopic_Subscription(name string, template *Template) (*AWSSNSTopic_Subscription, error) {

	result := &AWSSNSTopic_Subscription{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSSNSTopic_Subscription{}, errors.New("resource not found")

}
