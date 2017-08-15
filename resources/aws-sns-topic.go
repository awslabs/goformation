package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::SNS::Topic AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html
type AWSSNSTopic struct {

	// DisplayName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html#cfn-sns-topic-displayname
	DisplayName string `json:"DisplayName"`

	// Subscription AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html#cfn-sns-topic-subscription
	Subscription []AWSSNSTopic_Subscription `json:"Subscription"`

	// TopicName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-topic.html#cfn-sns-topic-topicname
	TopicName string `json:"TopicName"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSNSTopic) AWSCloudFormationType() string {
	return "AWS::SNS::Topic"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSSNSTopic) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSSNSTopicResources retrieves all AWSSNSTopic items from a CloudFormation template
func GetAllAWSSNSTopicResources(template *Template) map[string]*AWSSNSTopic {

	results := map[string]*AWSSNSTopic{}
	for name, resource := range template.Resources {
		result := &AWSSNSTopic{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSSNSTopicWithName retrieves all AWSSNSTopic items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSSNSTopicWithName(name string, template *Template) (*AWSSNSTopic, error) {

	result := &AWSSNSTopic{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSSNSTopic{}, errors.New("resource not found")

}
