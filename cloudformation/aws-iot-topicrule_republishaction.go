package cloudformation

import (
	"encoding/json"
)

// AWSIoTTopicRule_RepublishAction AWS CloudFormation Resource (AWS::IoT::TopicRule.RepublishAction)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-republishaction.html
type AWSIoTTopicRule_RepublishAction struct {

	// RoleArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-republishaction.html#cfn-iot-topicrule-republishaction-rolearn
	RoleArn *Value `json:"RoleArn,omitempty"`

	// Topic AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-republishaction.html#cfn-iot-topicrule-republishaction-topic
	Topic *Value `json:"Topic,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTTopicRule_RepublishAction) AWSCloudFormationType() string {
	return "AWS::IoT::TopicRule.RepublishAction"
}

func (r *AWSIoTTopicRule_RepublishAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
