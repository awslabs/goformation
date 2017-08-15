package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ElastiCache::SubnetGroup AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-subnetgroup.html
type AWSElastiCacheSubnetGroup struct {

	// CacheSubnetGroupName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-subnetgroup.html#cfn-elasticache-subnetgroup-cachesubnetgroupname
	CacheSubnetGroupName string `json:"CacheSubnetGroupName"`

	// Description AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-subnetgroup.html#cfn-elasticache-subnetgroup-description
	Description string `json:"Description"`

	// SubnetIds AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-subnetgroup.html#cfn-elasticache-subnetgroup-subnetids
	SubnetIds []string `json:"SubnetIds"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElastiCacheSubnetGroup) AWSCloudFormationType() string {
	return "AWS::ElastiCache::SubnetGroup"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElastiCacheSubnetGroup) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSElastiCacheSubnetGroupResources retrieves all AWSElastiCacheSubnetGroup items from a CloudFormation template
func GetAllAWSElastiCacheSubnetGroupResources(template *Template) map[string]*AWSElastiCacheSubnetGroup {

	results := map[string]*AWSElastiCacheSubnetGroup{}
	for name, resource := range template.Resources {
		result := &AWSElastiCacheSubnetGroup{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSElastiCacheSubnetGroupWithName retrieves all AWSElastiCacheSubnetGroup items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSElastiCacheSubnetGroupWithName(name string, template *Template) (*AWSElastiCacheSubnetGroup, error) {

	result := &AWSElastiCacheSubnetGroup{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSElastiCacheSubnetGroup{}, errors.New("resource not found")

}
