package cloudformation

import (
	"encoding/json"
)

// AWSSESConfigurationSetEventDestination_KinesisFirehoseDestination AWS CloudFormation Resource (AWS::SES::ConfigurationSetEventDestination.KinesisFirehoseDestination)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationseteventdestination-kinesisfirehosedestination.html
type AWSSESConfigurationSetEventDestination_KinesisFirehoseDestination struct {

	// DeliveryStreamARN AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationseteventdestination-kinesisfirehosedestination.html#cfn-ses-configurationseteventdestination-kinesisfirehosedestination-deliverystreamarn
	DeliveryStreamARN *Value `json:"DeliveryStreamARN,omitempty"`

	// IAMRoleARN AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-configurationseteventdestination-kinesisfirehosedestination.html#cfn-ses-configurationseteventdestination-kinesisfirehosedestination-iamrolearn
	IAMRoleARN *Value `json:"IAMRoleARN,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSESConfigurationSetEventDestination_KinesisFirehoseDestination) AWSCloudFormationType() string {
	return "AWS::SES::ConfigurationSetEventDestination.KinesisFirehoseDestination"
}

func (r *AWSSESConfigurationSetEventDestination_KinesisFirehoseDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
