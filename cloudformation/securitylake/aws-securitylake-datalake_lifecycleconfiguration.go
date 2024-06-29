// Code generated by "go generate". Please don't change this file directly.

package securitylake

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// DataLake_LifecycleConfiguration AWS CloudFormation Resource (AWS::SecurityLake::DataLake.LifecycleConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-datalake-lifecycleconfiguration.html
type DataLake_LifecycleConfiguration struct {

	// Expiration AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-datalake-lifecycleconfiguration.html#cfn-securitylake-datalake-lifecycleconfiguration-expiration
	Expiration *DataLake_Expiration `json:"Expiration,omitempty"`

	// Transitions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-securitylake-datalake-lifecycleconfiguration.html#cfn-securitylake-datalake-lifecycleconfiguration-transitions
	Transitions []DataLake_Transitions `json:"Transitions,omitempty"`

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
func (r *DataLake_LifecycleConfiguration) AWSCloudFormationType() string {
	return "AWS::SecurityLake::DataLake.LifecycleConfiguration"
}