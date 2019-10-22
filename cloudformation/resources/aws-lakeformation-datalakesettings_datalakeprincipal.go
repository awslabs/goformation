package resources

import "github.com/awslabs/goformation/cloudformation/policies"

// AWSLakeFormationDataLakeSettings_DataLakePrincipal AWS CloudFormation Resource (AWS::LakeFormation::DataLakeSettings.DataLakePrincipal)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-datalakesettings-datalakeprincipal.html
type AWSLakeFormationDataLakeSettings_DataLakePrincipal struct {

	// DataLakePrincipalIdentifier AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-lakeformation-datalakesettings-datalakeprincipal.html#cfn-lakeformation-datalakesettings-datalakeprincipal-datalakeprincipalidentifier
	DataLakePrincipalIdentifier string `json:"DataLakePrincipalIdentifier,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy policies.DeletionPolicy

	// _dependsOn stores the logical ID of the resources to be created before this resource
	_dependsOn []string

	// _metadata stores structured data associated with this resource
	_metadata map[string]interface{}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSLakeFormationDataLakeSettings_DataLakePrincipal) AWSCloudFormationType() string {
	return "AWS::LakeFormation::DataLakeSettings.DataLakePrincipal"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSLakeFormationDataLakeSettings_DataLakePrincipal) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSLakeFormationDataLakeSettings_DataLakePrincipal) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// Metadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSLakeFormationDataLakeSettings_DataLakePrincipal) Metadata() map[string]interface{} {
	return r._metadata
}

// SetMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSLakeFormationDataLakeSettings_DataLakePrincipal) SetMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSLakeFormationDataLakeSettings_DataLakePrincipal) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}
