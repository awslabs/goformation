// Code generated by "go generate". Please don't change this file directly.

package internetmonitor

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// Monitor_InternetMeasurementsLogDelivery AWS CloudFormation Resource (AWS::InternetMonitor::Monitor.InternetMeasurementsLogDelivery)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-internetmonitor-monitor-internetmeasurementslogdelivery.html
type Monitor_InternetMeasurementsLogDelivery struct {

	// S3Config AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-internetmonitor-monitor-internetmeasurementslogdelivery.html#cfn-internetmonitor-monitor-internetmeasurementslogdelivery-s3config
	S3Config *Monitor_S3Config `json:"S3Config,omitempty"`

	// AWSCloudFormationDeletionPolicy represents a CloudFormation DeletionPolicy
	AWSCloudFormationDeletionPolicy policies.DeletionPolicy `json:"-"`

	// AWSCloudFormationUpdateReplacePolicy represents a CloudFormation UpdateReplacePolicy
	AWSCloudFormationUpdateReplacePolicy policies.UpdateReplacePolicy `json:"-"`

	// AWSCloudFormationDependsOn stores the logical ID of the resources to be created before this resource
	AWSCloudFormationDependsOn []string `json:"-"`

	// AWSCloudFormationMetadata stores structured data associated with this resource
	AWSCloudFormationMetadata map[string]interface{} `json:"-"`

	// AWSCloudFormationCondition stores the logical ID of the condition that must be satisfied for this resource to be created
	AWSCloudFormationCondition string `json:"-"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *Monitor_InternetMeasurementsLogDelivery) AWSCloudFormationType() string {
	return "AWS::InternetMonitor::Monitor.InternetMeasurementsLogDelivery"
}
