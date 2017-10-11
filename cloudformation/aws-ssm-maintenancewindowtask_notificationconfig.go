package cloudformation

// AWSSSMMaintenanceWindowTask_NotificationConfig AWS CloudFormation Resource (AWS::SSM::MaintenanceWindowTask.NotificationConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-taskinvocationparameters-maintenancewindowruncommandparameters-notificationconfig.html
type AWSSSMMaintenanceWindowTask_NotificationConfig struct {

	// NotificationArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-taskinvocationparameters-maintenancewindowruncommandparameters-notificationconfig.html#cfn-ssm-maintenancewindowtask-taskinvocationparameters-maintenancewindowruncommandparameters-notificationconfig-notificationarn
	NotificationArn string `json:"NotificationArn,omitempty"`

	// NotificationEvents AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-taskinvocationparameters-maintenancewindowruncommandparameters-notificationconfig.html#cfn-ssm-maintenancewindowtask-taskinvocationparameters-maintenancewindowruncommandparameters-notificationconfig-notificationevents
	NotificationEvents []string `json:"NotificationEvents,omitempty"`

	// NotificationType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-taskinvocationparameters-maintenancewindowruncommandparameters-notificationconfig.html#cfn-ssm-maintenancewindowtask-taskinvocationparameters-maintenancewindowruncommandparameters-notificationconfig-notificationtype
	NotificationType string `json:"NotificationType,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSSMMaintenanceWindowTask_NotificationConfig) AWSCloudFormationType() string {
	return "AWS::SSM::MaintenanceWindowTask.NotificationConfig"
}
