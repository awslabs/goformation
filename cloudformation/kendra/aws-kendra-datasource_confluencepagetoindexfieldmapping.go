package kendra

import (
	"github.com/awslabs/goformation/v5/cloudformation/policies"
)

// DataSource_ConfluencePageToIndexFieldMapping AWS CloudFormation Resource (AWS::Kendra::DataSource.ConfluencePageToIndexFieldMapping)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluencepagetoindexfieldmapping.html
type DataSource_ConfluencePageToIndexFieldMapping struct {

	// DataSourceFieldName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluencepagetoindexfieldmapping.html#cfn-kendra-datasource-confluencepagetoindexfieldmapping-datasourcefieldname
	DataSourceFieldName string `json:"DataSourceFieldName"`

	// DateFieldFormat AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluencepagetoindexfieldmapping.html#cfn-kendra-datasource-confluencepagetoindexfieldmapping-datefieldformat
	DateFieldFormat *string `json:"DateFieldFormat,omitempty"`

	// IndexFieldName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-confluencepagetoindexfieldmapping.html#cfn-kendra-datasource-confluencepagetoindexfieldmapping-indexfieldname
	IndexFieldName string `json:"IndexFieldName"`

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
func (r *DataSource_ConfluencePageToIndexFieldMapping) AWSCloudFormationType() string {
	return "AWS::Kendra::DataSource.ConfluencePageToIndexFieldMapping"
}
