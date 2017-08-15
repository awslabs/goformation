package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSSNSTopicPolicy AWS CloudFormation Resource (AWS::SNS::TopicPolicy)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-policy.html
type AWSSNSTopicPolicy struct {

	// PolicyDocument AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-policy.html#cfn-sns-topicpolicy-policydocument
	PolicyDocument interface{} `json:"PolicyDocument"`

	// Topics AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-policy.html#cfn-sns-topicpolicy-topics
	Topics []string `json:"Topics"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSNSTopicPolicy) AWSCloudFormationType() string {
	return "AWS::SNS::TopicPolicy"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSSNSTopicPolicy) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSSNSTopicPolicyResources retrieves all AWSSNSTopicPolicy items from a CloudFormation template
func GetAllAWSSNSTopicPolicyResources(template *Template) map[string]*AWSSNSTopicPolicy {

	results := map[string]*AWSSNSTopicPolicy{}
	for name, resource := range template.Resources {
		result := &AWSSNSTopicPolicy{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSSNSTopicPolicyWithName retrieves all AWSSNSTopicPolicy items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSSNSTopicPolicyWithName(name string, template *Template) (*AWSSNSTopicPolicy, error) {

	result := &AWSSNSTopicPolicy{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSSNSTopicPolicy{}, errors.New("resource not found")

}
