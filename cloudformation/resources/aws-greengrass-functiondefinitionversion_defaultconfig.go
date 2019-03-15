package resources

import "github.com/awslabs/goformation/cloudformation/policies"

// AWSGreengrassFunctionDefinitionVersion_DefaultConfig AWS CloudFormation Resource (AWS::Greengrass::FunctionDefinitionVersion.DefaultConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinitionversion-defaultconfig.html
type AWSGreengrassFunctionDefinitionVersion_DefaultConfig struct {

	// Execution AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-greengrass-functiondefinitionversion-defaultconfig.html#cfn-greengrass-functiondefinitionversion-defaultconfig-execution
	Execution *AWSGreengrassFunctionDefinitionVersion_Execution `json:"Execution,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy policies.DeletionPolicy

	// _dependsOn stores the logical ID of the resources to be created before this resource
	_dependsOn []string

	// _metadata stores structured data associated with this resource
	_metadata map[string]interface{}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSGreengrassFunctionDefinitionVersion_DefaultConfig) AWSCloudFormationType() string {
	return "AWS::Greengrass::FunctionDefinitionVersion.DefaultConfig"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSGreengrassFunctionDefinitionVersion_DefaultConfig) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSGreengrassFunctionDefinitionVersion_DefaultConfig) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// Metadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSGreengrassFunctionDefinitionVersion_DefaultConfig) Metadata() map[string]interface{} {
	return r._metadata
}

// SetMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSGreengrassFunctionDefinitionVersion_DefaultConfig) SetMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSGreengrassFunctionDefinitionVersion_DefaultConfig) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}
