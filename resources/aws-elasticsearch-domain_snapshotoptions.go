package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Elasticsearch::Domain.SnapshotOptions AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-snapshotoptions.html
type AWSElasticsearchDomain_SnapshotOptions struct {

	// AutomatedSnapshotStartHour AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-snapshotoptions.html#cfn-elasticsearch-domain-snapshotoptions-automatedsnapshotstarthour
	AutomatedSnapshotStartHour int64 `json:"AutomatedSnapshotStartHour"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticsearchDomain_SnapshotOptions) AWSCloudFormationType() string {
	return "AWS::Elasticsearch::Domain.SnapshotOptions"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticsearchDomain_SnapshotOptions) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSElasticsearchDomain_SnapshotOptionsResources retrieves all AWSElasticsearchDomain_SnapshotOptions items from a CloudFormation template
func GetAllAWSElasticsearchDomain_SnapshotOptions(template *Template) map[string]*AWSElasticsearchDomain_SnapshotOptions {

	results := map[string]*AWSElasticsearchDomain_SnapshotOptions{}
	for name, resource := range template.Resources {
		result := &AWSElasticsearchDomain_SnapshotOptions{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSElasticsearchDomain_SnapshotOptionsWithName retrieves all AWSElasticsearchDomain_SnapshotOptions items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSElasticsearchDomain_SnapshotOptions(name string, template *Template) (*AWSElasticsearchDomain_SnapshotOptions, error) {

	result := &AWSElasticsearchDomain_SnapshotOptions{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSElasticsearchDomain_SnapshotOptions{}, errors.New("resource not found")

}
