package resources

import "github.com/awslabs/goformation/cloudformation/policies"

// AWSEC2LaunchTemplate_TagSpecification AWS CloudFormation Resource (AWS::EC2::LaunchTemplate.TagSpecification)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-tagspecification.html
type AWSEC2LaunchTemplate_TagSpecification struct {

	// ResourceType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-tagspecification.html#cfn-ec2-launchtemplate-tagspecification-resourcetype
	ResourceType string `json:"ResourceType,omitempty"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-launchtemplate-tagspecification.html#cfn-ec2-launchtemplate-tagspecification-tags
	Tags []Tag `json:"Tags,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy policies.DeletionPolicy

	// _dependsOn stores the logical ID of the resources to be created before this resource
	_dependsOn []string

	// _metadata stores structured data associated with this resource
	_metadata map[string]interface{}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSEC2LaunchTemplate_TagSpecification) AWSCloudFormationType() string {
	return "AWS::EC2::LaunchTemplate.TagSpecification"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSEC2LaunchTemplate_TagSpecification) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSEC2LaunchTemplate_TagSpecification) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// Metadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSEC2LaunchTemplate_TagSpecification) Metadata() map[string]interface{} {
	return r._metadata
}

// SetMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSEC2LaunchTemplate_TagSpecification) SetMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSEC2LaunchTemplate_TagSpecification) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}
