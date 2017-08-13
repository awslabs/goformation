package resources

// AWS::Config::DeliveryChannel.ConfigSnapshotDeliveryProperties AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-deliverychannel-configsnapshotdeliveryproperties.html
type AWSConfigDeliveryChannelConfigSnapshotDeliveryProperties struct {
    
    // DeliveryFrequency AWS CloudFormation Property
    // Required: false
    // See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-deliverychannel-configsnapshotdeliveryproperties.html#cfn-config-deliverychannel-configsnapshotdeliveryproperties-deliveryfrequency
    DeliveryFrequency string `json:"DeliveryFrequency"`
    
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSConfigDeliveryChannelConfigSnapshotDeliveryProperties) AWSCloudFormationType() string {
    return "AWS::Config::DeliveryChannel.ConfigSnapshotDeliveryProperties"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSConfigDeliveryChannelConfigSnapshotDeliveryProperties) AWSCloudFormationSpecificationVersion() string {
    return "1.4.2"
}