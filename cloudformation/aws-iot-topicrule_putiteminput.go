package cloudformation

// AWSIoTTopicRule_PutItemInput AWS CloudFormation Resource (AWS::IoT::TopicRule.PutItemInput)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-iot-dynamodbv2-putitem.html
type AWSIoTTopicRule_PutItemInput struct {

	// TableName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/cfn-iot-dynamodbv2-putitem.html#cfn-iot-dynamodbv2-putitem-tablename
	TableName string `json:"TableName,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTTopicRule_PutItemInput) AWSCloudFormationType() string {
	return "AWS::IoT::TopicRule.PutItemInput"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIoTTopicRule_PutItemInput) AWSCloudFormationSpecificationVersion() string {
	return "1.5.0"
}
