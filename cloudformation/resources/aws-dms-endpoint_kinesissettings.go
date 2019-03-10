package resources

import "github.com/awslabs/goformation/cloudformation/policies"

// AWSDMSEndpoint_KinesisSettings AWS CloudFormation Resource (AWS::DMS::Endpoint.KinesisSettings)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-kinesissettings.html
type AWSDMSEndpoint_KinesisSettings struct {

	// MessageFormat AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-kinesissettings.html#cfn-dms-endpoint-kinesissettings-messageformat
	MessageFormat string `json:"MessageFormat,omitempty"`

	// ServiceAccessRoleArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-kinesissettings.html#cfn-dms-endpoint-kinesissettings-serviceaccessrolearn
	ServiceAccessRoleArn string `json:"ServiceAccessRoleArn,omitempty"`

	// StreamArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-dms-endpoint-kinesissettings.html#cfn-dms-endpoint-kinesissettings-streamarn
	StreamArn string `json:"StreamArn,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy policies.DeletionPolicy

	// _dependsOn stores the logical ID of the resources to be created before this resource
	_dependsOn []string

	// _metadata stores structured data associated with this resource
	_metadata map[string]interface{}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSDMSEndpoint_KinesisSettings) AWSCloudFormationType() string {
	return "AWS::DMS::Endpoint.KinesisSettings"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSDMSEndpoint_KinesisSettings) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSDMSEndpoint_KinesisSettings) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// Metadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSDMSEndpoint_KinesisSettings) Metadata() map[string]interface{} {
	return r._metadata
}

// SetMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSDMSEndpoint_KinesisSettings) SetMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSDMSEndpoint_KinesisSettings) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}
