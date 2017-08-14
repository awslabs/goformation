package resources

// AWS::SNS::Subscription AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html
type AWSSNSSubscription struct {

	// Endpoint AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html#cfn-sns-endpoint

	Endpoint string `json:"Endpoint"`

	// Protocol AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html#cfn-sns-protocol

	Protocol string `json:"Protocol"`

	// TopicArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sns-subscription.html#topicarn

	TopicArn string `json:"TopicArn"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSNSSubscription) AWSCloudFormationType() string {
	return "AWS::SNS::Subscription"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSSNSSubscription) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
