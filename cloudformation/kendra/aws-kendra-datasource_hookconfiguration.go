package kendra

import (
	"github.com/awslabs/goformation/v5/cloudformation/policies"
)

// DataSource_HookConfiguration AWS CloudFormation Resource (AWS::Kendra::DataSource.HookConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-hookconfiguration.html
type DataSource_HookConfiguration struct {

	// InvocationCondition AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-hookconfiguration.html#cfn-kendra-datasource-hookconfiguration-invocationcondition
	InvocationCondition *DataSource_DocumentAttributeCondition `json:"InvocationCondition,omitempty"`

	// LambdaArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-hookconfiguration.html#cfn-kendra-datasource-hookconfiguration-lambdaarn
	LambdaArn string `json:"LambdaArn,omitempty"`

	// S3Bucket AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-hookconfiguration.html#cfn-kendra-datasource-hookconfiguration-s3bucket
	S3Bucket string `json:"S3Bucket,omitempty"`

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
func (r *DataSource_HookConfiguration) AWSCloudFormationType() string {
	return "AWS::Kendra::DataSource.HookConfiguration"
}
