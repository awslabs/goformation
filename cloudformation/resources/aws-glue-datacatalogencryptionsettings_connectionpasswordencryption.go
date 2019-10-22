package resources

import "github.com/awslabs/goformation/cloudformation/policies"

// AWSGlueDataCatalogEncryptionSettings_ConnectionPasswordEncryption AWS CloudFormation Resource (AWS::Glue::DataCatalogEncryptionSettings.ConnectionPasswordEncryption)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-datacatalogencryptionsettings-connectionpasswordencryption.html
type AWSGlueDataCatalogEncryptionSettings_ConnectionPasswordEncryption struct {

	// KmsKeyId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-datacatalogencryptionsettings-connectionpasswordencryption.html#cfn-glue-datacatalogencryptionsettings-connectionpasswordencryption-kmskeyid
	KmsKeyId string `json:"KmsKeyId,omitempty"`

	// ReturnConnectionPasswordEncrypted AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-datacatalogencryptionsettings-connectionpasswordencryption.html#cfn-glue-datacatalogencryptionsettings-connectionpasswordencryption-returnconnectionpasswordencrypted
	ReturnConnectionPasswordEncrypted bool `json:"ReturnConnectionPasswordEncrypted,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy policies.DeletionPolicy

	// _dependsOn stores the logical ID of the resources to be created before this resource
	_dependsOn []string

	// _metadata stores structured data associated with this resource
	_metadata map[string]interface{}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSGlueDataCatalogEncryptionSettings_ConnectionPasswordEncryption) AWSCloudFormationType() string {
	return "AWS::Glue::DataCatalogEncryptionSettings.ConnectionPasswordEncryption"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSGlueDataCatalogEncryptionSettings_ConnectionPasswordEncryption) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSGlueDataCatalogEncryptionSettings_ConnectionPasswordEncryption) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// Metadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSGlueDataCatalogEncryptionSettings_ConnectionPasswordEncryption) Metadata() map[string]interface{} {
	return r._metadata
}

// SetMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSGlueDataCatalogEncryptionSettings_ConnectionPasswordEncryption) SetMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSGlueDataCatalogEncryptionSettings_ConnectionPasswordEncryption) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}
