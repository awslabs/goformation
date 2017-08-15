package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::IoT::TopicRule.DynamoDBAction AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-dynamodb.html
type AWSIoTTopicRule_DynamoDBAction struct {

	// HashKeyField AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-dynamodb.html#cfn-iot-dynamodb-hashkeyfield
	HashKeyField string `json:"HashKeyField"`

	// HashKeyValue AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-dynamodb.html#cfn-iot-dynamodb-hashkeyvalue
	HashKeyValue string `json:"HashKeyValue"`

	// PayloadField AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-dynamodb.html#cfn-iot-dynamodb-payloadfield
	PayloadField string `json:"PayloadField"`

	// RangeKeyField AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-dynamodb.html#cfn-iot-dynamodb-rangekeyfield
	RangeKeyField string `json:"RangeKeyField"`

	// RangeKeyValue AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-dynamodb.html#cfn-iot-dynamodb-rangekeyvalue
	RangeKeyValue string `json:"RangeKeyValue"`

	// RoleArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-dynamodb.html#cfn-iot-dynamodb-rolearn
	RoleArn string `json:"RoleArn"`

	// TableName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-dynamodb.html#cfn-iot-dynamodb-tablename
	TableName string `json:"TableName"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTTopicRule_DynamoDBAction) AWSCloudFormationType() string {
	return "AWS::IoT::TopicRule.DynamoDBAction"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIoTTopicRule_DynamoDBAction) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSIoTTopicRule_DynamoDBActionResources retrieves all AWSIoTTopicRule_DynamoDBAction items from a CloudFormation template
func GetAllAWSIoTTopicRule_DynamoDBAction(template *Template) map[string]*AWSIoTTopicRule_DynamoDBAction {

	results := map[string]*AWSIoTTopicRule_DynamoDBAction{}
	for name, resource := range template.Resources {
		result := &AWSIoTTopicRule_DynamoDBAction{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSIoTTopicRule_DynamoDBActionWithName retrieves all AWSIoTTopicRule_DynamoDBAction items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSIoTTopicRule_DynamoDBAction(name string, template *Template) (*AWSIoTTopicRule_DynamoDBAction, error) {

	result := &AWSIoTTopicRule_DynamoDBAction{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSIoTTopicRule_DynamoDBAction{}, errors.New("resource not found")

}
