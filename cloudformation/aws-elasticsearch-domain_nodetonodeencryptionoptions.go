package cloudformation

// AWSElasticsearchDomain_NodeToNodeEncryptionOptions AWS CloudFormation Resource (AWS::Elasticsearch::Domain.NodeToNodeEncryptionOptions)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-nodetonodeencryptionoptions.html
type AWSElasticsearchDomain_NodeToNodeEncryptionOptions struct {

	// Enabled AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-nodetonodeencryptionoptions.html#cfn-elasticsearch-domain-nodetonodeencryptionoptions-enabled
	Enabled bool `json:"Enabled,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy DeletionPolicy
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticsearchDomain_NodeToNodeEncryptionOptions) AWSCloudFormationType() string {
	return "AWS::Elasticsearch::Domain.NodeToNodeEncryptionOptions"
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSElasticsearchDomain_NodeToNodeEncryptionOptions) SetDeletionPolicy(policy DeletionPolicy) {
	r._deletionPolicy = policy
}
