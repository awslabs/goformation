package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSElasticBeanstalkEnvironment AWS CloudFormation Resource (AWS::ElasticBeanstalk::Environment)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html
type AWSElasticBeanstalkEnvironment struct {

	// ApplicationName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-applicationname
	ApplicationName string `json:"ApplicationName"`

	// CNAMEPrefix AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-cnameprefix
	CNAMEPrefix string `json:"CNAMEPrefix"`

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-description
	Description string `json:"Description"`

	// EnvironmentName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-name
	EnvironmentName string `json:"EnvironmentName"`

	// OptionSettings AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-optionsettings
	OptionSettings []AWSElasticBeanstalkEnvironment_OptionSettings `json:"OptionSettings"`

	// SolutionStackName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-solutionstackname
	SolutionStackName string `json:"SolutionStackName"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-elasticbeanstalk-environment-tags
	Tags []Tag `json:"Tags"`

	// TemplateName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-templatename
	TemplateName string `json:"TemplateName"`

	// Tier AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-tier
	Tier AWSElasticBeanstalkEnvironment_Tier `json:"Tier"`

	// VersionLabel AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-environment.html#cfn-beanstalk-environment-versionlabel
	VersionLabel string `json:"VersionLabel"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticBeanstalkEnvironment) AWSCloudFormationType() string {
	return "AWS::ElasticBeanstalk::Environment"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticBeanstalkEnvironment) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSElasticBeanstalkEnvironmentResources retrieves all AWSElasticBeanstalkEnvironment items from a CloudFormation template
func GetAllAWSElasticBeanstalkEnvironmentResources(template *Template) map[string]*AWSElasticBeanstalkEnvironment {

	results := map[string]*AWSElasticBeanstalkEnvironment{}
	for name, resource := range template.Resources {
		result := &AWSElasticBeanstalkEnvironment{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSElasticBeanstalkEnvironmentWithName retrieves all AWSElasticBeanstalkEnvironment items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSElasticBeanstalkEnvironmentWithName(name string, template *Template) (*AWSElasticBeanstalkEnvironment, error) {

	result := &AWSElasticBeanstalkEnvironment{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSElasticBeanstalkEnvironment{}, errors.New("resource not found")

}
