package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSElasticsearchDomain AWS CloudFormation Resource (AWS::Elasticsearch::Domain)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticsearch-domain.html
type AWSElasticsearchDomain struct {

	// AccessPolicies AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticsearch-domain.html#cfn-elasticsearch-domain-accesspolicies
	AccessPolicies interface{} `json:"AccessPolicies"`

	// AdvancedOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticsearch-domain.html#cfn-elasticsearch-domain-advancedoptions
	AdvancedOptions map[string]string `json:"AdvancedOptions"`

	// DomainName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticsearch-domain.html#cfn-elasticsearch-domain-domainname
	DomainName string `json:"DomainName"`

	// EBSOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticsearch-domain.html#cfn-elasticsearch-domain-ebsoptions
	EBSOptions AWSElasticsearchDomain_EBSOptions `json:"EBSOptions"`

	// ElasticsearchClusterConfig AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticsearch-domain.html#cfn-elasticsearch-domain-elasticsearchclusterconfig
	ElasticsearchClusterConfig AWSElasticsearchDomain_ElasticsearchClusterConfig `json:"ElasticsearchClusterConfig"`

	// ElasticsearchVersion AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticsearch-domain.html#cfn-elasticsearch-domain-elasticsearchversion
	ElasticsearchVersion string `json:"ElasticsearchVersion"`

	// SnapshotOptions AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticsearch-domain.html#cfn-elasticsearch-domain-snapshotoptions
	SnapshotOptions AWSElasticsearchDomain_SnapshotOptions `json:"SnapshotOptions"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticsearch-domain.html#cfn-elasticsearch-domain-tags
	Tags []Tag `json:"Tags"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticsearchDomain) AWSCloudFormationType() string {
	return "AWS::Elasticsearch::Domain"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticsearchDomain) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSElasticsearchDomainResources retrieves all AWSElasticsearchDomain items from a CloudFormation template
func (t *Template) GetAllAWSElasticsearchDomainResources() map[string]*AWSElasticsearchDomain {

	results := map[string]*AWSElasticsearchDomain{}
	for name, resource := range t.Resources {
		result := &AWSElasticsearchDomain{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSElasticsearchDomainWithName retrieves all AWSElasticsearchDomain items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSElasticsearchDomainWithName(name string) (*AWSElasticsearchDomain, error) {

	result := &AWSElasticsearchDomain{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSElasticsearchDomain{}, errors.New("resource not found")

}
