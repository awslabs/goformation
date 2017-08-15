package resources

// AWS::Elasticsearch::Domain.ElasticsearchClusterConfig AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-elasticsearchclusterconfig.html
type AWSElasticsearchDomain_ElasticsearchClusterConfig struct {

	// DedicatedMasterCount AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-elasticsearchclusterconfig.html#cfn-elasticsearch-domain-elasticseachclusterconfig-dedicatedmastercount
	DedicatedMasterCount int64 `json:"DedicatedMasterCount"`

	// DedicatedMasterEnabled AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-elasticsearchclusterconfig.html#cfn-elasticsearch-domain-elasticseachclusterconfig-dedicatedmasterenabled
	DedicatedMasterEnabled bool `json:"DedicatedMasterEnabled"`

	// DedicatedMasterType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-elasticsearchclusterconfig.html#cfn-elasticsearch-domain-elasticseachclusterconfig-dedicatedmastertype
	DedicatedMasterType string `json:"DedicatedMasterType"`

	// InstanceCount AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-elasticsearchclusterconfig.html#cfn-elasticsearch-domain-elasticseachclusterconfig-instancecount
	InstanceCount int64 `json:"InstanceCount"`

	// InstanceType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-elasticsearchclusterconfig.html#cfn-elasticsearch-domain-elasticseachclusterconfig-instnacetype
	InstanceType string `json:"InstanceType"`

	// ZoneAwarenessEnabled AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-elasticsearchclusterconfig.html#cfn-elasticsearch-domain-elasticseachclusterconfig-zoneawarenessenabled
	ZoneAwarenessEnabled bool `json:"ZoneAwarenessEnabled"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticsearchDomain_ElasticsearchClusterConfig) AWSCloudFormationType() string {
	return "AWS::Elasticsearch::Domain.ElasticsearchClusterConfig"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticsearchDomain_ElasticsearchClusterConfig) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}
