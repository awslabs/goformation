package kendra

import (
	"github.com/awslabs/goformation/v5/cloudformation/policies"
)

// DataSource_GoogleDriveConfiguration AWS CloudFormation Resource (AWS::Kendra::DataSource.GoogleDriveConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-googledriveconfiguration.html
type DataSource_GoogleDriveConfiguration struct {

	// ExcludeMimeTypes AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-googledriveconfiguration.html#cfn-kendra-datasource-googledriveconfiguration-excludemimetypes
	ExcludeMimeTypes []string `json:"ExcludeMimeTypes,omitempty"`

	// ExcludeSharedDrives AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-googledriveconfiguration.html#cfn-kendra-datasource-googledriveconfiguration-excludeshareddrives
	ExcludeSharedDrives []string `json:"ExcludeSharedDrives,omitempty"`

	// ExcludeUserAccounts AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-googledriveconfiguration.html#cfn-kendra-datasource-googledriveconfiguration-excludeuseraccounts
	ExcludeUserAccounts []string `json:"ExcludeUserAccounts,omitempty"`

	// ExclusionPatterns AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-googledriveconfiguration.html#cfn-kendra-datasource-googledriveconfiguration-exclusionpatterns
	ExclusionPatterns []string `json:"ExclusionPatterns,omitempty"`

	// FieldMappings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-googledriveconfiguration.html#cfn-kendra-datasource-googledriveconfiguration-fieldmappings
	FieldMappings []DataSource_DataSourceToIndexFieldMapping `json:"FieldMappings,omitempty"`

	// InclusionPatterns AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-googledriveconfiguration.html#cfn-kendra-datasource-googledriveconfiguration-inclusionpatterns
	InclusionPatterns []string `json:"InclusionPatterns,omitempty"`

	// SecretArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-kendra-datasource-googledriveconfiguration.html#cfn-kendra-datasource-googledriveconfiguration-secretarn
	SecretArn string `json:"SecretArn,omitempty"`

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
func (r *DataSource_GoogleDriveConfiguration) AWSCloudFormationType() string {
	return "AWS::Kendra::DataSource.GoogleDriveConfiguration"
}
