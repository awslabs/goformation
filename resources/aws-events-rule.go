package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Events::Rule AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html
type AWSEventsRule struct {

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html#cfn-events-rule-description

	Description string `json:"Description"`

	// EventPattern AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html#cfn-events-rule-eventpattern

	EventPattern interface{} `json:"EventPattern"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html#cfn-events-rule-name

	Name string `json:"Name"`

	// RoleArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html#cfn-events-rule-rolearn

	RoleArn string `json:"RoleArn"`

	// ScheduleExpression AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html#cfn-events-rule-scheduleexpression

	ScheduleExpression string `json:"ScheduleExpression"`

	// State AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html#cfn-events-rule-state

	State string `json:"State"`

	// Targets AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-events-rule.html#cfn-events-rule-targets

	Targets []AWSEventsRule_Target `json:"Targets"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEventsRule) AWSCloudFormationType() string {
	return "AWS::Events::Rule"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSEventsRule) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSEventsRuleResources retrieves all AWSEventsRule items from a CloudFormation template
func GetAllAWSEventsRule(template *Template) map[string]*AWSEventsRule {

	results := map[string]*AWSEventsRule{}
	for name, resource := range template.Resources {
		result := &AWSEventsRule{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSEventsRuleWithName retrieves all AWSEventsRule items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSEventsRule(name string, template *Template) (*AWSEventsRule, error) {

	result := &AWSEventsRule{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSEventsRule{}, errors.New("resource not found")

}
