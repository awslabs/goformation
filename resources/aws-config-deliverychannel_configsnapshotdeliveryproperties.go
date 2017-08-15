package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Config::DeliveryChannel.ConfigSnapshotDeliveryProperties AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-deliverychannel-configsnapshotdeliveryproperties.html
type AWSConfigDeliveryChannel_ConfigSnapshotDeliveryProperties struct {

	// DeliveryFrequency AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-deliverychannel-configsnapshotdeliveryproperties.html#cfn-config-deliverychannel-configsnapshotdeliveryproperties-deliveryfrequency
	DeliveryFrequency string `json:"DeliveryFrequency"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSConfigDeliveryChannel_ConfigSnapshotDeliveryProperties) AWSCloudFormationType() string {
	return "AWS::Config::DeliveryChannel.ConfigSnapshotDeliveryProperties"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSConfigDeliveryChannel_ConfigSnapshotDeliveryProperties) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
