package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSElastiCacheSecurityGroup AWS CloudFormation Resource (AWS::ElastiCache::SecurityGroup)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-security-group.html
type AWSElastiCacheSecurityGroup struct {

	// Description AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-security-group.html#cfn-elasticache-securitygroup-description
	Description string `json:"Description"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElastiCacheSecurityGroup) AWSCloudFormationType() string {
	return "AWS::ElastiCache::SecurityGroup"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElastiCacheSecurityGroup) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSElastiCacheSecurityGroupResources retrieves all AWSElastiCacheSecurityGroup items from a CloudFormation template
func (t *Template) GetAllAWSElastiCacheSecurityGroupResources() map[string]*AWSElastiCacheSecurityGroup {

	results := map[string]*AWSElastiCacheSecurityGroup{}
	for name, resource := range t.Resources {
		result := &AWSElastiCacheSecurityGroup{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSElastiCacheSecurityGroupWithName retrieves all AWSElastiCacheSecurityGroup items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSElastiCacheSecurityGroupWithName(name string) (*AWSElastiCacheSecurityGroup, error) {

	result := &AWSElastiCacheSecurityGroup{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSElastiCacheSecurityGroup{}, errors.New("resource not found")

}
