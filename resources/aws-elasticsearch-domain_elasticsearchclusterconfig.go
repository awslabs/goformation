package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

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

// GetAllAWSElasticsearchDomain_ElasticsearchClusterConfigResources retrieves all AWSElasticsearchDomain_ElasticsearchClusterConfig items from a CloudFormation template
func GetAllAWSElasticsearchDomain_ElasticsearchClusterConfig(template *Template) map[string]*AWSElasticsearchDomain_ElasticsearchClusterConfig {

	results := map[string]*AWSElasticsearchDomain_ElasticsearchClusterConfig{}
	for name, resource := range template.Resources {
		result := &AWSElasticsearchDomain_ElasticsearchClusterConfig{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSElasticsearchDomain_ElasticsearchClusterConfigWithName retrieves all AWSElasticsearchDomain_ElasticsearchClusterConfig items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSElasticsearchDomain_ElasticsearchClusterConfig(name string, template *Template) (*AWSElasticsearchDomain_ElasticsearchClusterConfig, error) {

	result := &AWSElasticsearchDomain_ElasticsearchClusterConfig{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSElasticsearchDomain_ElasticsearchClusterConfig{}, errors.New("resource not found")

}
