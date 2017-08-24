package cloudformation

// AWSIoTTopicRule_DynamoDBv2Action AWS CloudFormation Resource (AWS::IoT::TopicRule.DynamoDBv2Action)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-dynamodbv2.html
type AWSIoTTopicRule_DynamoDBv2Action struct {

	// PutItem AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-dynamodbv2.html#cfn-iot-dynamodbv2-putitem
	PutItem *AWSIoTTopicRule_PutItemInput `json:"PutItem,omitempty"`

	// RoleArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-dynamodbv2.html#cfn-iot-dynamodbv2-rolearn
	RoleArn string `json:"RoleArn,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTTopicRule_DynamoDBv2Action) AWSCloudFormationType() string {
	return "AWS::IoT::TopicRule.DynamoDBv2Action"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIoTTopicRule_DynamoDBv2Action) AWSCloudFormationSpecificationVersion() string {
	return "1.5.0"
}
