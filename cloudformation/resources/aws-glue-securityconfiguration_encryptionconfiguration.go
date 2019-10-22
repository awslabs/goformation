package resources

import "github.com/awslabs/goformation/cloudformation/policies"

// AWSGlueSecurityConfiguration_EncryptionConfiguration AWS CloudFormation Resource (AWS::Glue::SecurityConfiguration.EncryptionConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-securityconfiguration-encryptionconfiguration.html
type AWSGlueSecurityConfiguration_EncryptionConfiguration struct {

	// CloudWatchEncryption AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-securityconfiguration-encryptionconfiguration.html#cfn-glue-securityconfiguration-encryptionconfiguration-cloudwatchencryption
	CloudWatchEncryption *AWSGlueSecurityConfiguration_CloudWatchEncryption `json:"CloudWatchEncryption,omitempty"`

	// JobBookmarksEncryption AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-securityconfiguration-encryptionconfiguration.html#cfn-glue-securityconfiguration-encryptionconfiguration-jobbookmarksencryption
	JobBookmarksEncryption *AWSGlueSecurityConfiguration_JobBookmarksEncryption `json:"JobBookmarksEncryption,omitempty"`

	// S3Encryptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-glue-securityconfiguration-encryptionconfiguration.html#cfn-glue-securityconfiguration-encryptionconfiguration-s3encryptions
	S3Encryptions *AWSGlueSecurityConfiguration_S3Encryptions `json:"S3Encryptions,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy policies.DeletionPolicy

	// _dependsOn stores the logical ID of the resources to be created before this resource
	_dependsOn []string

	// _metadata stores structured data associated with this resource
	_metadata map[string]interface{}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSGlueSecurityConfiguration_EncryptionConfiguration) AWSCloudFormationType() string {
	return "AWS::Glue::SecurityConfiguration.EncryptionConfiguration"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSGlueSecurityConfiguration_EncryptionConfiguration) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSGlueSecurityConfiguration_EncryptionConfiguration) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// Metadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSGlueSecurityConfiguration_EncryptionConfiguration) Metadata() map[string]interface{} {
	return r._metadata
}

// SetMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSGlueSecurityConfiguration_EncryptionConfiguration) SetMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSGlueSecurityConfiguration_EncryptionConfiguration) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}
