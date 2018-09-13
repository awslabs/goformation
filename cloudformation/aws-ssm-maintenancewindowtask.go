package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSSSMMaintenanceWindowTask AWS CloudFormation Resource (AWS::SSM::MaintenanceWindowTask)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html
type AWSSSMMaintenanceWindowTask struct {

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html#cfn-ssm-maintenancewindowtask-description
	Description *Value `json:"Description,omitempty"`

	// LoggingInfo AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html#cfn-ssm-maintenancewindowtask-logginginfo
	LoggingInfo *AWSSSMMaintenanceWindowTask_LoggingInfo `json:"LoggingInfo,omitempty"`

	// MaxConcurrency AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html#cfn-ssm-maintenancewindowtask-maxconcurrency
	MaxConcurrency *Value `json:"MaxConcurrency,omitempty"`

	// MaxErrors AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html#cfn-ssm-maintenancewindowtask-maxerrors
	MaxErrors *Value `json:"MaxErrors,omitempty"`

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html#cfn-ssm-maintenancewindowtask-name
	Name *Value `json:"Name,omitempty"`

	// Priority AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html#cfn-ssm-maintenancewindowtask-priority
	Priority *Value `json:"Priority,omitempty"`

	// ServiceRoleArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html#cfn-ssm-maintenancewindowtask-servicerolearn
	ServiceRoleArn *Value `json:"ServiceRoleArn,omitempty"`

	// Targets AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html#cfn-ssm-maintenancewindowtask-targets
	Targets []AWSSSMMaintenanceWindowTask_Target `json:"Targets,omitempty"`

	// TaskArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html#cfn-ssm-maintenancewindowtask-taskarn
	TaskArn *Value `json:"TaskArn,omitempty"`

	// TaskInvocationParameters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html#cfn-ssm-maintenancewindowtask-taskinvocationparameters
	TaskInvocationParameters *AWSSSMMaintenanceWindowTask_TaskInvocationParameters `json:"TaskInvocationParameters,omitempty"`

	// TaskParameters AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html#cfn-ssm-maintenancewindowtask-taskparameters
	TaskParameters interface{} `json:"TaskParameters,omitempty"`

	// TaskType AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html#cfn-ssm-maintenancewindowtask-tasktype
	TaskType *Value `json:"TaskType,omitempty"`

	// WindowId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-maintenancewindowtask.html#cfn-ssm-maintenancewindowtask-windowid
	WindowId *Value `json:"WindowId,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSSMMaintenanceWindowTask) AWSCloudFormationType() string {
	return "AWS::SSM::MaintenanceWindowTask"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSSSMMaintenanceWindowTask) MarshalJSON() ([]byte, error) {
	type Properties AWSSSMMaintenanceWindowTask
	return json.Marshal(&struct {
		Type       string
		Properties Properties
	}{
		Type:       r.AWSCloudFormationType(),
		Properties: (Properties)(*r),
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *AWSSSMMaintenanceWindowTask) UnmarshalJSON(b []byte) error {
	type Properties AWSSSMMaintenanceWindowTask
	res := &struct {
		Type       string
		Properties *Properties
	}{}
	if err := json.Unmarshal(b, &res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}

	// If the resource has no Properties set, it could be nil
	if res.Properties != nil {
		*r = AWSSSMMaintenanceWindowTask(*res.Properties)
	}

	return nil
}

// GetAllAWSSSMMaintenanceWindowTaskResources retrieves all AWSSSMMaintenanceWindowTask items from an AWS CloudFormation template
func (t *Template) GetAllAWSSSMMaintenanceWindowTaskResources() map[string]AWSSSMMaintenanceWindowTask {
	results := map[string]AWSSSMMaintenanceWindowTask{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSSSMMaintenanceWindowTask:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::SSM::MaintenanceWindowTask" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSSSMMaintenanceWindowTask{}
						if err := result.UnmarshalJSON(b); err == nil {
							results[name] = *result
						}
					}
				}
			}
		}
	}
	return results
}

// GetAWSSSMMaintenanceWindowTaskWithName retrieves all AWSSSMMaintenanceWindowTask items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSSSMMaintenanceWindowTaskWithName(name string) (AWSSSMMaintenanceWindowTask, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSSSMMaintenanceWindowTask:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::SSM::MaintenanceWindowTask" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						result := &AWSSSMMaintenanceWindowTask{}
						if err := result.UnmarshalJSON(b); err == nil {
							return *result, nil
						}
					}
				}
			}
		}
	}
	return AWSSSMMaintenanceWindowTask{}, errors.New("resource not found")
}
