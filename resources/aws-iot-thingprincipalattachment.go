package resources

// AWS::IoT::ThingPrincipalAttachment AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thingprincipalattachment.html
type AWSIoTThingPrincipalAttachment struct {

	// Principal AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thingprincipalattachment.html#cfn-iot-thingprincipalattachment-principal
	Principal string `json:"Principal"`

	// ThingName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thingprincipalattachment.html#cfn-iot-thingprincipalattachment-thingname
	ThingName string `json:"ThingName"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTThingPrincipalAttachment) AWSCloudFormationType() string {
	return "AWS::IoT::ThingPrincipalAttachment"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIoTThingPrincipalAttachment) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
