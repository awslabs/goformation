package cloudformation

// AWSSSMMaintenanceWindowTask_MaintenanceWindowRunCommandParameters AWS CloudFormation Resource (AWS::SSM::MaintenanceWindowTask.MaintenanceWindowRunCommandParameters)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-taskinvocationparameters-maintenancewindowruncommandparameters.html
type AWSSSMMaintenanceWindowTask_MaintenanceWindowRunCommandParameters struct {

	// Comment AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-taskinvocationparameters-maintenancewindowruncommandparameters.html#cfn-ssm-maintenancewindowtask-taskinvocationparameters-maintenancewindowruncommandparameters-comment
	Comment string `json:"Comment,omitempty"`

	// DocumentHash AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-taskinvocationparameters-maintenancewindowruncommandparameters.html#cfn-ssm-maintenancewindowtask-taskinvocationparameters-maintenancewindowruncommandparameters-documenthash
	DocumentHash string `json:"DocumentHash,omitempty"`

	// DocumentHashType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-taskinvocationparameters-maintenancewindowruncommandparameters.html#cfn-ssm-maintenancewindowtask-taskinvocationparameters-maintenancewindowruncommandparameters-documenthashtype
	DocumentHashType string `json:"DocumentHashType,omitempty"`

	// NotificationConfig AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-taskinvocationparameters-maintenancewindowruncommandparameters.html#cfn-ssm-maintenancewindowtask-taskinvocationparameters-maintenancewindowruncommandparameters-notificationconfig
	NotificationConfig *AWSSSMMaintenanceWindowTask_NotificationConfig `json:"NotificationConfig,omitempty"`

	// OutputS3BucketName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-taskinvocationparameters-maintenancewindowruncommandparameters.html#cfn-ssm-maintenancewindowtask-taskinvocationparameters-maintenancewindowruncommandparameters-outputs3bucketname
	OutputS3BucketName string `json:"OutputS3BucketName,omitempty"`

	// OutputS3KeyPrefix AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-taskinvocationparameters-maintenancewindowruncommandparameters.html#cfn-ssm-maintenancewindowtask-taskinvocationparameters-maintenancewindowruncommandparameters-outputs3keyprefix
	OutputS3KeyPrefix string `json:"OutputS3KeyPrefix,omitempty"`

	// Parameters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-taskinvocationparameters-maintenancewindowruncommandparameters.html#cfn-ssm-maintenancewindowtask-taskinvocationparameters-maintenancewindowruncommandparameters-parameters
	Parameters interface{} `json:"Parameters,omitempty"`

	// ServiceRoleArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-taskinvocationparameters-maintenancewindowruncommandparameters.html#cfn-ssm-maintenancewindowtask-taskinvocationparameters-maintenancewindowruncommandparameters-servicerolearn
	ServiceRoleArn string `json:"ServiceRoleArn,omitempty"`

	// TimeoutSeconds AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-taskinvocationparameters-maintenancewindowruncommandparameters.html#cfn-ssm-maintenancewindowtask-taskinvocationparameters-maintenancewindowruncommandparameters-timeoutseconds
	TimeoutSeconds int `json:"TimeoutSeconds,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSSMMaintenanceWindowTask_MaintenanceWindowRunCommandParameters) AWSCloudFormationType() string {
	return "AWS::SSM::MaintenanceWindowTask.MaintenanceWindowRunCommandParameters"
}
