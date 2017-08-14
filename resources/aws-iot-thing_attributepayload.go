package resources

// AWS::IoT::Thing.AttributePayload AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-thing-attributepayload.html
type AWSIoTThingAttributePayload struct {

	// Attributes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-thing-attributepayload.html#cfn-iot-thing-attributepayload-attributes
	Attributes map[string]AWSIoTThingAttributePayloadstring `json:"Attributes"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTThingAttributePayload) AWSCloudFormationType() string {
	return "AWS::IoT::Thing.AttributePayload"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIoTThingAttributePayload) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
