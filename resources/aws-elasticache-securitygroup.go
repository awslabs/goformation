package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ElastiCache::SecurityGroup AWS CloudFormation Resource
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
func GetAllAWSElastiCacheSecurityGroup(template *Template) map[string]*AWSElastiCacheSecurityGroup {

	results := map[string]*AWSElastiCacheSecurityGroup{}
	for name, resource := range template.Resources {
		result := &AWSElastiCacheSecurityGroup{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSElastiCacheSecurityGroupWithName retrieves all AWSElastiCacheSecurityGroup items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSElastiCacheSecurityGroup(name string, template *Template) (*AWSElastiCacheSecurityGroup, error) {

	result := &AWSElastiCacheSecurityGroup{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSElastiCacheSecurityGroup{}, errors.New("resource not found")

}
