// Code generated by "go generate". Please don't change this file directly.

package qbusiness

import (
	"github.com/awslabs/goformation/v7/cloudformation/policies"
)

// DataSource_DocumentAttributeValue AWS CloudFormation Resource (AWS::QBusiness::DataSource.DocumentAttributeValue)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-documentattributevalue.html
type DataSource_DocumentAttributeValue struct {

	// DateValue AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-documentattributevalue.html#cfn-qbusiness-datasource-documentattributevalue-datevalue
	DateValue *string `json:"DateValue,omitempty"`

	// LongValue AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-documentattributevalue.html#cfn-qbusiness-datasource-documentattributevalue-longvalue
	LongValue *float64 `json:"LongValue,omitempty"`

	// StringListValue AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-documentattributevalue.html#cfn-qbusiness-datasource-documentattributevalue-stringlistvalue
	StringListValue []string `json:"StringListValue,omitempty"`

	// StringValue AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-qbusiness-datasource-documentattributevalue.html#cfn-qbusiness-datasource-documentattributevalue-stringvalue
	StringValue *string `json:"StringValue,omitempty"`

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
func (r *DataSource_DocumentAttributeValue) AWSCloudFormationType() string {
	return "AWS::QBusiness::DataSource.DocumentAttributeValue"
}
