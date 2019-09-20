package resources

import "github.com/TachyonNexus/goformation/cloudformation/policies"

// AWSGlueDataCatalogEncryptionSettings_DataCatalogEncryptionSettings AWS CloudFormation Resource (AWS::Glue::DataCatalogEncryptionSettings.DataCatalogEncryptionSettings)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-datacatalogencryptionsettings-datacatalogencryptionsettings.html
type AWSGlueDataCatalogEncryptionSettings_DataCatalogEncryptionSettings struct {

	// ConnectionPasswordEncryption AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-datacatalogencryptionsettings-datacatalogencryptionsettings.html#cfn-glue-datacatalogencryptionsettings-datacatalogencryptionsettings-connectionpasswordencryption
	ConnectionPasswordEncryption *AWSGlueDataCatalogEncryptionSettings_ConnectionPasswordEncryption `json:"ConnectionPasswordEncryption,omitempty"`

	// EncryptionAtRest AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-datacatalogencryptionsettings-datacatalogencryptionsettings.html#cfn-glue-datacatalogencryptionsettings-datacatalogencryptionsettings-encryptionatrest
	EncryptionAtRest *AWSGlueDataCatalogEncryptionSettings_EncryptionAtRest `json:"EncryptionAtRest,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy policies.DeletionPolicy

	// _dependsOn stores the logical ID of the resources to be created before this resource
	_dependsOn []string

	// _metadata stores structured data associated with this resource
	_metadata map[string]interface{}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSGlueDataCatalogEncryptionSettings_DataCatalogEncryptionSettings) AWSCloudFormationType() string {
	return "AWS::Glue::DataCatalogEncryptionSettings.DataCatalogEncryptionSettings"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSGlueDataCatalogEncryptionSettings_DataCatalogEncryptionSettings) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSGlueDataCatalogEncryptionSettings_DataCatalogEncryptionSettings) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// Metadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSGlueDataCatalogEncryptionSettings_DataCatalogEncryptionSettings) Metadata() map[string]interface{} {
	return r._metadata
}

// SetMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSGlueDataCatalogEncryptionSettings_DataCatalogEncryptionSettings) SetMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// DeletionPolicy returns the AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSGlueDataCatalogEncryptionSettings_DataCatalogEncryptionSettings) DeletionPolicy() policies.DeletionPolicy {
	return r._deletionPolicy
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSGlueDataCatalogEncryptionSettings_DataCatalogEncryptionSettings) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}
