package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSElastiCacheSecurityGroupIngress AWS CloudFormation Resource (AWS::ElastiCache::SecurityGroupIngress)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-security-group-ingress.html
type AWSElastiCacheSecurityGroupIngress struct {

	// CacheSecurityGroupName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-security-group-ingress.html#cfn-elasticache-securitygroupingress-cachesecuritygroupname
	CacheSecurityGroupName string `json:"CacheSecurityGroupName"`

	// EC2SecurityGroupName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-security-group-ingress.html#cfn-elasticache-securitygroupingress-ec2securitygroupname
	EC2SecurityGroupName string `json:"EC2SecurityGroupName"`

	// EC2SecurityGroupOwnerId AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-security-group-ingress.html#cfn-elasticache-securitygroupingress-ec2securitygroupownerid
	EC2SecurityGroupOwnerId string `json:"EC2SecurityGroupOwnerId"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElastiCacheSecurityGroupIngress) AWSCloudFormationType() string {
	return "AWS::ElastiCache::SecurityGroupIngress"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElastiCacheSecurityGroupIngress) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSElastiCacheSecurityGroupIngressResources retrieves all AWSElastiCacheSecurityGroupIngress items from a CloudFormation template
func (t *Template) GetAllAWSElastiCacheSecurityGroupIngressResources() map[string]*AWSElastiCacheSecurityGroupIngress {

	results := map[string]*AWSElastiCacheSecurityGroupIngress{}
	for name, resource := range t.Resources {
		result := &AWSElastiCacheSecurityGroupIngress{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSElastiCacheSecurityGroupIngressWithName retrieves all AWSElastiCacheSecurityGroupIngress items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSElastiCacheSecurityGroupIngressWithName(name string) (*AWSElastiCacheSecurityGroupIngress, error) {

	result := &AWSElastiCacheSecurityGroupIngress{}
	if resource, ok := t.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSElastiCacheSecurityGroupIngress{}, errors.New("resource not found")

}
