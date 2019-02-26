package cloudformation

// AWSAmazonMQConfigurationAssociation_ConfigurationId AWS CloudFormation Resource (AWS::AmazonMQ::ConfigurationAssociation.ConfigurationId)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-configurationassociation-configurationid.html
type AWSAmazonMQConfigurationAssociation_ConfigurationId struct {

	// Id AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-configurationassociation-configurationid.html#cfn-amazonmq-configurationassociation-configurationid-id
	Id string `json:"Id,omitempty"`

	// Revision AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-amazonmq-configurationassociation-configurationid.html#cfn-amazonmq-configurationassociation-configurationid-revision
	Revision int `json:"Revision,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSAmazonMQConfigurationAssociation_ConfigurationId) AWSCloudFormationType() string {
	return "AWS::AmazonMQ::ConfigurationAssociation.ConfigurationId"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSAmazonMQConfigurationAssociation_ConfigurationId) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
