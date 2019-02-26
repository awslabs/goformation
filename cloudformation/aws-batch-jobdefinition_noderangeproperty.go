package cloudformation

// AWSBatchJobDefinition_NodeRangeProperty AWS CloudFormation Resource (AWS::Batch::JobDefinition.NodeRangeProperty)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-noderangeproperty.html
type AWSBatchJobDefinition_NodeRangeProperty struct {

	// Container AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-noderangeproperty.html#cfn-batch-jobdefinition-noderangeproperty-container
	Container *AWSBatchJobDefinition_ContainerProperties `json:"Container,omitempty"`

	// TargetNodes AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-batch-jobdefinition-noderangeproperty.html#cfn-batch-jobdefinition-noderangeproperty-targetnodes
	TargetNodes string `json:"TargetNodes,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSBatchJobDefinition_NodeRangeProperty) AWSCloudFormationType() string {
	return "AWS::Batch::JobDefinition.NodeRangeProperty"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSBatchJobDefinition_NodeRangeProperty) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
