package resources

import "github.com/awslabs/goformation/v2/cloudformation/policies"

// AWSConfigRemediationConfiguration_ExecutionControls AWS CloudFormation Resource (AWS::Config::RemediationConfiguration.ExecutionControls)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-remediationconfiguration-executioncontrols.html
type AWSConfigRemediationConfiguration_ExecutionControls struct {

	// SsmControls AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-remediationconfiguration-executioncontrols.html#cfn-config-remediationconfiguration-executioncontrols-ssmcontrols
	SsmControls *AWSConfigRemediationConfiguration_SsmControls `json:"SsmControls,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy policies.DeletionPolicy

	// _dependsOn stores the logical ID of the resources to be created before this resource
	_dependsOn []string

	// _metadata stores structured data associated with this resource
	_metadata map[string]interface{}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSConfigRemediationConfiguration_ExecutionControls) AWSCloudFormationType() string {
	return "AWS::Config::RemediationConfiguration.ExecutionControls"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSConfigRemediationConfiguration_ExecutionControls) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSConfigRemediationConfiguration_ExecutionControls) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// Metadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSConfigRemediationConfiguration_ExecutionControls) Metadata() map[string]interface{} {
	return r._metadata
}

// SetMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSConfigRemediationConfiguration_ExecutionControls) SetMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// DeletionPolicy returns the AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSConfigRemediationConfiguration_ExecutionControls) DeletionPolicy() policies.DeletionPolicy {
	return r._deletionPolicy
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSConfigRemediationConfiguration_ExecutionControls) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}
