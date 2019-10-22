package resources

import "github.com/awslabs/goformation/cloudformation/policies"

// AWSGreengrassResourceDefinition_ResourceInstance AWS CloudFormation Resource (AWS::Greengrass::ResourceDefinition.ResourceInstance)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-resourceinstance.html
type AWSGreengrassResourceDefinition_ResourceInstance struct {

	// Id AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-resourceinstance.html#cfn-greengrass-resourcedefinition-resourceinstance-id
	Id string `json:"Id,omitempty"`

	// Name AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-resourceinstance.html#cfn-greengrass-resourcedefinition-resourceinstance-name
	Name string `json:"Name,omitempty"`

	// ResourceDataContainer AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-resourcedefinition-resourceinstance.html#cfn-greengrass-resourcedefinition-resourceinstance-resourcedatacontainer
	ResourceDataContainer *AWSGreengrassResourceDefinition_ResourceDataContainer `json:"ResourceDataContainer,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy policies.DeletionPolicy

	// _dependsOn stores the logical ID of the resources to be created before this resource
	_dependsOn []string

	// _metadata stores structured data associated with this resource
	_metadata map[string]interface{}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSGreengrassResourceDefinition_ResourceInstance) AWSCloudFormationType() string {
	return "AWS::Greengrass::ResourceDefinition.ResourceInstance"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSGreengrassResourceDefinition_ResourceInstance) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSGreengrassResourceDefinition_ResourceInstance) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// Metadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSGreengrassResourceDefinition_ResourceInstance) Metadata() map[string]interface{} {
	return r._metadata
}

// SetMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSGreengrassResourceDefinition_ResourceInstance) SetMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSGreengrassResourceDefinition_ResourceInstance) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}
