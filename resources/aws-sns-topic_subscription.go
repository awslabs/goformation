package resources

// AWS::SNS::Topic.Subscription AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sns-subscription.html
type AWSSNSTopicSubscription struct {

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
func (r *AWSSNSTopicSubscription) AWSCloudFormationType() string {
	return "AWS::SNS::Topic.Subscription"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSSNSTopicSubscription) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
