package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::Elasticsearch::Domain.EBSOptions AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-ebsoptions.html
type AWSElasticsearchDomain_EBSOptions struct {

	// EBSEnabled AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-ebsoptions.html#cfn-elasticsearch-domain-ebsoptions-ebsenabled

	EBSEnabled bool `json:"EBSEnabled"`

	// Iops AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-ebsoptions.html#cfn-elasticsearch-domain-ebsoptions-iops

	Iops int64 `json:"Iops"`

	// VolumeSize AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-ebsoptions.html#cfn-elasticsearch-domain-ebsoptions-volumesize

	VolumeSize int64 `json:"VolumeSize"`

	// VolumeType AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticsearch-domain-ebsoptions.html#cfn-elasticsearch-domain-ebsoptions-volumetype

	VolumeType string `json:"VolumeType"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticsearchDomain_EBSOptions) AWSCloudFormationType() string {
	return "AWS::Elasticsearch::Domain.EBSOptions"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticsearchDomain_EBSOptions) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSElasticsearchDomain_EBSOptionsResources retrieves all AWSElasticsearchDomain_EBSOptions items from a CloudFormation template
func GetAllAWSElasticsearchDomain_EBSOptions(template *Template) map[string]*AWSElasticsearchDomain_EBSOptions {

	results := map[string]*AWSElasticsearchDomain_EBSOptions{}
	for name, resource := range template.Resources {
		result := &AWSElasticsearchDomain_EBSOptions{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSElasticsearchDomain_EBSOptionsWithName retrieves all AWSElasticsearchDomain_EBSOptions items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSElasticsearchDomain_EBSOptions(name string, template *Template) (*AWSElasticsearchDomain_EBSOptions, error) {

	result := &AWSElasticsearchDomain_EBSOptions{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSElasticsearchDomain_EBSOptions{}, errors.New("resource not found")

}
