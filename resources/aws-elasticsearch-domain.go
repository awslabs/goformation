package resources

// AWS::Elasticsearch::Domain AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticsearch-domain.html
type AWSElasticsearchDomain struct {

	// AccessPolicies AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticsearch-domain.html#cfn-elasticsearch-domain-accesspolicies
	AccessPolicies object `json:"AccessPolicies"`

	// AdvancedOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticsearch-domain.html#cfn-elasticsearch-domain-advancedoptions
	AdvancedOptions map[string]AWSElasticsearchDomainstring `json:"AdvancedOptions"`

	// DomainName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticsearch-domain.html#cfn-elasticsearch-domain-domainname
	DomainName string `json:"DomainName"`

	// EBSOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticsearch-domain.html#cfn-elasticsearch-domain-ebsoptions
	EBSOptions AWSElasticsearchDomainEBSOptions `json:"EBSOptions"`

	// ElasticsearchClusterConfig AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticsearch-domain.html#cfn-elasticsearch-domain-elasticsearchclusterconfig
	ElasticsearchClusterConfig AWSElasticsearchDomainElasticsearchClusterConfig `json:"ElasticsearchClusterConfig"`

	// ElasticsearchVersion AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticsearch-domain.html#cfn-elasticsearch-domain-elasticsearchversion
	ElasticsearchVersion string `json:"ElasticsearchVersion"`

	// SnapshotOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticsearch-domain.html#cfn-elasticsearch-domain-snapshotoptions
	SnapshotOptions AWSElasticsearchDomainSnapshotOptions `json:"SnapshotOptions"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticsearch-domain.html#cfn-elasticsearch-domain-tags
	Tags []AWSElasticsearchDomainTag `json:"Tags"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticsearchDomain) AWSCloudFormationType() string {
	return "AWS::Elasticsearch::Domain"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticsearchDomain) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
