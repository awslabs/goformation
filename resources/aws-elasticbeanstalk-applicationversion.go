package resources

import (
	"errors"

	"github.com/mitchellh/mapstructure"
)

// AWS::ElasticBeanstalk::ApplicationVersion AWS CloudFormation Resource
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-version.html
type AWSElasticBeanstalkApplicationVersion struct {

	// ApplicationName AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-version.html#cfn-elasticbeanstalk-applicationversion-applicationname
	ApplicationName string `json:"ApplicationName"`

	// Description AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-version.html#cfn-elasticbeanstalk-applicationversion-description
	Description string `json:"Description"`

	// SourceBundle AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-beanstalk-version.html#cfn-elasticbeanstalk-applicationversion-sourcebundle
	SourceBundle AWSElasticBeanstalkApplicationVersion_SourceBundle `json:"SourceBundle"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticBeanstalkApplicationVersion) AWSCloudFormationType() string {
	return "AWS::ElasticBeanstalk::ApplicationVersion"
}

// AWSCloudFormationSpecificationVersion returns the AWS Specification Version that this resource was generated from
func (r *AWSElasticBeanstalkApplicationVersion) AWSCloudFormationSpecificationVersion() string {
	return "1.4.2"
}

// GetAllAWSElasticBeanstalkApplicationVersionResources retrieves all AWSElasticBeanstalkApplicationVersion items from a CloudFormation template
func GetAllAWSElasticBeanstalkApplicationVersionResources(template *Template) map[string]*AWSElasticBeanstalkApplicationVersion {

	results := map[string]*AWSElasticBeanstalkApplicationVersion{}
	for name, resource := range template.Resources {
		result := &AWSElasticBeanstalkApplicationVersion{}
		if err := mapstructure.Decode(resource, result); err == nil {
			results[name] = result
		}
	}
	return results

}

// GetAWSElasticBeanstalkApplicationVersionWithName retrieves all AWSElasticBeanstalkApplicationVersion items from a CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func GetAWSElasticBeanstalkApplicationVersionWithName(name string, template *Template) (*AWSElasticBeanstalkApplicationVersion, error) {

	result := &AWSElasticBeanstalkApplicationVersion{}
	if resource, ok := template.Resources[name]; ok {
		if err := mapstructure.Decode(resource, result); err == nil {
			return result, nil
		}
	}

	return &AWSElasticBeanstalkApplicationVersion{}, errors.New("resource not found")

}
