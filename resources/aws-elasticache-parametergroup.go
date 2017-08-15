package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSElastiCacheParameterGroup AWS CloudFormation Resource (AWS::ElastiCache::ParameterGroup)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-parameter-group.html
type AWSElastiCacheParameterGroup struct {

	// CacheParameterGroupFamily AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-parameter-group.html#cfn-elasticache-parametergroup-cacheparametergroupfamily
	CacheParameterGroupFamily string `json:"CacheParameterGroupFamily"`

	// Description AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-parameter-group.html#cfn-elasticache-parametergroup-description
	Description string `json:"Description"`

	// Properties AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-parameter-group.html#cfn-elasticache-parametergroup-properties
	Properties map[string]string `json:"Properties"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElastiCacheParameterGroup) AWSCloudFormationType() string {
	return "AWS::ElastiCache::ParameterGroup"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElastiCacheParameterGroup) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSElastiCacheParameterGroupResources retrieves all AWSElastiCacheParameterGroup items from a CloudFormation template
func GetAllAWSElastiCacheParameterGroupResources(template *Template) map[string]*AWSElastiCacheParameterGroup {

	results := map[string]*AWSElastiCacheParameterGroup{}
	for name, resource := range template.Resources {
		result := &AWSElastiCacheParameterGroup{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSElastiCacheParameterGroupWithName retrieves all AWSElastiCacheParameterGroup items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSElastiCacheParameterGroupWithName(name string, template *Template) (*AWSElastiCacheParameterGroup, error) {

	result := &AWSElastiCacheParameterGroup{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSElastiCacheParameterGroup{}, errors.New("resource not found")

}
