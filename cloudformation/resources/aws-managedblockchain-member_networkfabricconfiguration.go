package resources

import "github.com/awslabs/goformation/cloudformation/policies"

// AWSManagedBlockchainMember_NetworkFabricConfiguration AWS CloudFormation Resource (AWS::ManagedBlockchain::Member.NetworkFabricConfiguration)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-managedblockchain-member-networkfabricconfiguration.html
type AWSManagedBlockchainMember_NetworkFabricConfiguration struct {

	// Edition AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-managedblockchain-member-networkfabricconfiguration.html#cfn-managedblockchain-member-networkfabricconfiguration-edition
	Edition string `json:"Edition,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy policies.DeletionPolicy

	// _dependsOn stores the logical ID of the resources to be created before this resource
	_dependsOn []string

	// _metadata stores structured data associated with this resource
	_metadata map[string]interface{}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSManagedBlockchainMember_NetworkFabricConfiguration) AWSCloudFormationType() string {
	return "AWS::ManagedBlockchain::Member.NetworkFabricConfiguration"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSManagedBlockchainMember_NetworkFabricConfiguration) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSManagedBlockchainMember_NetworkFabricConfiguration) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// Metadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSManagedBlockchainMember_NetworkFabricConfiguration) Metadata() map[string]interface{} {
	return r._metadata
}

// SetMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSManagedBlockchainMember_NetworkFabricConfiguration) SetMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSManagedBlockchainMember_NetworkFabricConfiguration) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}
