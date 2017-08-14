package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::IoT::TopicRule.LambdaAction AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-lambda.html
type AWSIoTTopicRule_LambdaAction struct {

	// FunctionArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-lambda.html#cfn-iot-lambda-functionarn

	FunctionArn string `json:"FunctionArn"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTTopicRule_LambdaAction) AWSCloudFormationType() string {
	return "AWS::IoT::TopicRule.LambdaAction"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIoTTopicRule_LambdaAction) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSIoTTopicRule_LambdaActionResources retrieves all AWSIoTTopicRule_LambdaAction items from a CloudFormation template
func GetAllAWSIoTTopicRule_LambdaAction(template *Template) map[string]*AWSIoTTopicRule_LambdaAction {

	results := map[string]*AWSIoTTopicRule_LambdaAction{}
	for name, resource := range template.Resources {
		result := &AWSIoTTopicRule_LambdaAction{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSIoTTopicRule_LambdaActionWithName retrieves all AWSIoTTopicRule_LambdaAction items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSIoTTopicRule_LambdaAction(name string, template *Template) (*AWSIoTTopicRule_LambdaAction, error) {

	result := &AWSIoTTopicRule_LambdaAction{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSIoTTopicRule_LambdaAction{}, errors.New("resource not found")

}
