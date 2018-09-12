package cloudformation

import (
	"encoding/json"
)

// AWSConfigDeliveryChannel_ConfigSnapshotDeliveryProperties AWS CloudFormation Resource (AWS::Config::DeliveryChannel.ConfigSnapshotDeliveryProperties)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-deliverychannel-configsnapshotdeliveryproperties.html
type AWSConfigDeliveryChannel_ConfigSnapshotDeliveryProperties struct {

	// DeliveryFrequency AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-deliverychannel-configsnapshotdeliveryproperties.html#cfn-config-deliverychannel-configsnapshotdeliveryproperties-deliveryfrequency
	DeliveryFrequency *Value `json:"DeliveryFrequency,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSConfigDeliveryChannel_ConfigSnapshotDeliveryProperties) AWSCloudFormationType() string {
	return "AWS::Config::DeliveryChannel.ConfigSnapshotDeliveryProperties"
}

func (r *AWSConfigDeliveryChannel_ConfigSnapshotDeliveryProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
