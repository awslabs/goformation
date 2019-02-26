package cloudformation

// AWSBatchJobDefinition_NodeProperties AWS CloudFormation Resource (AWS::Batch::JobDefinition.NodeProperties)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-nodeproperties.html
type AWSBatchJobDefinition_NodeProperties struct {

	// MainNode AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-nodeproperties.html#cfn-batch-jobdefinition-nodeproperties-mainnode
	MainNode int `json:"MainNode,omitempty"`

	// NodeRangeProperties AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-nodeproperties.html#cfn-batch-jobdefinition-nodeproperties-noderangeproperties
	NodeRangeProperties []AWSBatchJobDefinition_NodeRangeProperty `json:"NodeRangeProperties,omitempty"`

	// NumNodes AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-nodeproperties.html#cfn-batch-jobdefinition-nodeproperties-numnodes
	NumNodes int `json:"NumNodes,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSBatchJobDefinition_NodeProperties) AWSCloudFormationType() string {
	return "AWS::Batch::JobDefinition.NodeProperties"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSBatchJobDefinition_NodeProperties) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
