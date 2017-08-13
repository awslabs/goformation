package resources

// AWS::IoT::PolicyPrincipalAttachment AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-policyprincipalattachment.html
type AWSIoTPolicyPrincipalAttachment struct {
    
    // PolicyName AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-policyprincipalattachment.html#cfn-iot-policyprincipalattachment-policyname
    PolicyName string `json:"PolicyName"`
    
    // Principal AWS CloudFormation Property
    // Required: true
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-iot-policyprincipalattachment.html#cfn-iot-policyprincipalattachment-principal
    Principal string `json:"Principal"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSIoTPolicyPrincipalAttachment) AWSCloudFormationType() string {
    return "AWS::IoT::PolicyPrincipalAttachment"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSIoTPolicyPrincipalAttachment) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}