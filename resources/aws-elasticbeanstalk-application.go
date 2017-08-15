package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWSElasticBeanstalkApplication AWS CloudFormation Resource (AWS::ElasticBeanstalk::Application)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk.html
type AWSElasticBeanstalkApplication struct {

	// ApplicationName AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk.html#cfn-elasticbeanstalk-application-name
	ApplicationName string `json:"ApplicationName"`

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk.html#cfn-elasticbeanstalk-application-description
	Description string `json:"Description"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticBeanstalkApplication) AWSCloudFormationType() string {
	return "AWS::ElasticBeanstalk::Application"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticBeanstalkApplication) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSElasticBeanstalkApplicationResources retrieves all AWSElasticBeanstalkApplication items from a CloudFormation template
func GetAllAWSElasticBeanstalkApplicationResources(template *Template) map[string]*AWSElasticBeanstalkApplication {

	results := map[string]*AWSElasticBeanstalkApplication{}
	for name, resource := range template.Resources {
		result := &AWSElasticBeanstalkApplication{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSElasticBeanstalkApplicationWithName retrieves all AWSElasticBeanstalkApplication items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSElasticBeanstalkApplicationWithName(name string, template *Template) (*AWSElasticBeanstalkApplication, error) {

	result := &AWSElasticBeanstalkApplication{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSElasticBeanstalkApplication{}, errors.New("resource not found")

}
