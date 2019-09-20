package resources

import "github.com/TachyonNexus/goformation/cloudformation/policies"

// AWSServerlessApplication_ApplicationLocation AWS CloudFormation Resource (AWS::Serverless::Application.ApplicationLocation)
// See: https://github.com/TachyonNexus/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessapplication
type AWSServerlessApplication_ApplicationLocation struct {

	// ApplicationId AWS CloudFormation Property
	// Required: true
	// See: https://github.com/TachyonNexus/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessapplication
	ApplicationId string `json:"ApplicationId,omitempty"`

	// SemanticVersion AWS CloudFormation Property
	// Required: true
	// See: https://github.com/TachyonNexus/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessapplication
	SemanticVersion string `json:"SemanticVersion,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy policies.DeletionPolicy

	// _dependsOn stores the logical ID of the resources to be created before this resource
	_dependsOn []string

	// _metadata stores structured data associated with this resource
	_metadata map[string]interface{}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSServerlessApplication_ApplicationLocation) AWSCloudFormationType() string {
	return "AWS::Serverless::Application.ApplicationLocation"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSServerlessApplication_ApplicationLocation) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSServerlessApplication_ApplicationLocation) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// Metadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSServerlessApplication_ApplicationLocation) Metadata() map[string]interface{} {
	return r._metadata
}

// SetMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSServerlessApplication_ApplicationLocation) SetMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// DeletionPolicy returns the AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSServerlessApplication_ApplicationLocation) DeletionPolicy() policies.DeletionPolicy {
	return r._deletionPolicy
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSServerlessApplication_ApplicationLocation) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}
