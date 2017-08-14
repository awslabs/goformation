package resources

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
	Subscription []AWSSNSTopicSubscription `json:"Subscription"`

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
