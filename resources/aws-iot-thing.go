package resources

// AWS::IoT::Thing AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thing.html
type AWSIoTThing struct {
    
    // AttributePayload AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thing.html#cfn-iot-thing-attributepayload
    AttributePayload AWSIoTThingAttributePayload `json:"AttributePayload"`
    
    // ThingName AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-thing.html#cfn-iot-thing-thingname
    ThingName string `json:"ThingName"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTThing) AWSCloudFormationType() string {
    return "AWS::IoT::Thing"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIoTThing) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}