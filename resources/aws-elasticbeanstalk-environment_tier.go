package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ElasticBeanstalk::Environment.Tier AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment-tier.html
type AWSElasticBeanstalkEnvironment_Tier struct {

	// Name AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment-tier.html#cfn-beanstalk-env-tier-name

	Name string `json:"Name"`

	// Type AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment-tier.html#cfn-beanstalk-env-tier-type

	Type string `json:"Type"`

	// Version AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment-tier.html#cfn-beanstalk-env-tier-version

	Version string `json:"Version"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticBeanstalkEnvironment_Tier) AWSCloudFormationType() string {
	return "AWS::ElasticBeanstalk::Environment.Tier"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticBeanstalkEnvironment_Tier) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSElasticBeanstalkEnvironment_TierResources retrieves all AWSElasticBeanstalkEnvironment_Tier items from a CloudFormation template
func GetAllAWSElasticBeanstalkEnvironment_Tier(template *Template) map[string]*AWSElasticBeanstalkEnvironment_Tier {

	results := map[string]*AWSElasticBeanstalkEnvironment_Tier{}
	for name, resource := range template.Resources {
		result := &AWSElasticBeanstalkEnvironment_Tier{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSElasticBeanstalkEnvironment_TierWithName retrieves all AWSElasticBeanstalkEnvironment_Tier items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetWithNameAWSElasticBeanstalkEnvironment_Tier(name string, template *Template) (*AWSElasticBeanstalkEnvironment_Tier, error) {

	result := &AWSElasticBeanstalkEnvironment_Tier{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSElasticBeanstalkEnvironment_Tier{}, errors.New("resource not found")

}
