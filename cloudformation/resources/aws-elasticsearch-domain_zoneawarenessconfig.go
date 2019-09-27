package resources

import "github.com/awslabs/goformation/cloudformation/policies"

// AWSElasticsearchDomain_ZoneAwarenessConfig AWS CloudFormation Resource (AWS::Elasticsearch::Domain.ZoneAwarenessConfig)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-zoneawarenessconfig.html
type AWSElasticsearchDomain_ZoneAwarenessConfig struct {

	// AvailabilityZoneCount AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-zoneawarenessconfig.html#cfn-elasticsearch-domain-zoneawarenessconfig-availabilityzonecount
	AvailabilityZoneCount int `json:"AvailabilityZoneCount,omitempty"`

	// _deletionPolicy represents a CloudFormation DeletionPolicy
	_deletionPolicy policies.DeletionPolicy

	// _dependsOn stores the logical ID of the resources to be created before this resource
	_dependsOn []string

	// _metadata stores structured data associated with this resource
	_metadata map[string]interface{}
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticsearchDomain_ZoneAwarenessConfig) AWSCloudFormationType() string {
	return "AWS::Elasticsearch::Domain.ZoneAwarenessConfig"
}

// DependsOn returns a slice of logical ID names this resource depends on.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSElasticsearchDomain_ZoneAwarenessConfig) DependsOn() []string {
	return r._dependsOn
}

// SetDependsOn specify that the creation of this resource follows another.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-dependson.html
func (r *AWSElasticsearchDomain_ZoneAwarenessConfig) SetDependsOn(dependencies []string) {
	r._dependsOn = dependencies
}

// Metadata returns the metadata associated with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSElasticsearchDomain_ZoneAwarenessConfig) Metadata() map[string]interface{} {
	return r._metadata
}

// SetMetadata enables you to associate structured data with this resource.
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-metadata.html
func (r *AWSElasticsearchDomain_ZoneAwarenessConfig) SetMetadata(metadata map[string]interface{}) {
	r._metadata = metadata
}

// DeletionPolicy returns the AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSElasticsearchDomain_ZoneAwarenessConfig) DeletionPolicy() policies.DeletionPolicy {
	return r._deletionPolicy
}

// SetDeletionPolicy applies an AWS CloudFormation DeletionPolicy to this resource
// see: https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-attribute-deletionpolicy.html
func (r *AWSElasticsearchDomain_ZoneAwarenessConfig) SetDeletionPolicy(policy policies.DeletionPolicy) {
	r._deletionPolicy = policy
}
