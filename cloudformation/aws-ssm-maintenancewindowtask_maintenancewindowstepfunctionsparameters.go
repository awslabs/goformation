package cloudformation

// AWSSSMMaintenanceWindowTask_MaintenanceWindowStepFunctionsParameters AWS CloudFormation Resource (AWS::SSM::MaintenanceWindowTask.MaintenanceWindowStepFunctionsParameters)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-taskinvocationparameters-maintenancewindowstepfunctionsparameters.html
type AWSSSMMaintenanceWindowTask_MaintenanceWindowStepFunctionsParameters struct {

	// Input AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-taskinvocationparameters-maintenancewindowstepfunctionsparameters.html#cfn-ssm-maintenancewindowtask-taskinvocationparameters-maintenancewindowstepfunctionsparameters-input
	Input string `json:"Input,omitempty"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ssm-maintenancewindowtask-taskinvocationparameters-maintenancewindowstepfunctionsparameters.html#cfn-ssm-maintenancewindowtask-taskinvocationparameters-maintenancewindowstepfunctionsparameters-name
	Name string `json:"Name,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSSMMaintenanceWindowTask_MaintenanceWindowStepFunctionsParameters) AWSCloudFormationType() string {
	return "AWS::SSM::MaintenanceWindowTask.MaintenanceWindowStepFunctionsParameters"
}
