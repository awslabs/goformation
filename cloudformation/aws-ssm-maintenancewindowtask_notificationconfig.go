package cloudformation

import (
	"encoding/json"
)

// AWSSSMMaintenanceWindowTask_NotificationConfig AWS CloudFormation Resource (AWS::SSM::MaintenanceWindowTask.NotificationConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-notificationconfig.html
type AWSSSMMaintenanceWindowTask_NotificationConfig struct {

	// NotificationArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-notificationconfig.html#cfn-ssm-maintenancewindowtask-notificationconfig-notificationarn
	NotificationArn *Value `json:"NotificationArn,omitempty"`

	// NotificationEvents AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-notificationconfig.html#cfn-ssm-maintenancewindowtask-notificationconfig-notificationevents
	NotificationEvents []*Value `json:"NotificationEvents,omitempty"`

	// NotificationType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-notificationconfig.html#cfn-ssm-maintenancewindowtask-notificationconfig-notificationtype
	NotificationType *Value `json:"NotificationType,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSSMMaintenanceWindowTask_NotificationConfig) AWSCloudFormationType() string {
	return "AWS::SSM::MaintenanceWindowTask.NotificationConfig"
}

func (r *AWSSSMMaintenanceWindowTask_NotificationConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(*r)
}
